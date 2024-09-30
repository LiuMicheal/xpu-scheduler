/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// MinExtenderPriority defines the min priority value for extender.
	MinExtenderPriority int64 = 0

	// MaxExtenderPriority defines the max priority value for extender.
	MaxExtenderPriority int64 = 10
)

// ExtenderPreemptionResult represents the result returned by preemption phase of extender.
type ExtenderPreemptionResult struct {
	NodeNameToMetaVictims map[string]*MetaVictims
}

// ExtenderPreemptionArgs represents the arguments needed by the extender to preempt pods on nodes.
type ExtenderPreemptionArgs struct {
	// Pod being scheduled
	Pod *v1.Pod
	// Victims map generated by scheduler preemption phase
	// Only set NodeNameToMetaVictims if Extender.NodeCacheCapable == true. Otherwise, only set NodeNameToVictims.
	NodeNameToVictims     map[string]*Victims
	NodeNameToMetaVictims map[string]*MetaVictims
}

// Victims represents:
//   pods:  a group of pods expected to be preempted.
//   numPDBViolations: the count of violations of PodDisruptionBudget
type Victims struct {
	Pods             []*v1.Pod
	NumPDBViolations int64
}

// MetaPod represent identifier for a v1.Pod
type MetaPod struct {
	UID string
}

// MetaVictims represents:
//   pods:  a group of pods expected to be preempted.
//     Only Pod identifiers will be sent and user are expect to get v1.Pod in their own way.
//   numPDBViolations: the count of violations of PodDisruptionBudget
type MetaVictims struct {
	Pods             []*MetaPod
	NumPDBViolations int64
}

// ExtenderArgs represents the arguments needed by the extender to filter/prioritize
// nodes for a pod.
type ExtenderArgs struct {
	// Pod being scheduled
	Pod *v1.Pod
	// List of candidate nodes where the pod can be scheduled; to be populated
	// only if Extender.NodeCacheCapable == false
	Nodes *v1.NodeList
	// List of candidate node names where the pod can be scheduled; to be
	// populated only if Extender.NodeCacheCapable == true
	NodeNames *[]string
}

// FailedNodesMap represents the filtered out nodes, with node names and failure messages
type FailedNodesMap map[string]string

// ExtenderFilterResult represents the results of a filter call to an extender
type ExtenderFilterResult struct {
	// Filtered set of nodes where the pod can be scheduled; to be populated
	// only if Extender.NodeCacheCapable == false
	Nodes *v1.NodeList
	// Filtered set of nodes where the pod can be scheduled; to be populated
	// only if Extender.NodeCacheCapable == true
	NodeNames *[]string
	// Filtered out nodes where the pod can't be scheduled and the failure messages
	FailedNodes FailedNodesMap
	// Filtered out nodes where the pod can't be scheduled and preemption would
	// not change anything. The value is the failure message same as FailedNodes.
	// Nodes specified here takes precedence over FailedNodes.
	FailedAndUnresolvableNodes FailedNodesMap
	// Error message indicating failure
	Error string
}

// ExtenderBindingArgs represents the arguments to an extender for binding a pod to a node.
type ExtenderBindingArgs struct {
	// PodName is the name of the pod being bound
	PodName string
	// PodNamespace is the namespace of the pod being bound
	PodNamespace string
	// PodUID is the UID of the pod being bound
	PodUID types.UID
	// Node selected by the scheduler
	Node string
}

// ExtenderBindingResult represents the result of binding of a pod to a node from an extender.
type ExtenderBindingResult struct {
	// Error message indicating failure
	Error string
}

// HostPriority represents the priority of scheduling to a particular host, higher priority is better.
type HostPriority struct {
	// Name of the host
	Host string
	// Score associated with the host
	Score int64
}

// HostPriorityList declares a []HostPriority type.
type HostPriorityList []HostPriority
