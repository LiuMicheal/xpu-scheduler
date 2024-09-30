package plugins

import (
	"context"
	pb "xpu-scheduler/pkg/plugins/proto"
	"time"

	grpc "google.golang.org/grpc"
	insecure "google.golang.org/grpc/credentials/insecure"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	klog "k8s.io/klog/v2"
	"k8s.io/kubernetes/pkg/scheduler/framework"
)

const Name = "xpu-plugin"

var allocate_flag = 0

type NPUnodeinfo_t struct {
	// lock         *sync.Mutex
	AicoreUsage  uint32
	AicpuUsage   uint32
	CtrlcpuUsage uint32
	MemoryUsage  uint32
}

var NPUCache map[string]NPUnodeinfo_t

var maxaicore int64 = 200
var maxaicpu int64 = 400
var maxmem int64 = 8192

type NpuSched struct {
	handle framework.Handle
}

func (s *NpuSched) Name() string {
	return Name
}

// Get the node load through GRPC
func GetUsed(node *v1.Node) {
	// Get the node name and access the node through the mapping in hosts
	// add := node.Node().Status.Addresses[0].Address
	nodename := node.Status.Addresses[1].Address
	// Connect to GRPCserver through nodename.
	// The premise of this method is that the hosts file of the node where the scheduler is located has the nodename mapped to the node.
	connect, err := grpc.Dial(nodename+":50001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		// klog.V(3).Infof("Remote GRPC connection failed:%v", err)
		return
	}
	defer connect.Close()
	// Get NPU utilization through grpc remote procedure call
	client := pb.NewGetUsedRateClient(connect)
	resp, err := client.GetUsedRate(context.Background(), &pb.Request{Cycle: 10})
	if err != nil {
		// klog.V(3).Infof("Remote GRPC connection failed:%v", err)
		// If the connection fails, an error is returned and the node cannot obtain the load
		// klog.V(3).Infof("Unable to connect to node via GRPC %s，IP%s:%v", nodename, add, err)
		// Assign the node load in the node load map to 100%, indicating that it is not schedulable.
		usage := NPUnodeinfo_t{
			AicoreUsage:  100,
			AicpuUsage:   100,
			CtrlcpuUsage: 100,
			MemoryUsage:  100,
		}
		NPUCache[nodename] = usage

		return
	}
	usage := NPUnodeinfo_t{
		AicoreUsage:  resp.AicoreUsage,
		AicpuUsage:   resp.AicpuUsage,
		CtrlcpuUsage: resp.CtrlcpuUsage,
		MemoryUsage:  resp.MemoryUsage,
	}
	// NPUCache[nodename].lock.Lock()
	NPUCache[nodename] = usage
	// NPUCache[nodename].lock.Unlock()
}
func (s *NpuSched) Filter(ctx context.Context, state *framework.CycleState, pod *v1.Pod, node *framework.NodeInfo) *framework.Status {
	klog.V(3).Infof("filter pod: %v", pod.Name)
	// klog.V(3).Infof("filter pod node: %v", node.Node().Status.Addresses[0].Address)
	// klog.V(3).Infof("filter pod node: %v", node.Node().Status.Addresses[1].Address)
	nodename := node.Node().Name
	cpureq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("cpu")]
	memoryreq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("memory")]
	// Vcpu := node.Allocatable.MilliCPU - 4000
	ctrlcpufree := (100 - NPUCache[nodename].CtrlcpuUsage) / 100 * 4000
	aicorefree := (100 - NPUCache[nodename].AicoreUsage) / 100 * 200
	aicpufree := (100 - NPUCache[nodename].AicpuUsage) / 100 * 400
	memfree := (100 - NPUCache[nodename].MemoryUsage) / 100 * 8192
	Vaicore := node.Allocatable.ScalarResources[v1.ResourceName("nwpu/vnpu-aicore")]
	Vaicpu := node.Allocatable.ScalarResources[v1.ResourceName("nwpu/vnpu-aicpu")]
	Vmem := node.Allocatable.ScalarResources[v1.ResourceName("nwpu/vnpu-mem")]
	if node.Allocatable.ScalarResources == nil {
		// if cpureq.Value() < Vcpu && memoryreq.Value() < Vmem {// Traditional fixed resource allocation
		// 	return framework.NewStatus(framework.Success, "")
		// }
		if cpureq.Value() < int64(ctrlcpufree) && memoryreq.Value() < int64(memfree) { // Elastic resource allocation
			return framework.NewStatus(framework.Success, "")
		}
		return framework.NewStatus(framework.Unschedulable, "node resource not enough")
	}
	klog.V(3).Infof("node allocatable: %v", Vaicore)
	klog.V(3).Infof("node allocatable: %v", Vaicpu)
	klog.V(3).Infof("node allocatable: %v", Vmem)

	npumemoryreq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("nwpu/vnpu-mem")]
	aicorereq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("nwpu/vnpu-aicore")]
	aicpureq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("nwpu/vnpu-aicpu")]
	// klog.V(3).Infof("pod request aicore: %v", aicorereq.Value())
	// klog.V(3).Infof("pod request aicpu: %v", aicpureq.Value())
	// if aicorereq.Value() < Vaicore && aicpureq.Value() < Vaicpu && npumemoryreq.Value() < Vmem {// Traditional fixed resource allocation
	// 	return framework.NewStatus(framework.Success, "")
	// }
	if aicorereq.Value() < int64(aicorefree) && aicpureq.Value() < int64(aicpufree) && npumemoryreq.Value() < int64(memfree) { // Elastic resource allocation
		return framework.NewStatus(framework.Success, "")
	}
	return framework.NewStatus(framework.Unschedulable, "node pressure large")
}

