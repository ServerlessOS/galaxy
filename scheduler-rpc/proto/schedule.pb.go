// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: proto/schedule.proto

package scheduler_rpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PeerSchedulerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *PeerSchedulerInfo) Reset() {
	*x = PeerSchedulerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerSchedulerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerSchedulerInfo) ProtoMessage() {}

func (x *PeerSchedulerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerSchedulerInfo.ProtoReflect.Descriptor instead.
func (*PeerSchedulerInfo) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{0}
}

func (x *PeerSchedulerInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *PeerSchedulerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type PeerSchedulersUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List   []*PeerSchedulerInfo `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	Action string               `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
}

func (x *PeerSchedulersUpdate) Reset() {
	*x = PeerSchedulersUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerSchedulersUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerSchedulersUpdate) ProtoMessage() {}

func (x *PeerSchedulersUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerSchedulersUpdate.ProtoReflect.Descriptor instead.
func (*PeerSchedulersUpdate) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{1}
}

func (x *PeerSchedulersUpdate) GetList() []*PeerSchedulerInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *PeerSchedulersUpdate) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type PeerSchedulersUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int64 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *PeerSchedulersUpdateReply) Reset() {
	*x = PeerSchedulersUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerSchedulersUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerSchedulersUpdateReply) ProtoMessage() {}

func (x *PeerSchedulersUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerSchedulersUpdateReply.ProtoReflect.Descriptor instead.
func (*PeerSchedulersUpdateReply) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{2}
}

func (x *PeerSchedulersUpdateReply) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type NodeResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	HaveCpu  int64  `protobuf:"varint,2,opt,name=haveCpu,proto3" json:"haveCpu,omitempty"`
	HaveMem  int64  `protobuf:"varint,3,opt,name=haveMem,proto3" json:"haveMem,omitempty"`
	Address  string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Port     string `protobuf:"bytes,5,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *NodeResource) Reset() {
	*x = NodeResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResource) ProtoMessage() {}

func (x *NodeResource) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResource.ProtoReflect.Descriptor instead.
func (*NodeResource) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{3}
}

func (x *NodeResource) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *NodeResource) GetHaveCpu() int64 {
	if x != nil {
		return x.HaveCpu
	}
	return 0
}

func (x *NodeResource) GetHaveMem() int64 {
	if x != nil {
		return x.HaveMem
	}
	return 0
}

func (x *NodeResource) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NodeResource) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type NodeResourceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List       []*NodeResource `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Action     string          `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
	SourceAddr string          `protobuf:"bytes,3,opt,name=SourceAddr,proto3" json:"SourceAddr,omitempty"`
	TargetAddr string          `protobuf:"bytes,4,opt,name=TargetAddr,proto3" json:"TargetAddr,omitempty"`
}

func (x *NodeResourceUpdate) Reset() {
	*x = NodeResourceUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResourceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResourceUpdate) ProtoMessage() {}

func (x *NodeResourceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResourceUpdate.ProtoReflect.Descriptor instead.
func (*NodeResourceUpdate) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{4}
}

func (x *NodeResourceUpdate) GetList() []*NodeResource {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *NodeResourceUpdate) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *NodeResourceUpdate) GetSourceAddr() string {
	if x != nil {
		return x.SourceAddr
	}
	return ""
}

func (x *NodeResourceUpdate) GetTargetAddr() string {
	if x != nil {
		return x.TargetAddr
	}
	return ""
}

type NodeResourceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int64 `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *NodeResourceReply) Reset() {
	*x = NodeResourceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeResourceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeResourceReply) ProtoMessage() {}

func (x *NodeResourceReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeResourceReply.ProtoReflect.Descriptor instead.
func (*NodeResourceReply) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{5}
}

