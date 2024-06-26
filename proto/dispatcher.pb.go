// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.11.2
// source: proto/dispatcher.proto

package proto

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

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  int64  `protobuf:"varint,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	FuncName   string `protobuf:"bytes,2,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	RequireCpu int64  `protobuf:"varint,3,opt,name=RequireCpu,proto3" json:"RequireCpu,omitempty"`
	RequireMem int64  `protobuf:"varint,4,opt,name=RequireMem,proto3" json:"RequireMem,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *UserRequest) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *UserRequest) GetRequireCpu() int64 {
	if x != nil {
		return x.RequireCpu
	}
	return 0
}

func (x *UserRequest) GetRequireMem() int64 {
	if x != nil {
		return x.RequireMem
	}
	return 0
}

type UserRequestReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId   int64  `protobuf:"varint,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	FuncName    string `protobuf:"bytes,2,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	Destination string `protobuf:"bytes,3,opt,name=Destination,proto3" json:"Destination,omitempty"`
}

func (x *UserRequestReply) Reset() {
	*x = UserRequestReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequestReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequestReply) ProtoMessage() {}

func (x *UserRequestReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequestReply.ProtoReflect.Descriptor instead.
func (*UserRequestReply) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{1}
}

func (x *UserRequestReply) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *UserRequestReply) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *UserRequestReply) GetDestination() string {
	if x != nil {
		return x.Destination
	}
	return ""
}

type UserRequestList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*UserRequest `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *UserRequestList) Reset() {
	*x = UserRequestList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequestList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequestList) ProtoMessage() {}

func (x *UserRequestList) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequestList.ProtoReflect.Descriptor instead.
func (*UserRequestList) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{2}
}

func (x *UserRequestList) GetList() []*UserRequest {
	if x != nil {
		return x.List
	}
	return nil
}

type InstanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64  `protobuf:"varint,1,opt,name=RequestId,proto3" json:"RequestId,omitempty"`
	FuncName  string `protobuf:"bytes,2,opt,name=FuncName,proto3" json:"FuncName,omitempty"`
	Address   string `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *InstanceInfo) Reset() {
	*x = InstanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceInfo) ProtoMessage() {}

func (x *InstanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceInfo.ProtoReflect.Descriptor instead.
func (*InstanceInfo) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{3}
}

func (x *InstanceInfo) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *InstanceInfo) GetFuncName() string {
	if x != nil {
		return x.FuncName
	}
	return ""
}

func (x *InstanceInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type InstanceUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List   []*InstanceInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Action string          `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
}

func (x *InstanceUpdate) Reset() {
	*x = InstanceUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceUpdate) ProtoMessage() {}

func (x *InstanceUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceUpdate.ProtoReflect.Descriptor instead.
func (*InstanceUpdate) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{4}
}

func (x *InstanceUpdate) GetList() []*InstanceInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *InstanceUpdate) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type InstanceUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int64 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *InstanceUpdateReply) Reset() {
	*x = InstanceUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceUpdateReply) ProtoMessage() {}

func (x *InstanceUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceUpdateReply.ProtoReflect.Descriptor instead.
func (*InstanceUpdateReply) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{5}
}

func (x *InstanceUpdateReply) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

type SchedulerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=NodeName,proto3" json:"NodeName,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (x *SchedulerInfo) Reset() {
	*x = SchedulerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerInfo) ProtoMessage() {}

func (x *SchedulerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerInfo.ProtoReflect.Descriptor instead.
func (*SchedulerInfo) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{6}
}

func (x *SchedulerInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *SchedulerInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type SchedulerViewUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List   []*SchedulerInfo `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	Action string           `protobuf:"bytes,2,opt,name=Action,proto3" json:"Action,omitempty"`
}

func (x *SchedulerViewUpdate) Reset() {
	*x = SchedulerViewUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerViewUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerViewUpdate) ProtoMessage() {}

func (x *SchedulerViewUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerViewUpdate.ProtoReflect.Descriptor instead.
func (*SchedulerViewUpdate) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{7}
}

func (x *SchedulerViewUpdate) GetList() []*SchedulerInfo {
	if x != nil {
		return x.List
	}
	return nil
}

func (x *SchedulerViewUpdate) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

type SchedulerViewUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State int64 `protobuf:"varint,1,opt,name=State,proto3" json:"State,omitempty"`
}

func (x *SchedulerViewUpdateReply) Reset() {
	*x = SchedulerViewUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_dispatcher_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SchedulerViewUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SchedulerViewUpdateReply) ProtoMessage() {}

