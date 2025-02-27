//
//Copyright 2020-present Open Networking Foundation.
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// (-- api-linter: core::0191::java-outer-classname=disabled
//     aip.dev/not-precedent: We need to do this because java clients should use onos classic. --)
// (-- api-linter: core::0191::java-package=disabled --)
// (-- api-linter: core::0191::java-multiple-files=disabled --)
//api-linter -I /Users/renner/go/src/github.com/p4lang/p4runtime/proto -I /Users/renner/go/src/github.com/googleapis/api-common-protos pipelines.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.13.0
// source: onos/control/pipelines.proto

package control

import (
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	v1 "github.com/p4lang/p4runtime/go/p4/config/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

//Pipeline resource
type Pipeline struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Pipeline name in the format of "arch/*/pipelines/*".
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User assigned metadata key,values
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// p4info compiler artifact describing the pipeline contents
	P4Info *v1.P4Info `protobuf:"bytes,3,opt,name=p4info,proto3" json:"p4info,omitempty"`
	// target specific binary configuration, could just be json in bmv2 case but all treated same as opaque []byte
	Config []byte `protobuf:"bytes,4,opt,name=config,proto3" json:"config,omitempty"`
	// this is usually generated by the registry surface by pulling the param ids from p4info
	// purpose is for efficient registration space.
	// i.e. for the various roles which want to register particular entity IDs
	Surface []*v1.P4Ids `protobuf:"bytes,5,rep,name=surface,proto3" json:"surface,omitempty"`
	// role registration for this pipeline resource
	Roles map[uint64]*RoleConfig `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Pipeline) Reset() {
	*x = Pipeline{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onos_control_pipelines_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pipeline) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pipeline) ProtoMessage() {}

func (x *Pipeline) ProtoReflect() protoreflect.Message {
	mi := &file_onos_control_pipelines_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pipeline.ProtoReflect.Descriptor instead.
func (*Pipeline) Descriptor() ([]byte, []int) {
	return file_onos_control_pipelines_proto_rawDescGZIP(), []int{0}
}

func (x *Pipeline) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Pipeline) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Pipeline) GetP4Info() *v1.P4Info {
	if x != nil {
		return x.P4Info
	}
	return nil
}

func (x *Pipeline) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Pipeline) GetSurface() []*v1.P4Ids {
	if x != nil {
		return x.Surface
	}
	return nil
}

func (x *Pipeline) GetRoles() map[uint64]*RoleConfig {
	if x != nil {
		return x.Roles
	}
	return nil
}

type RoleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	P4Objects                        *P4Objects `protobuf:"bytes,1,opt,name=p4_objects,json=p4Objects,proto3" json:"p4_objects,omitempty"`
	AllowPacketOut                   bool       `protobuf:"varint,2,opt,name=allow_packet_out,json=allowPacketOut,proto3" json:"allow_packet_out,omitempty"`
	AllowPacketIn                    bool       `protobuf:"varint,3,opt,name=allow_packet_in,json=allowPacketIn,proto3" json:"allow_packet_in,omitempty"`
	AllowSetForwardingPipelineConfig bool       `protobuf:"varint,4,opt,name=allow_set_forwarding_pipeline_config,json=allowSetForwardingPipelineConfig,proto3" json:"allow_set_forwarding_pipeline_config,omitempty"`
	Extension                        *any.Any   `protobuf:"bytes,5,opt,name=extension,proto3" json:"extension,omitempty"`
}

func (x *RoleConfig) Reset() {
	*x = RoleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onos_control_pipelines_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleConfig) ProtoMessage() {}

func (x *RoleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_onos_control_pipelines_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleConfig.ProtoReflect.Descriptor instead.
func (*RoleConfig) Descriptor() ([]byte, []int) {
	return file_onos_control_pipelines_proto_rawDescGZIP(), []int{1}
}

func (x *RoleConfig) GetP4Objects() *P4Objects {
	if x != nil {
		return x.P4Objects
	}
	return nil
}

func (x *RoleConfig) GetAllowPacketOut() bool {
	if x != nil {
		return x.AllowPacketOut
	}
	return false
}

func (x *RoleConfig) GetAllowPacketIn() bool {
	if x != nil {
		return x.AllowPacketIn
	}
	return false
}

func (x *RoleConfig) GetAllowSetForwardingPipelineConfig() bool {
	if x != nil {
		return x.AllowSetForwardingPipelineConfig
	}
	return false
}

func (x *RoleConfig) GetExtension() *any.Any {
	if x != nil {
		return x.Extension
	}
	return nil
}

type P4Objects struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternIds        []uint32 `protobuf:"varint,1,rep,packed,name=extern_ids,json=externIds,proto3" json:"extern_ids,omitempty"`
	TableIds         []uint32 `protobuf:"varint,2,rep,packed,name=table_ids,json=tableIds,proto3" json:"table_ids,omitempty"`
	ActionIds        []uint32 `protobuf:"varint,3,rep,packed,name=action_ids,json=actionIds,proto3" json:"action_ids,omitempty"`
	ActionProfileIds []uint32 `protobuf:"varint,4,rep,packed,name=action_profile_ids,json=actionProfileIds,proto3" json:"action_profile_ids,omitempty"`
	CounterIds       []uint32 `protobuf:"varint,5,rep,packed,name=counter_ids,json=counterIds,proto3" json:"counter_ids,omitempty"` // indirect counters
	MeterIds         []uint32 `protobuf:"varint,6,rep,packed,name=meter_ids,json=meterIds,proto3" json:"meter_ids,omitempty"`       // indirect meters
}

func (x *P4Objects) Reset() {
	*x = P4Objects{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onos_control_pipelines_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *P4Objects) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*P4Objects) ProtoMessage() {}

func (x *P4Objects) ProtoReflect() protoreflect.Message {
	mi := &file_onos_control_pipelines_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use P4Objects.ProtoReflect.Descriptor instead.
func (*P4Objects) Descriptor() ([]byte, []int) {
	return file_onos_control_pipelines_proto_rawDescGZIP(), []int{2}
}

func (x *P4Objects) GetExternIds() []uint32 {
	if x != nil {
		return x.ExternIds
	}
	return nil
}

func (x *P4Objects) GetTableIds() []uint32 {
	if x != nil {
		return x.TableIds
	}
	return nil
}

func (x *P4Objects) GetActionIds() []uint32 {
	if x != nil {
		return x.ActionIds
	}
	return nil
}

func (x *P4Objects) GetActionProfileIds() []uint32 {
	if x != nil {
		return x.ActionProfileIds
	}
	return nil
}

func (x *P4Objects) GetCounterIds() []uint32 {
	if x != nil {
		return x.CounterIds
	}
	return nil
}

func (x *P4Objects) GetMeterIds() []uint32 {
	if x != nil {
		return x.MeterIds
	}
	return nil
}

//CreatePipelineRequest wraps main resource pipeline with namespace parent on create
type CreatePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The namespace in which the pipeline should be created.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// The pipeline object being created
	Pipeline *Pipeline `protobuf:"bytes,2,opt,name=pipeline,proto3" json:"pipeline,omitempty"`
}

func (x *CreatePipelineRequest) Reset() {
	*x = CreatePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onos_control_pipelines_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePipelineRequest) ProtoMessage() {}

func (x *CreatePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onos_control_pipelines_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePipelineRequest.ProtoReflect.Descriptor instead.
func (*CreatePipelineRequest) Descriptor() ([]byte, []int) {
	return file_onos_control_pipelines_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePipelineRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *CreatePipelineRequest) GetPipeline() *Pipeline {
	if x != nil {
		return x.Pipeline
	}
	return nil
}

//DeletePipelineRequest ...
type DeletePipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Pipeline name in the format of "arch/*/pipelines/*" to delete
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeletePipelineRequest) Reset() {
	*x = DeletePipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onos_control_pipelines_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePipelineRequest) ProtoMessage() {}

func (x *DeletePipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onos_control_pipelines_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePipelineRequest.ProtoReflect.Descriptor instead.
func (*DeletePipelineRequest) Descriptor() ([]byte, []int) {
	return file_onos_control_pipelines_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//GetPipelineRequest ...
type GetPipelineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Pipeline name in the format of "arch/*/pipelines/*" to retrieve
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPipelineRequest) Reset() {
	*x = GetPipelineRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_onos_control_pipelines_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPipelineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPipelineRequest) ProtoMessage() {}

func (x *GetPipelineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_onos_control_pipelines_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPipelineRequest.ProtoReflect.Descriptor instead.
func (*GetPipelineRequest) Descriptor() ([]byte, []int) {
	return file_onos_control_pipelines_proto_rawDescGZIP(), []int{5}
}

func (x *GetPipelineRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_onos_control_pipelines_proto protoreflect.FileDescriptor

var file_onos_control_pipelines_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x6f, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2f, 0x70,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f,
	0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x70, 0x34, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x34, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x03, 0x0a, 0x08, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69,
	0x6e, 0x65, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x6e,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x06, 0x70, 0x34,
	0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x34, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x34, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x34, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x03, 0xe0,
	0x41, 0x02, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x07, 0x73, 0x75,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x34,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x34, 0x49, 0x64, 0x73,
	0x52, 0x07, 0x73, 0x75, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x55, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x31, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0x47, 0xea, 0x41, 0x44, 0x0a, 0x20, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20,
	0x61, 0x72, 0x63, 0x68, 0x2f, 0x7b, 0x61, 0x72, 0x63, 0x68, 0x7d, 0x2f, 0x70, 0x69, 0x70, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x7d,
	0x22, 0x9d, 0x02, 0x0a, 0x0a, 0x52, 0x6f, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x39, 0x0a, 0x0a, 0x70, 0x34, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x34, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x52,
	0x09, 0x70, 0x34, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x10, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x65,
	0x74, 0x4f, 0x75, 0x74, 0x12, 0x26, 0x0a, 0x0f, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x70, 0x61,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61,
	0x6c, 0x6c, 0x6f, 0x77, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x12, 0x4e, 0x0a, 0x24,
	0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x20, 0x61, 0x6c, 0x6c, 0x6f,
	0x77, 0x53, 0x65, 0x74, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x0a, 0x09,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x22, 0xd2, 0x01, 0x0a, 0x09, 0x50, 0x34, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x09, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d,
	0x52, 0x08, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x09,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x10, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x03, 0xe0, 0x41, 0x02, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x5f, 0x0a, 0x08,
	0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x42, 0x28, 0xe0, 0x41, 0x02, 0xfa, 0x41,
	0x22, 0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x52, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x30, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x2d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x03, 0xe0, 0x41, 0x02, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xa8,
	0x03, 0x0a, 0x10, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x12, 0x96, 0x01, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x69,
	0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x26, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x41, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x61,
	0x72, 0x63, 0x68, 0x2f, 0x2a, 0x7d, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73,
	0x3a, 0x08, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0xda, 0x41, 0x0f, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x2c, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x7e, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x26,
	0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x2a, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x61, 0x72, 0x63, 0x68, 0x2f, 0x2a, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e,
	0x65, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x7b, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x6e,
	0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6f, 0x6e, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x1f, 0x12, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x61,
	0x72, 0x63, 0x68, 0x2f, 0x2a, 0x2f, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x2f,
	0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2f, 0x6f, 0x6e, 0x6f, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x6f, 0x6e, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_onos_control_pipelines_proto_rawDescOnce sync.Once
	file_onos_control_pipelines_proto_rawDescData = file_onos_control_pipelines_proto_rawDesc
)

func file_onos_control_pipelines_proto_rawDescGZIP() []byte {
	file_onos_control_pipelines_proto_rawDescOnce.Do(func() {
		file_onos_control_pipelines_proto_rawDescData = protoimpl.X.CompressGZIP(file_onos_control_pipelines_proto_rawDescData)
	})
	return file_onos_control_pipelines_proto_rawDescData
}

var file_onos_control_pipelines_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_onos_control_pipelines_proto_goTypes = []interface{}{
	(*Pipeline)(nil),              // 0: onos.control.v1.Pipeline
	(*RoleConfig)(nil),            // 1: onos.control.v1.RoleConfig
	(*P4Objects)(nil),             // 2: onos.control.v1.P4Objects
	(*CreatePipelineRequest)(nil), // 3: onos.control.v1.CreatePipelineRequest
	(*DeletePipelineRequest)(nil), // 4: onos.control.v1.DeletePipelineRequest
	(*GetPipelineRequest)(nil),    // 5: onos.control.v1.GetPipelineRequest
	nil,                           // 6: onos.control.v1.Pipeline.LabelsEntry
	nil,                           // 7: onos.control.v1.Pipeline.RolesEntry
	(*v1.P4Info)(nil),             // 8: p4.config.v1.P4Info
	(*v1.P4Ids)(nil),              // 9: p4.config.v1.P4Ids
	(*any.Any)(nil),               // 10: google.protobuf.Any
	(*empty.Empty)(nil),           // 11: google.protobuf.Empty
}
var file_onos_control_pipelines_proto_depIdxs = []int32{
	6,  // 0: onos.control.v1.Pipeline.labels:type_name -> onos.control.v1.Pipeline.LabelsEntry
	8,  // 1: onos.control.v1.Pipeline.p4info:type_name -> p4.config.v1.P4Info
	9,  // 2: onos.control.v1.Pipeline.surface:type_name -> p4.config.v1.P4Ids
	7,  // 3: onos.control.v1.Pipeline.roles:type_name -> onos.control.v1.Pipeline.RolesEntry
	2,  // 4: onos.control.v1.RoleConfig.p4_objects:type_name -> onos.control.v1.P4Objects
	10, // 5: onos.control.v1.RoleConfig.extension:type_name -> google.protobuf.Any
	0,  // 6: onos.control.v1.CreatePipelineRequest.pipeline:type_name -> onos.control.v1.Pipeline
	1,  // 7: onos.control.v1.Pipeline.RolesEntry.value:type_name -> onos.control.v1.RoleConfig
	3,  // 8: onos.control.v1.PipelineRegistry.CreatePipeline:input_type -> onos.control.v1.CreatePipelineRequest
	4,  // 9: onos.control.v1.PipelineRegistry.DeletePipeline:input_type -> onos.control.v1.DeletePipelineRequest
	5,  // 10: onos.control.v1.PipelineRegistry.GetPipeline:input_type -> onos.control.v1.GetPipelineRequest
	0,  // 11: onos.control.v1.PipelineRegistry.CreatePipeline:output_type -> onos.control.v1.Pipeline
	11, // 12: onos.control.v1.PipelineRegistry.DeletePipeline:output_type -> google.protobuf.Empty
	0,  // 13: onos.control.v1.PipelineRegistry.GetPipeline:output_type -> onos.control.v1.Pipeline
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_onos_control_pipelines_proto_init() }
func file_onos_control_pipelines_proto_init() {
	if File_onos_control_pipelines_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_onos_control_pipelines_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pipeline); i {
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
		file_onos_control_pipelines_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleConfig); i {
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
		file_onos_control_pipelines_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*P4Objects); i {
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
		file_onos_control_pipelines_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePipelineRequest); i {
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
		file_onos_control_pipelines_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePipelineRequest); i {
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
		file_onos_control_pipelines_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPipelineRequest); i {
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
			RawDescriptor: file_onos_control_pipelines_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_onos_control_pipelines_proto_goTypes,
		DependencyIndexes: file_onos_control_pipelines_proto_depIdxs,
		MessageInfos:      file_onos_control_pipelines_proto_msgTypes,
	}.Build()
	File_onos_control_pipelines_proto = out.File
	file_onos_control_pipelines_proto_rawDesc = nil
	file_onos_control_pipelines_proto_goTypes = nil
	file_onos_control_pipelines_proto_depIdxs = nil
}