func (s *NpuSched) Score(ctx context.Context, state *framework.CycleState, pod *v1.Pod, nodeName string) (int64, *framework.Status) {

	// klog.V(3).Infof(nodeName, "AIcpu usage:", s.NPUCache[nodeName].AicpuUsage)
	// klog.V(3).Infof(nodeName, "AIcore usage:", s.NPUCache[nodeName].AicoreUsage)
	// klog.V(3).Infof(nodeName, "Ctrlcpu usage:", s.NPUCache[nodeName].CtrlcpuUsage)
	// klog.V(3).Infof(nodeName, "Memory usage:", s.NPUCache[nodeName].MemoryUsage)

	// s.NPUCache[nodeName].lock.Lock()
	node := s.handle.SnapshotSharedLister().NodeInfos().Get(nodeName)
	// node := nodeList.Get(nodeName)
	cpureq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("cpu")]
	memoryreq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("memory")]
	npumemoryreq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("nwpu/vnpu-mem")]
	aicorereq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("nwpu/vnpu-aicore")]
	aicpureq := pod.Spec.Containers[0].Resources.Limits[v1.ResourceName("nwpu/vnpu-aicpu")]
	aicorefree := 100 - NPUCache[nodeName].AicoreUsage
	aicpufree := 100 - NPUCache[nodeName].AicpuUsage
	ctrlcpufree := 100 - NPUCache[nodeName].CtrlcpuUsage
	memoryfree := 100 - NPUCache[nodeName].MemoryUsage
	aicoreavl := node.Allocatable.ScalarResources[v1.ResourceName("nwpu/vnpu-aicore")]
	aicpuavl := node.Allocatable.ScalarResources[v1.ResourceName("nwpu/vnpu-aicpu")]
	memavl := node.Allocatable.ScalarResources[v1.ResourceName("nwpu/vnpu-mem")]
	// s.NPUCache[nodeName].lock.Unlock()
	weightaicore := float64(aicorereq.Value()) / float64(maxaicore)
	weightaicpu := float64(aicpureq.Value()) / float64(maxaicpu)
	weightctrlcpu := float64(cpureq.Value()) / 4000
	weightmemory := float64((memoryreq.Value() + npumemoryreq.Value())) / float64(maxmem)
	socre1 := weightaicore*float64(aicorefree) + weightaicpu*float64(aicpufree) + weightctrlcpu*float64(ctrlcpufree) + weightmemory*float64(memoryfree)
	socre2 := weightaicore*float64(aicoreavl) + weightaicpu*float64(aicpuavl) + weightmemory*float64(memavl)
	socre := (socre1 + socre2) / 2
	klog.Infoln(nodeName+" score is :", socre)
	return int64(socre), framework.NewStatus(framework.Success, nodeName)
}

// Another interface declared by ScorePlugin. If Score is implemented, this interface must also be implemented and return nil
func (s *NpuSched) ScoreExtensions() framework.ScoreExtensions {
	return nil
}

func Cachestart() {
	for {

		config, err := clientcmd.BuildConfigFromFlags("", "/root/.kube/config")
		client, err := kubernetes.NewForConfig(config)
		nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
		if err != nil {
			klog.V(3).Infof("Failed to obtain node list:%v", err)
			continue
		}
		for _, node := range nodes.Items {
			GetUsed(&node)
		}
		time.Sleep(time.Second * 10)

		//     client, err := kubernetes.NewForConfigAndClient(clientcmd.NewDefaultClientConfig(), nil)

		// 	nodes, err := client.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
		// 	if err != nil {
		// 		klog.V(3).Infof("Failed to obtain node list:%v", err)
		// 		continue
		// 	}
		// 	for _, node := range nodes {
		// 		GetUsed(node)
		// 	}
		// 	time.Sleep(time.Second * 10)

	}
}

func New(obj runtime.Object, f framework.Handle) (framework.Plugin, error) {

	klog.V(3).Infof("NpuSched plug-in start")
	NpuSchedPlugin := NpuSched{
		handle: f,
	}

	return &NpuSchedPlugin, nil
}