func (x *SchedulerViewUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_dispatcher_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SchedulerViewUpdateReply.ProtoReflect.Descriptor instead.
func (*SchedulerViewUpdateReply) Descriptor() ([]byte, []int) {
	return file_proto_dispatcher_proto_rawDescGZIP(), []int{8}
}

func (x *SchedulerViewUpdateReply) GetState() int64 {
	if x != nil {
		return x.State
	}
	return 0
}

var File_proto_dispatcher_proto protoreflect.FileDescriptor

var file_proto_dispatcher_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x70, 0x75, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x43, 0x70, 0x75, 0x12, 0x1e,
	0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4d, 0x65, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x4d, 0x65, 0x6d, 0x22, 0x6e,
	0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3e,
	0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x62,
	0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08,
	0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x46, 0x75, 0x6e, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x56, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x18, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6c, 0x69,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2b, 0x0a, 0x13, 0x49, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x45, 0x0a, 0x0d, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x6f, 0x64, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5c,
	0x0a, 0x13, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04,
	0x6c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x30, 0x0a, 0x18,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x32, 0xcd,
	0x02, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x12, 0x47, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1b, 0x2e, 0x64, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x56, 0x69, 0x65, 0x77, 0x12, 0x1a, 0x2e, 0x64,
	0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x63, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x1a, 0x1f, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61,
	0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x5e, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x69,
	0x65, 0x77, 0x12, 0x1f, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x1a, 0x24, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x2e, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x72, 0x56, 0x69, 0x65, 0x77, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x12, 0x17, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c,
	0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x13,
	0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_dispatcher_proto_rawDescOnce sync.Once
	file_proto_dispatcher_proto_rawDescData = file_proto_dispatcher_proto_rawDesc
)

func file_proto_dispatcher_proto_rawDescGZIP() []byte {
	file_proto_dispatcher_proto_rawDescOnce.Do(func() {
		file_proto_dispatcher_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_dispatcher_proto_rawDescData)
	})
	return file_proto_dispatcher_proto_rawDescData
}

