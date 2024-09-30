# Chang'E: A Container Heterogeneous Task Scheduling System for Embedded XPU and RapidIO

Welcome to the **Chang'E** scheduling system! This repository contains the implementation of the **Task Scheduling Module** for the Chang'E system, which is designed for efficiently scheduling tasks across heterogeneous computing platforms, particularly those deploying CPUs, GPUs, and NPUs in extreme environments like aerospace.

The **Task Scheduling Module** in Chang'E extends the default Kubernetes scheduler to support joint scheduling of heterogeneous resources such as CPU, GPU, and NPU using our fine-grained **Resource Deep Matching (RDM)** algorithm.

## Features

- **Heterogeneous Task Scheduling**: This module dynamically schedules tasks across heterogeneous resources (CPU, GPU, and NPU), optimizing resource allocation based on real-time status and resource demands.
- **Kubernetes Integration**: Extends the Kubernetes scheduler with a custom plug-in designed for joint scheduling of XPU resources.
- **Flexible Scoring and Filtering**: Implements Kubernetes' `FilterPlugin` and `ScorePlugin` to ensure tasks are allocated to the best-suited nodes based on the specific requirements of the PODs.
- **Open Source**: The Task Scheduling Module is fully open-sourced and can be easily integrated into existing Kubernetes environments.

## Example Use Case
Imagine an aerospace scenario where multiple tasks, such as space imagery analysis or target tracking, need to be executed efficiently using heterogeneous computing resources. The Chang'E Task Scheduling Module dynamically allocates these tasks to the most suitable computing nodes, considering real-time resource availability and task priorities. This approach helps maximize resource utilization while maintaining optimal performance under varying workloads.

## System Requirements

The following software versions are required for the Task Scheduling Module:

- **Operating System**: Ubuntu 18.04
- **Kubernetes**: Version 1.24.0
- **Docker**: Version 20.10.21
- **Golang**: Version 1.20.1

## Install Kubernetes
Follow the official Kubernetes installation tutorial by selecting your preferred method of installation (e.g., using `kubeadm`, `minikube`, or cloud-based services like GKE or EKS). Here's a basic installation guide using `kubeadm`:

1. **Install dependencies**:
   - Ensure Docker is installed and running.
   - Install `kubeadm`, `kubelet`, and `kubectl` on your nodes by following the [Kubernetes official documentation](https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/install-kubeadm/).

2. **Initialize the control-plane node**:
   ```shell
   sudo kubeadm init --pod-network-cidr=192.168.0.0/16
3. **Set up your local kubectl environment**:
    ```shell
    mkdir -p $HOME/.kube
    sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
    sudo chown $(id -u):$(id -g) $HOME/.kube/config
4. **Install a Pod network add-on (e.g., Calico or Flannel)**:
    ```shell
    kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml
5. **Join worker nodes to the cluster: On each worker node, use the `kubeadm join` command outputted during the `kubeadm init` step on the control-plane node. It should look something like this:**
    ```shell
    sudo kubeadm join <master-ip>:6443 --token <token> --discovery-token-ca-cert-hash sha256:<hash>
For more detailed instructions, refer to the [Kubernetes official documentation](https://kubernetes.io/docs/home/).

## Build XPU-Scheduler
Build the execuble binary file:
```shell
$ make build
```
## Deploy XPU-Scheduler
Deploy the XPU scheduler to your Kubernetes cluster:
```shell
$ kubectl apply -f deploy/xpu-scheduler.yaml
```

## Test XPU-Scheduler
To validate the functionality of the scheduler, apply the test configurations:
```shell
$ kubectl apply -f deploy/test-scheduler.yaml
```

Then watch sample-scheduler pod logs.