func (x *NodeResourceReply) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type ScheduleRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*ScheduleRequest `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ScheduleRequestList) Reset() {
	*x = ScheduleRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequestList) ProtoMessage() {}

func (x *ScheduleRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequestList.ProtoReflect.Descriptor instead.
func (*ScheduleRequestList) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{6}
}

func (x *ScheduleRequestList) GetList() []*ScheduleRequest {
	if x != nil {
		return x.List
	}
	return nil
}

type ScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId      int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	FuncName       string `protobuf:"bytes,2,opt,name=funcName,proto3" json:"funcName,omitempty"`
	RequireCpu     int64  `protobuf:"varint,3,opt,name=requireCpu,proto3" json:"requireCpu,omitempty"`
	RequireMem     int64  `protobuf:"varint,4,opt,name=requireMem,proto3" json:"requireMem,omitempty"`
	DispatcherAddr string `protobuf:"bytes,5,opt,name=dispatcherAddr,proto3" json:"dispatcherAddr,omitempty"`
}

func (x *ScheduleRequest) Reset() {
	*x = ScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleRequest) ProtoMessage() {}

func (x *ScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleRequest.ProtoReflect.Descriptor instead.
func (*ScheduleRequest) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{7}
}

func (x *ScheduleRequest) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ScheduleRequest) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *ScheduleRequest) GetRequireCpu() int64 {
	if x != nil {
		return x.RequireCpu
	}
	return 0
}

func (x *ScheduleRequest) GetRequireMem() int64 {
	if x != nil {
		return x.RequireMem
	}
	return 0
}

func (x *ScheduleRequest) GetDispatcherAddr() string {
	if x != nil {
		return x.DispatcherAddr
	}
	return ""
}

type ScheduleReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId      int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	FuncName       string `protobuf:"bytes,2,opt,name=funcName,proto3" json:"funcName,omitempty"`
	DeployPosition string `protobuf:"bytes,3,opt,name=deployPosition,proto3" json:"deployPosition,omitempty"`
}

func (x *ScheduleReply) Reset() {
	*x = ScheduleReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_schedule_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScheduleReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScheduleReply) ProtoMessage() {}

func (x *ScheduleReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_schedule_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScheduleReply.ProtoReflect.Descriptor instead.
func (*ScheduleReply) Descriptor() ([]byte, []int) {
	return file_proto_schedule_proto_rawDescGZIP(), []int{8}
}

func (x *ScheduleReply) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ScheduleReply) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *ScheduleReply) GetDeployPosition() string {
	if x != nil {
		return x.DeployPosition
	}
	return ""
}

var File_proto_schedule_proto protoreflect.FileDescriptor

var file_proto_schedule_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x22, 0x49, 0x0a, 0x11, 0x50, 0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x60, 0x0a, 0x14,
	0x50, 0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x31,
	0x0a, 0x19, 0x50, 0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x8c, 0x01, 0x0a, 0x0c, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x68, 0x61, 0x76, 0x65, 0x43, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x68, 0x61, 0x76, 0x65, 0x43, 0x70, 0x75, 0x12, 0x18, 0x0a, 0x07, 0x68, 0x61, 0x76, 0x65,
	0x4d, 0x65, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x68, 0x61, 0x76, 0x65, 0x4d,
	0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x22, 0x99, 0x01, 0x0a, 0x12, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1e, 0x0a, 0x0a,
	0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x22, 0x29, 0x0a, 0x11,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2e,
	0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0xb3,
	0x01, 0x0a, 0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x70, 0x75, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4d, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4d, 0x65, 0x6d, 0x12, 0x26, 0x0a, 0x0e,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x41, 0x64, 0x64, 0x72, 0x22, 0x71, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xd1, 0x02, 0x0a, 0x09, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x08, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x1e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x54, 0x0a,
	0x13, 0x55, 0x70, 0x61, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72,
	0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x4e, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x0c, 0x50, 0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x18, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x13, 0x50,
	0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x12, 0x1f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x24, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x73, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e,
	0x2f, 0x3b, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_schedule_proto_rawDescOnce sync.Once
	file_proto_schedule_proto_rawDescData = file_proto_schedule_proto_rawDesc
)

func file_proto_schedule_proto_rawDescGZIP() []byte {
	file_proto_schedule_proto_rawDescOnce.Do(func() {
		file_proto_schedule_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_schedule_proto_rawDescData)
	})
	return file_proto_schedule_proto_rawDescData
}

var file_proto_schedule_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_schedule_proto_goTypes = []interface{}{
	(*PeerSchedulerInfo)(nil),         // 0: scheduler.PeerSchedulerInfo
	(*PeerSchedulersUpdate)(nil),      // 1: scheduler.PeerSchedulersUpdate
	(*PeerSchedulersUpdateReply)(nil), // 2: scheduler.PeerSchedulersUpdateReply
	(*NodeResource)(nil),              // 3: scheduler.NodeResource
	(*NodeResourceUpdate)(nil),        // 4: scheduler.NodeResourceUpdate
	(*NodeResourceReply)(nil),         // 5: scheduler.NodeResourceReply
	(*ScheduleRequestList)(nil),       // 6: scheduler.ScheduleRequestList
	(*ScheduleRequest)(nil),           // 7: scheduler.ScheduleRequest
	(*ScheduleReply)(nil),             // 8: scheduler.ScheduleReply
}
var file_proto_schedule_proto_depIdxs = []int32{
	0, // 0: scheduler.PeerSchedulersUpdate.List:type_name -> scheduler.PeerSchedulerInfo
	3, // 1: scheduler.NodeResourceUpdate.list:type_name -> scheduler.NodeResource
	7, // 2: scheduler.ScheduleRequestList.list:type_name -> scheduler.ScheduleRequest
	6, // 3: scheduler.Scheduler.Schedule:input_type -> scheduler.ScheduleRequestList
	4, // 4: scheduler.Scheduler.UpadateNodeResource:input_type -> scheduler.NodeResourceUpdate
	7, // 5: scheduler.Scheduler.PeerSchedule:input_type -> scheduler.ScheduleRequest
	1, // 6: scheduler.Scheduler.PeerSchedulerUpdate:input_type -> scheduler.PeerSchedulersUpdate
	8, // 7: scheduler.Scheduler.Schedule:output_type -> scheduler.ScheduleReply
	5, // 8: scheduler.Scheduler.UpadateNodeResource:output_type -> scheduler.NodeResourceReply
	8, // 9: scheduler.Scheduler.PeerSchedule:output_type -> scheduler.ScheduleReply
	2, // 10: scheduler.Scheduler.PeerSchedulerUpdate:output_type -> scheduler.PeerSchedulersUpdateReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_schedule_proto_init() }
func file_proto_schedule_proto_init() {
	if File_proto_schedule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_schedule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerSchedulerInfo); i {
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
		file_proto_schedule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerSchedulersUpdate); i {
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
		file_proto_schedule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerSchedulersUpdateReply); i {
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
		file_proto_schedule_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResource); i {
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
		file_proto_schedule_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResourceUpdate); i {
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
		file_proto_schedule_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeResourceReply); i {
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
		file_proto_schedule_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequestList); i {
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
		file_proto_schedule_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleRequest); i {
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
		file_proto_schedule_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScheduleReply); i {
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
			RawDescriptor: file_proto_schedule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_schedule_proto_goTypes,
		DependencyIndexes: file_proto_schedule_proto_depIdxs,
		MessageInfos:      file_proto_schedule_proto_msgTypes,
	}.Build()
	File_proto_schedule_proto = out.File
	file_proto_schedule_proto_rawDesc = nil
	file_proto_schedule_proto_goTypes = nil
	file_proto_schedule_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SchedulerClient is the client API for Scheduler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SchedulerClient interface {
	Schedule(ctx context.Context, in *ScheduleRequestList, opts ...grpc.CallOption) (*ScheduleReply, error)
	UpadateNodeResource(ctx context.Context, in *NodeResourceUpdate, opts ...grpc.CallOption) (*NodeResourceReply, error)
	PeerSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleReply, error)
	PeerSchedulerUpdate(ctx context.Context, in *PeerSchedulersUpdate, opts ...grpc.CallOption) (*PeerSchedulersUpdateReply, error)
}

type schedulerClient struct {
	cc grpc.ClientConnInterface
}

func NewSchedulerClient(cc grpc.ClientConnInterface) SchedulerClient {
	return &schedulerClient{cc}
}

func (c *schedulerClient) Schedule(ctx context.Context, in *ScheduleRequestList, opts ...grpc.CallOption) (*ScheduleReply, error) {
	out := new(ScheduleReply)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/Schedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) UpadateNodeResource(ctx context.Context, in *NodeResourceUpdate, opts ...grpc.CallOption) (*NodeResourceReply, error) {
	out := new(NodeResourceReply)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/UpadateNodeResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) PeerSchedule(ctx context.Context, in *ScheduleRequest, opts ...grpc.CallOption) (*ScheduleReply, error) {
	out := new(ScheduleReply)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/PeerSchedule", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *schedulerClient) PeerSchedulerUpdate(ctx context.Context, in *PeerSchedulersUpdate, opts ...grpc.CallOption) (*PeerSchedulersUpdateReply, error) {
	out := new(PeerSchedulersUpdateReply)
	err := c.cc.Invoke(ctx, "/scheduler.Scheduler/PeerSchedulerUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SchedulerServer is the server API for Scheduler service.
type SchedulerServer interface {
	Schedule(context.Context, *ScheduleRequestList) (*ScheduleReply, error)
	UpadateNodeResource(context.Context, *NodeResourceUpdate) (*NodeResourceReply, error)
	PeerSchedule(context.Context, *ScheduleRequest) (*ScheduleReply, error)
	PeerSchedulerUpdate(context.Context, *PeerSchedulersUpdate) (*PeerSchedulersUpdateReply, error)
}

// UnimplementedSchedulerServer can be embedded to have forward compatible implementations.
type UnimplementedSchedulerServer struct {
}

func (*UnimplementedSchedulerServer) Schedule(context.Context, *ScheduleRequestList) (*ScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Schedule not implemented")
}
func (*UnimplementedSchedulerServer) UpadateNodeResource(context.Context, *NodeResourceUpdate) (*NodeResourceReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpadateNodeResource not implemented")
}
func (*UnimplementedSchedulerServer) PeerSchedule(context.Context, *ScheduleRequest) (*ScheduleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeerSchedule not implemented")
}
func (*UnimplementedSchedulerServer) PeerSchedulerUpdate(context.Context, *PeerSchedulersUpdate) (*PeerSchedulersUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PeerSchedulerUpdate not implemented")
}

func RegisterSchedulerServer(s *grpc.Server, srv SchedulerServer) {
	s.RegisterService(&_Scheduler_serviceDesc, srv)
}

func _Scheduler_Schedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequestList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).Schedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/Schedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).Schedule(ctx, req.(*ScheduleRequestList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_UpadateNodeResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeResourceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).UpadateNodeResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/UpadateNodeResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).UpadateNodeResource(ctx, req.(*NodeResourceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_PeerSchedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ScheduleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).PeerSchedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/PeerSchedule",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).PeerSchedule(ctx, req.(*ScheduleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scheduler_PeerSchedulerUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PeerSchedulersUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SchedulerServer).PeerSchedulerUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scheduler.Scheduler/PeerSchedulerUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SchedulerServer).PeerSchedulerUpdate(ctx, req.(*PeerSchedulersUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scheduler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scheduler.Scheduler",
	HandlerType: (*SchedulerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Schedule",
			Handler:    _Scheduler_Schedule_Handler,
		},
		{
			MethodName: "UpadateNodeResource",
			Handler:    _Scheduler_UpadateNodeResource_Handler,
		},
		{
			MethodName: "PeerSchedule",
			Handler:    _Scheduler_PeerSchedule_Handler,
		},
		{
			MethodName: "PeerSchedulerUpdate",
			Handler:    _Scheduler_PeerSchedulerUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/schedule.proto",
}