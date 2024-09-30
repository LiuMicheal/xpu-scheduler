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
	v1beta1 "k8s.io/api/certificates/v1beta1"
)

// CertificateSigningRequestSpecApplyConfiguration represents an declarative configuration of the CertificateSigningRequestSpec type for use
// with apply.
type CertificateSigningRequestSpecApplyConfiguration struct {
	Request           []byte                        `json:"request,omitempty"`
	SignerName        *string                       `json:"signerName,omitempty"`
	ExpirationSeconds *int32                        `json:"expirationSeconds,omitempty"`
	Usages            []v1beta1.KeyUsage            `json:"usages,omitempty"`
	Username          *string                       `json:"username,omitempty"`
	UID               *string                       `json:"uid,omitempty"`
	Groups            []string                      `json:"groups,omitempty"`
	Extra             map[string]v1beta1.ExtraValue `json:"extra,omitempty"`
}

// CertificateSigningRequestSpecApplyConfiguration constructs an declarative configuration of the CertificateSigningRequestSpec type for use with
// apply.
func CertificateSigningRequestSpec() *CertificateSigningRequestSpecApplyConfiguration {
	return &CertificateSigningRequestSpecApplyConfiguration{}
}

// WithRequest adds the given value to the Request field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Request field.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithRequest(values ...byte) *CertificateSigningRequestSpecApplyConfiguration {
	for i := range values {
		b.Request = append(b.Request, values[i])
	}
	return b
}

// WithSignerName sets the SignerName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SignerName field is set to the value of the last call.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithSignerName(value string) *CertificateSigningRequestSpecApplyConfiguration {
	b.SignerName = &value
	return b
}

// WithExpirationSeconds sets the ExpirationSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ExpirationSeconds field is set to the value of the last call.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithExpirationSeconds(value int32) *CertificateSigningRequestSpecApplyConfiguration {
	b.ExpirationSeconds = &value
	return b
}

// WithUsages adds the given value to the Usages field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Usages field.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithUsages(values ...v1beta1.KeyUsage) *CertificateSigningRequestSpecApplyConfiguration {
	for i := range values {
		b.Usages = append(b.Usages, values[i])
	}
	return b
}

// WithUsername sets the Username field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Username field is set to the value of the last call.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithUsername(value string) *CertificateSigningRequestSpecApplyConfiguration {
	b.Username = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithUID(value string) *CertificateSigningRequestSpecApplyConfiguration {
	b.UID = &value
	return b
}

// WithGroups adds the given value to the Groups field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Groups field.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithGroups(values ...string) *CertificateSigningRequestSpecApplyConfiguration {
	for i := range values {
		b.Groups = append(b.Groups, values[i])
	}
	return b
}

// WithExtra puts the entries into the Extra field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Extra field,
// overwriting an existing map entries in Extra field with the same key.
func (b *CertificateSigningRequestSpecApplyConfiguration) WithExtra(entries map[string]v1beta1.ExtraValue) *CertificateSigningRequestSpecApplyConfiguration {
	if b.Extra == nil && len(entries) > 0 {
		b.Extra = make(map[string]v1beta1.ExtraValue, len(entries))
	}
	for k, v := range entries {
		b.Extra[k] = v
	}
	return b
}
