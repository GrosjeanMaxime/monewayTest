// Copyright 2019 Google LLC.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.12.3
// source: google/devtools/clouderrorreporting/v1beta1/error_group_service.proto

package clouderrorreporting

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// A request to return an individual group.
type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The group resource name. Written as
	// `projects/{projectID}/groups/{group_name}`. Call
	// [`groupStats.list`](https://cloud.google.com/error-reporting/reference/rest/v1beta1/projects.groupStats/list)
	// to return a list of groups belonging to this project.
	//
	// Example: `projects/my-project-123/groups/my-group`
	GroupName string `protobuf:"bytes,1,opt,name=group_name,json=groupName,proto3" json:"group_name,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

// A request to replace the existing data for the given group.
type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. The group which replaces the resource on the server.
	Group *ErrorGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGroupRequest) GetGroup() *ErrorGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

var File_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto protoreflect.FileDescriptor

var file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDesc = []byte{
	0x0a, 0x45, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65,
	0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x67, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x54, 0x0a, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x2f,
	0x0a, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x68, 0x0a, 0x12, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x52, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x32, 0xfb, 0x03, 0x0a, 0x11, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc1, 0x01, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x3e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a,
	0x7d, 0xda, 0x41, 0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0xc9,
	0x01, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x3f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x37, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c,
	0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x40, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32,
	0x1a, 0x29, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x7b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f,
	0x2a, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2f, 0x2a, 0x7d, 0x3a, 0x05, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0xda, 0x41, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x56, 0xca, 0x41, 0x22, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0xd2, 0x41, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x77, 0x77, 0x77, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f,
	0x72, 0x6d, 0x42, 0xa3, 0x02, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x16, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x5e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x6f, 0x6f, 0x6c, 0x73,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67,
	0xf8, 0x01, 0x01, 0xaa, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x56, 0x31, 0x42, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x23, 0x47, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x5c, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xea,
	0x02, 0x26, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a,
	0x3a, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescOnce sync.Once
	file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescData = file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDesc
)

func file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescGZIP() []byte {
	file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescOnce.Do(func() {
		file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescData)
	})
	return file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDescData
}

var file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_goTypes = []interface{}{
	(*GetGroupRequest)(nil),    // 0: google.devtools.clouderrorreporting.v1beta1.GetGroupRequest
	(*UpdateGroupRequest)(nil), // 1: google.devtools.clouderrorreporting.v1beta1.UpdateGroupRequest
	(*ErrorGroup)(nil),         // 2: google.devtools.clouderrorreporting.v1beta1.ErrorGroup
}
var file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_depIdxs = []int32{
	2, // 0: google.devtools.clouderrorreporting.v1beta1.UpdateGroupRequest.group:type_name -> google.devtools.clouderrorreporting.v1beta1.ErrorGroup
	0, // 1: google.devtools.clouderrorreporting.v1beta1.ErrorGroupService.GetGroup:input_type -> google.devtools.clouderrorreporting.v1beta1.GetGroupRequest
	1, // 2: google.devtools.clouderrorreporting.v1beta1.ErrorGroupService.UpdateGroup:input_type -> google.devtools.clouderrorreporting.v1beta1.UpdateGroupRequest
	2, // 3: google.devtools.clouderrorreporting.v1beta1.ErrorGroupService.GetGroup:output_type -> google.devtools.clouderrorreporting.v1beta1.ErrorGroup
	2, // 4: google.devtools.clouderrorreporting.v1beta1.ErrorGroupService.UpdateGroup:output_type -> google.devtools.clouderrorreporting.v1beta1.ErrorGroup
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_init() }
func file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_init() {
	if File_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto != nil {
		return
	}
	file_google_devtools_clouderrorreporting_v1beta1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_goTypes,
		DependencyIndexes: file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_depIdxs,
		MessageInfos:      file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_msgTypes,
	}.Build()
	File_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto = out.File
	file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_rawDesc = nil
	file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_goTypes = nil
	file_google_devtools_clouderrorreporting_v1beta1_error_group_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ErrorGroupServiceClient is the client API for ErrorGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ErrorGroupServiceClient interface {
	// Get the specified group.
	GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error)
	// Replace the data for the specified group.
	// Fails if the group does not exist.
	UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error)
}

type errorGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewErrorGroupServiceClient(cc grpc.ClientConnInterface) ErrorGroupServiceClient {
	return &errorGroupServiceClient{cc}
}

func (c *errorGroupServiceClient) GetGroup(ctx context.Context, in *GetGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error) {
	out := new(ErrorGroup)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *errorGroupServiceClient) UpdateGroup(ctx context.Context, in *UpdateGroupRequest, opts ...grpc.CallOption) (*ErrorGroup, error) {
	out := new(ErrorGroup)
	err := c.cc.Invoke(ctx, "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ErrorGroupServiceServer is the server API for ErrorGroupService service.
type ErrorGroupServiceServer interface {
	// Get the specified group.
	GetGroup(context.Context, *GetGroupRequest) (*ErrorGroup, error)
	// Replace the data for the specified group.
	// Fails if the group does not exist.
	UpdateGroup(context.Context, *UpdateGroupRequest) (*ErrorGroup, error)
}

// UnimplementedErrorGroupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedErrorGroupServiceServer struct {
}

func (*UnimplementedErrorGroupServiceServer) GetGroup(context.Context, *GetGroupRequest) (*ErrorGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (*UnimplementedErrorGroupServiceServer) UpdateGroup(context.Context, *UpdateGroupRequest) (*ErrorGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}

func RegisterErrorGroupServiceServer(s *grpc.Server, srv ErrorGroupServiceServer) {
	s.RegisterService(&_ErrorGroupService_serviceDesc, srv)
}

func _ErrorGroupService_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorGroupServiceServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorGroupServiceServer).GetGroup(ctx, req.(*GetGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ErrorGroupService_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorGroupServiceServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.devtools.clouderrorreporting.v1beta1.ErrorGroupService/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorGroupServiceServer).UpdateGroup(ctx, req.(*UpdateGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ErrorGroupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.devtools.clouderrorreporting.v1beta1.ErrorGroupService",
	HandlerType: (*ErrorGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGroup",
			Handler:    _ErrorGroupService_GetGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _ErrorGroupService_UpdateGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/devtools/clouderrorreporting/v1beta1/error_group_service.proto",
}
