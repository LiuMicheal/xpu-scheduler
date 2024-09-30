/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "k8s.io/api/admissionregistration/v1beta1"
)

// RuleApplyConfiguration represents an declarative configuration of the Rule type for use
// with apply.
type RuleApplyConfiguration struct {
	APIGroups   []string           `json:"apiGroups,omitempty"`
	APIVersions []string           `json:"apiVersions,omitempty"`
	Resources   []string           `json:"resources,omitempty"`
	Scope       *v1beta1.ScopeType `json:"scope,omitempty"`
}

// RuleApplyConfiguration constructs an declarative configuration of the Rule type for use with
// apply.
func Rule() *RuleApplyConfiguration {
	return &RuleApplyConfiguration{}
}

// WithAPIGroups adds the given value to the APIGroups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIGroups field.
func (b *RuleApplyConfiguration) WithAPIGroups(values ...string) *RuleApplyConfiguration {
	for i := range values {
		b.APIGroups = append(b.APIGroups, values[i])
	}
	return b
}

// WithAPIVersions adds the given value to the APIVersions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the APIVersions field.
func (b *RuleApplyConfiguration) WithAPIVersions(values ...string) *RuleApplyConfiguration {
	for i := range values {
		b.APIVersions = append(b.APIVersions, values[i])
	}
	return b
}

// WithResources adds the given value to the Resources field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Resources field.
func (b *RuleApplyConfiguration) WithResources(values ...string) *RuleApplyConfiguration {
	for i := range values {
		b.Resources = append(b.Resources, values[i])
	}
	return b
}

// WithScope sets the Scope field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Scope field is set to the value of the last call.
func (b *RuleApplyConfiguration) WithScope(value v1beta1.ScopeType) *RuleApplyConfiguration {
	b.Scope = &value
	return b
}
