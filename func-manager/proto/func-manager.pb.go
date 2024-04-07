// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: func-manager.proto

package proto

import (
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

type GeneralRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Labels    string `protobuf:"bytes,3,opt,name=labels,proto3" json:"labels,omitempty"`
}

func (x *GeneralRequest) Reset() {
	*x = GeneralRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralRequest) ProtoMessage() {}

func (x *GeneralRequest) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralRequest.ProtoReflect.Descriptor instead.
func (*GeneralRequest) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{0}
}

func (x *GeneralRequest) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *GeneralRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GeneralRequest) GetLabels() string {
	if x != nil {
		return x.Labels
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address   string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Port      string `protobuf:"bytes,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterReq) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *RegisterReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterReq) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

type CreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request     *GeneralRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
	Annotations string          `protobuf:"bytes,2,opt,name=Annotations,proto3" json:"Annotations,omitempty"`
	Document    string          `protobuf:"bytes,3,opt,name=Document,proto3" json:"Document,omitempty"`
}

func (x *CreateReq) Reset() {
	*x = CreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateReq) ProtoMessage() {}

func (x *CreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateReq.ProtoReflect.Descriptor instead.
func (*CreateReq) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{2}
}

func (x *CreateReq) GetRequest() *GeneralRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *CreateReq) GetAnnotations() string {
	if x != nil {
		return x.Annotations
	}
	return ""
}

func (x *CreateReq) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GeneralRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteReq) GetRequest() *GeneralRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request *GeneralRequest `protobuf:"bytes,1,opt,name=request,proto3" json:"request,omitempty"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{4}
}

func (x *GetReq) GetRequest() *GeneralRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

type ListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId int64 `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
}

func (x *ListReq) Reset() {
	*x = ListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListReq) ProtoMessage() {}

func (x *ListReq) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListReq.ProtoReflect.Descriptor instead.
func (*ListReq) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{5}
}

func (x *ListReq) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

type GetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	StatusCode int64  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Document   string `protobuf:"bytes,3,opt,name=Document,proto3" json:"Document,omitempty"`
}

func (x *GetResp) Reset() {
	*x = GetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResp) ProtoMessage() {}

func (x *GetResp) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResp.ProtoReflect.Descriptor instead.
func (*GetResp) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{6}
}

func (x *GetResp) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *GetResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GetResp) GetDocument() string {
	if x != nil {
		return x.Document
	}
	return ""
}

type CreateResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	StatusCode       int64  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Description      string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ErrorInformation string `protobuf:"bytes,4,opt,name=errorInformation,proto3" json:"errorInformation,omitempty"`
}

func (x *CreateResp) Reset() {
	*x = CreateResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResp) ProtoMessage() {}

func (x *CreateResp) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResp.ProtoReflect.Descriptor instead.
func (*CreateResp) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{7}
}

func (x *CreateResp) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *CreateResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CreateResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateResp) GetErrorInformation() string {
	if x != nil {
		return x.ErrorInformation
	}
	return ""
}

type DeleteResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	StatusCode       int64  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Description      string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ErrorInformation string `protobuf:"bytes,4,opt,name=errorInformation,proto3" json:"errorInformation,omitempty"`
}

func (x *DeleteResp) Reset() {
	*x = DeleteResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResp) ProtoMessage() {}

func (x *DeleteResp) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResp.ProtoReflect.Descriptor instead.
func (*DeleteResp) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteResp) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *DeleteResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DeleteResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DeleteResp) GetErrorInformation() string {
	if x != nil {
		return x.ErrorInformation
	}
	return ""
}

type RegisterResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId        int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	StatusCode       int64  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Description      string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	ErrorInformation string `protobuf:"bytes,4,opt,name=errorInformation,proto3" json:"errorInformation,omitempty"`
}

func (x *RegisterResp) Reset() {
	*x = RegisterResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResp) ProtoMessage() {}

func (x *RegisterResp) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResp.ProtoReflect.Descriptor instead.
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{9}
}

func (x *RegisterResp) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *RegisterResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *RegisterResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RegisterResp) GetErrorInformation() string {
	if x != nil {
		return x.ErrorInformation
	}
	return ""
}

type ListResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId  int64  `protobuf:"varint,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	StatusCode int64  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	List       string `protobuf:"bytes,3,opt,name=list,proto3" json:"list,omitempty"` //yaml转字符串传输
}

func (x *ListResp) Reset() {
	*x = ListResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_func_manager_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResp) ProtoMessage() {}