var file_proto_dispatcher_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_dispatcher_proto_goTypes = []interface{}{
	(*UserRequest)(nil),              // 0: dispatcher.UserRequest
	(*UserRequestReply)(nil),         // 1: dispatcher.UserRequestReply
	(*UserRequestList)(nil),          // 2: dispatcher.UserRequestList
	(*InstanceInfo)(nil),             // 3: dispatcher.InstanceInfo
	(*InstanceUpdate)(nil),           // 4: dispatcher.InstanceUpdate
	(*InstanceUpdateReply)(nil),      // 5: dispatcher.InstanceUpdateReply
	(*SchedulerInfo)(nil),            // 6: dispatcher.SchedulerInfo
	(*SchedulerViewUpdate)(nil),      // 7: dispatcher.SchedulerViewUpdate
	(*SchedulerViewUpdateReply)(nil), // 8: dispatcher.SchedulerViewUpdateReply
}
var file_proto_dispatcher_proto_depIdxs = []int32{
	0, // 0: dispatcher.UserRequestList.list:type_name -> dispatcher.UserRequest
	3, // 1: dispatcher.InstanceUpdate.list:type_name -> dispatcher.InstanceInfo
	6, // 2: dispatcher.SchedulerViewUpdate.list:type_name -> dispatcher.SchedulerInfo
	2, // 3: dispatcher.dispatcher.Dispatch:input_type -> dispatcher.UserRequestList
	4, // 4: dispatcher.dispatcher.UpdateInstanceView:input_type -> dispatcher.InstanceUpdate
	7, // 5: dispatcher.dispatcher.UpdateSchedulerView:input_type -> dispatcher.SchedulerViewUpdate
	0, // 6: dispatcher.dispatcher.Statis:input_type -> dispatcher.UserRequest
	1, // 7: dispatcher.dispatcher.Dispatch:output_type -> dispatcher.UserRequestReply
	5, // 8: dispatcher.dispatcher.UpdateInstanceView:output_type -> dispatcher.InstanceUpdateReply
	8, // 9: dispatcher.dispatcher.UpdateSchedulerView:output_type -> dispatcher.SchedulerViewUpdateReply
	1, // 10: dispatcher.dispatcher.Statis:output_type -> dispatcher.UserRequestReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_dispatcher_proto_init() }
func file_proto_dispatcher_proto_init() {
	if File_proto_dispatcher_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_dispatcher_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_proto_dispatcher_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequestReply); i {
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
		file_proto_dispatcher_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequestList); i {
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
		file_proto_dispatcher_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceInfo); i {
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
		file_proto_dispatcher_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceUpdate); i {
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
		file_proto_dispatcher_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceUpdateReply); i {
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
		file_proto_dispatcher_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerInfo); i {
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
		file_proto_dispatcher_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerViewUpdate); i {
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
		file_proto_dispatcher_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SchedulerViewUpdateReply); i {
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
			RawDescriptor: file_proto_dispatcher_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_dispatcher_proto_goTypes,
		DependencyIndexes: file_proto_dispatcher_proto_depIdxs,
		MessageInfos:      file_proto_dispatcher_proto_msgTypes,
	}.Build()
	File_proto_dispatcher_proto = out.File
	file_proto_dispatcher_proto_rawDesc = nil
	file_proto_dispatcher_proto_goTypes = nil
	file_proto_dispatcher_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DispatcherClient is the client API for Dispatcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DispatcherClient interface {
	Dispatch(ctx context.Context, in *UserRequestList, opts ...grpc.CallOption) (*UserRequestReply, error)
	UpdateInstanceView(ctx context.Context, in *InstanceUpdate, opts ...grpc.CallOption) (*InstanceUpdateReply, error)
	UpdateSchedulerView(ctx context.Context, in *SchedulerViewUpdate, opts ...grpc.CallOption) (*SchedulerViewUpdateReply, error)
	Statis(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRequestReply, error)
}

type dispatcherClient struct {
	cc grpc.ClientConnInterface
}

func NewDispatcherClient(cc grpc.ClientConnInterface) DispatcherClient {
	return &dispatcherClient{cc}
}

func (c *dispatcherClient) Dispatch(ctx context.Context, in *UserRequestList, opts ...grpc.CallOption) (*UserRequestReply, error) {
	out := new(UserRequestReply)
	err := c.cc.Invoke(ctx, "/dispatcher.dispatcher/Dispatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) UpdateInstanceView(ctx context.Context, in *InstanceUpdate, opts ...grpc.CallOption) (*InstanceUpdateReply, error) {
	out := new(InstanceUpdateReply)
	err := c.cc.Invoke(ctx, "/dispatcher.dispatcher/UpdateInstanceView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) UpdateSchedulerView(ctx context.Context, in *SchedulerViewUpdate, opts ...grpc.CallOption) (*SchedulerViewUpdateReply, error) {
	out := new(SchedulerViewUpdateReply)
	err := c.cc.Invoke(ctx, "/dispatcher.dispatcher/UpdateSchedulerView", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dispatcherClient) Statis(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserRequestReply, error) {
	out := new(UserRequestReply)
	err := c.cc.Invoke(ctx, "/dispatcher.dispatcher/Statis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DispatcherServer is the server API for Dispatcher service.
type DispatcherServer interface {
	Dispatch(context.Context, *UserRequestList) (*UserRequestReply, error)
	UpdateInstanceView(context.Context, *InstanceUpdate) (*InstanceUpdateReply, error)
	UpdateSchedulerView(context.Context, *SchedulerViewUpdate) (*SchedulerViewUpdateReply, error)
	Statis(context.Context, *UserRequest) (*UserRequestReply, error)
}

// UnimplementedDispatcherServer can be embedded to have forward compatible implementations.
type UnimplementedDispatcherServer struct {
}

func (*UnimplementedDispatcherServer) Dispatch(context.Context, *UserRequestList) (*UserRequestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Dispatch not implemented")
}
func (*UnimplementedDispatcherServer) UpdateInstanceView(context.Context, *InstanceUpdate) (*InstanceUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInstanceView not implemented")
}
func (*UnimplementedDispatcherServer) UpdateSchedulerView(context.Context, *SchedulerViewUpdate) (*SchedulerViewUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSchedulerView not implemented")
}
func (*UnimplementedDispatcherServer) Statis(context.Context, *UserRequest) (*UserRequestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Statis not implemented")
}

func RegisterDispatcherServer(s *grpc.Server, srv DispatcherServer) {
	s.RegisterService(&_Dispatcher_serviceDesc, srv)
}

func _Dispatcher_Dispatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequestList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Dispatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.dispatcher/Dispatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Dispatch(ctx, req.(*UserRequestList))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_UpdateInstanceView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InstanceUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).UpdateInstanceView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.dispatcher/UpdateInstanceView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).UpdateInstanceView(ctx, req.(*InstanceUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_UpdateSchedulerView_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SchedulerViewUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).UpdateSchedulerView(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.dispatcher/UpdateSchedulerView",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).UpdateSchedulerView(ctx, req.(*SchedulerViewUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dispatcher_Statis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DispatcherServer).Statis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dispatcher.dispatcher/Statis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DispatcherServer).Statis(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Dispatcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dispatcher.dispatcher",
	HandlerType: (*DispatcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Dispatch",
			Handler:    _Dispatcher_Dispatch_Handler,
		},
		{
			MethodName: "UpdateInstanceView",
			Handler:    _Dispatcher_UpdateInstanceView_Handler,
		},
		{
			MethodName: "UpdateSchedulerView",
			Handler:    _Dispatcher_UpdateSchedulerView_Handler,
		},
		{
			MethodName: "Statis",
			Handler:    _Dispatcher_Statis_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/dispatcher.proto",
}
