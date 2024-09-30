//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiserver "k8s.io/apiserver/pkg/apis/apiserver"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AdmissionConfiguration)(nil), (*apiserver.AdmissionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(a.(*AdmissionConfiguration), b.(*apiserver.AdmissionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.AdmissionConfiguration)(nil), (*AdmissionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_AdmissionConfiguration_To_v1_AdmissionConfiguration(a.(*apiserver.AdmissionConfiguration), b.(*AdmissionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*AdmissionPluginConfiguration)(nil), (*apiserver.AdmissionPluginConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(a.(*AdmissionPluginConfiguration), b.(*apiserver.AdmissionPluginConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*apiserver.AdmissionPluginConfiguration)(nil), (*AdmissionPluginConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_apiserver_AdmissionPluginConfiguration_To_v1_AdmissionPluginConfiguration(a.(*apiserver.AdmissionPluginConfiguration), b.(*AdmissionPluginConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in *AdmissionConfiguration, out *apiserver.AdmissionConfiguration, s conversion.Scope) error {
	out.Plugins = *(*[]apiserver.AdmissionPluginConfiguration)(unsafe.Pointer(&in.Plugins))
	return nil
}

// Convert_v1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration is an autogenerated conversion function.
func Convert_v1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in *AdmissionConfiguration, out *apiserver.AdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_v1_AdmissionConfiguration_To_apiserver_AdmissionConfiguration(in, out, s)
}

func autoConvert_apiserver_AdmissionConfiguration_To_v1_AdmissionConfiguration(in *apiserver.AdmissionConfiguration, out *AdmissionConfiguration, s conversion.Scope) error {
	out.Plugins = *(*[]AdmissionPluginConfiguration)(unsafe.Pointer(&in.Plugins))
	return nil
}

// Convert_apiserver_AdmissionConfiguration_To_v1_AdmissionConfiguration is an autogenerated conversion function.
func Convert_apiserver_AdmissionConfiguration_To_v1_AdmissionConfiguration(in *apiserver.AdmissionConfiguration, out *AdmissionConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_AdmissionConfiguration_To_v1_AdmissionConfiguration(in, out, s)
}

func autoConvert_v1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in *AdmissionPluginConfiguration, out *apiserver.AdmissionPluginConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.Path = in.Path
	out.Configuration = (*runtime.Unknown)(unsafe.Pointer(in.Configuration))
	return nil
}

// Convert_v1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration is an autogenerated conversion function.
func Convert_v1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in *AdmissionPluginConfiguration, out *apiserver.AdmissionPluginConfiguration, s conversion.Scope) error {
	return autoConvert_v1_AdmissionPluginConfiguration_To_apiserver_AdmissionPluginConfiguration(in, out, s)
}

func autoConvert_apiserver_AdmissionPluginConfiguration_To_v1_AdmissionPluginConfiguration(in *apiserver.AdmissionPluginConfiguration, out *AdmissionPluginConfiguration, s conversion.Scope) error {
	out.Name = in.Name
	out.Path = in.Path
	out.Configuration = (*runtime.Unknown)(unsafe.Pointer(in.Configuration))
	return nil
}

// Convert_apiserver_AdmissionPluginConfiguration_To_v1_AdmissionPluginConfiguration is an autogenerated conversion function.
func Convert_apiserver_AdmissionPluginConfiguration_To_v1_AdmissionPluginConfiguration(in *apiserver.AdmissionPluginConfiguration, out *AdmissionPluginConfiguration, s conversion.Scope) error {
	return autoConvert_apiserver_AdmissionPluginConfiguration_To_v1_AdmissionPluginConfiguration(in, out, s)
}
