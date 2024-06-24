// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by kubetype-gen. DO NOT EDIT.

package v1

import (
	networkingv1 "github.com/alibaba/higress/api/networking/v1"
	v1alpha1 "istio.io/api/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// <!-- crd generation tags
// +cue-gen:Http2Rpc:groupName:networking.higress.io
// +cue-gen:Http2Rpc:version:v1
// +cue-gen:Http2Rpc:storageVersion
// +cue-gen:Http2Rpc:annotations:helm.sh/resource-policy=keep
// +cue-gen:Http2Rpc:subresource:status
// +cue-gen:Http2Rpc:scope:Namespaced
// +cue-gen:Http2Rpc:resource:categories=higress-io,plural=http2rpcs
// +cue-gen:Http2Rpc:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.higress.io/v1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type Http2Rpc struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1.Http2Rpc `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Http2RpcList is a collection of Http2Rpcs.
type Http2RpcList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Http2Rpc `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// please upgrade the proto package
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// <!-- crd generation tags
// +cue-gen:McpBridge:groupName:networking.higress.io
// +cue-gen:McpBridge:version:v1
// +cue-gen:McpBridge:storageVersion
// +cue-gen:McpBridge:annotations:helm.sh/resource-policy=keep
// +cue-gen:McpBridge:subresource:status
// +cue-gen:McpBridge:scope:Namespaced
// +cue-gen:McpBridge:resource:categories=higress-io,plural=mcpbridges
// +cue-gen:McpBridge:preserveUnknownFields:false
// -->
//
// <!-- go code generation tags
// +kubetype-gen
// +kubetype-gen:groupVersion=networking.higress.io/v1
// +genclient
// +k8s:deepcopy-gen=true
// -->
type McpBridge struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec networkingv1.McpBridge `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status v1alpha1.IstioStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// McpBridgeList is a collection of McpBridges.
type McpBridgeList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []McpBridge `json:"items" protobuf:"bytes,2,rep,name=items"`
}