func (x *ListResp) ProtoReflect() protoreflect.Message {
	mi := &file_func_manager_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResp.ProtoReflect.Descriptor instead.
func (*ListResp) Descriptor() ([]byte, []int) {
	return file_func_manager_proto_rawDescGZIP(), []int{10}
}

func (x *ListResp) GetRequestId() int64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *ListResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ListResp) GetList() string {
	if x != nil {
		return x.List
	}
	return ""
}

var File_func_manager_proto protoreflect.FileDescriptor

var file_func_manager_proto_rawDesc = []byte{
	0x0a, 0x12, 0x66, 0x75, 0x6e, 0x63, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x22, 0x5a, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x22, 0x6d, 0x0a,
	0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x80, 0x01, 0x0a,
	0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x66, 0x75,
	0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x42, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a, 0x07,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x35, 0x0a,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x27, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x63, 0x0a,
	0x07, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0x98, 0x01, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x01,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x9a, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5c, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x32, 0xbc, 0x02, 0x0a, 0x0b, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x16, 0x2e,
	0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x66, 0x75, 0x6e,
	0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x32, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x66, 0x75, 0x6e, 0x63,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x12, 0x35, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x66, 0x75, 0x6e, 0x63,
	0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0f, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x18, 0x2e, 0x66, 0x75,
	0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x66, 0x75, 0x6e, 0x63, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_func_manager_proto_rawDescOnce sync.Once
	file_func_manager_proto_rawDescData = file_func_manager_proto_rawDesc
)

func file_func_manager_proto_rawDescGZIP() []byte {
	file_func_manager_proto_rawDescOnce.Do(func() {
		file_func_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_func_manager_proto_rawDescData)
	})
	return file_func_manager_proto_rawDescData
}

var file_func_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_func_manager_proto_goTypes = []interface{}{
	(*GeneralRequest)(nil), // 0: funcManager.GeneralRequest
	(*RegisterReq)(nil),    // 1: funcManager.RegisterReq
	(*CreateReq)(nil),      // 2: funcManager.CreateReq
	(*DeleteReq)(nil),      // 3: funcManager.DeleteReq
	(*GetReq)(nil),         // 4: funcManager.GetReq
	(*ListReq)(nil),        // 5: funcManager.ListReq
	(*GetResp)(nil),        // 6: funcManager.GetResp
	(*CreateResp)(nil),     // 7: funcManager.CreateResp
	(*DeleteResp)(nil),     // 8: funcManager.DeleteResp
	(*RegisterResp)(nil),   // 9: funcManager.RegisterResp
	(*ListResp)(nil),       // 10: funcManager.ListResp
}
var file_func_manager_proto_depIdxs = []int32{
	0,  // 0: funcManager.CreateReq.request:type_name -> funcManager.GeneralRequest
	0,  // 1: funcManager.DeleteReq.request:type_name -> funcManager.GeneralRequest
	0,  // 2: funcManager.GetReq.request:type_name -> funcManager.GeneralRequest
	2,  // 3: funcManager.funcManager.Create:input_type -> funcManager.CreateReq
	3,  // 4: funcManager.funcManager.Delete:input_type -> funcManager.DeleteReq
	4,  // 5: funcManager.funcManager.Get:input_type -> funcManager.GetReq
	5,  // 6: funcManager.funcManager.List:input_type -> funcManager.ListReq
	1,  // 7: funcManager.funcManager.registerGateway:input_type -> funcManager.RegisterReq
	7,  // 8: funcManager.funcManager.Create:output_type -> funcManager.CreateResp
	8,  // 9: funcManager.funcManager.Delete:output_type -> funcManager.DeleteResp
	6,  // 10: funcManager.funcManager.Get:output_type -> funcManager.GetResp
	10, // 11: funcManager.funcManager.List:output_type -> funcManager.ListResp
	9,  // 12: funcManager.funcManager.registerGateway:output_type -> funcManager.RegisterResp
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_func_manager_proto_init() }
func file_func_manager_proto_init() {
	if File_func_manager_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_func_manager_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralRequest); i {
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
		file_func_manager_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterReq); i {
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
		file_func_manager_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateReq); i {
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
		file_func_manager_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_func_manager_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_func_manager_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListReq); i {
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
		file_func_manager_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResp); i {
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
		file_func_manager_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResp); i {
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
		file_func_manager_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResp); i {
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
		file_func_manager_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResp); i {
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
		file_func_manager_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResp); i {
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
			RawDescriptor: file_func_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_func_manager_proto_goTypes,
		DependencyIndexes: file_func_manager_proto_depIdxs,
		MessageInfos:      file_func_manager_proto_msgTypes,
	}.Build()
	File_func_manager_proto = out.File
	file_func_manager_proto_rawDesc = nil
	file_func_manager_proto_goTypes = nil
	file_func_manager_proto_depIdxs = nil
}
