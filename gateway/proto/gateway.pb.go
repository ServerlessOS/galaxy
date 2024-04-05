// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: gateway.proto

package gateway_rpc

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

type UpdateGatewayListReq_OperationType int32

const (
	UpdateGatewayListReq_APPEND   UpdateGatewayListReq_OperationType = 0
	UpdateGatewayListReq_REDUCE   UpdateGatewayListReq_OperationType = 1
	UpdateGatewayListReq_OVERRIDE UpdateGatewayListReq_OperationType = 2
)

// Enum value maps for UpdateGatewayListReq_OperationType.
var (
	UpdateGatewayListReq_OperationType_name = map[int32]string{
		0: "APPEND",
		1: "REDUCE",
		2: "OVERRIDE",
	}
	UpdateGatewayListReq_OperationType_value = map[string]int32{
		"APPEND":   0,
		"REDUCE":   1,
		"OVERRIDE": 2,
	}
)

func (x UpdateGatewayListReq_OperationType) Enum() *UpdateGatewayListReq_OperationType {
	p := new(UpdateGatewayListReq_OperationType)
	*p = x
	return p
}

func (x UpdateGatewayListReq_OperationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UpdateGatewayListReq_OperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_gateway_proto_enumTypes[0].Descriptor()
}

func (UpdateGatewayListReq_OperationType) Type() protoreflect.EnumType {
	return &file_gateway_proto_enumTypes[0]
}

func (x UpdateGatewayListReq_OperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UpdateGatewayListReq_OperationType.Descriptor instead.
func (UpdateGatewayListReq_OperationType) EnumDescriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0, 0}
}

type UpdateGatewayListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        UpdateGatewayListReq_OperationType `protobuf:"varint,1,opt,name=type,proto3,enum=gateway_rpc.UpdateGatewayListReq_OperationType" json:"type,omitempty"` // 操作类型 追加/减少/覆盖
	GatewayList map[string]string                  `protobuf:"bytes,2,rep,name=gatewayList,proto3" json:"gatewayList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UpdateGatewayListReq) Reset() {
	*x = UpdateGatewayListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGatewayListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGatewayListReq) ProtoMessage() {}

func (x *UpdateGatewayListReq) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGatewayListReq.ProtoReflect.Descriptor instead.
func (*UpdateGatewayListReq) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateGatewayListReq) GetType() UpdateGatewayListReq_OperationType {
	if x != nil {
		return x.Type
	}
	return UpdateGatewayListReq_APPEND
}

func (x *UpdateGatewayListReq) GetGatewayList() map[string]string {
	if x != nil {
		return x.GatewayList
	}
	return nil
}

type GeneralResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode       int64  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Description      string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	ErrorInformation string `protobuf:"bytes,3,opt,name=errorInformation,proto3" json:"errorInformation,omitempty"`
}

func (x *GeneralResp) Reset() {
	*x = GeneralResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneralResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneralResp) ProtoMessage() {}

func (x *GeneralResp) ProtoReflect() protoreflect.Message {
	mi := &file_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneralResp.ProtoReflect.Descriptor instead.
func (*GeneralResp) Descriptor() ([]byte, []int) {
	return file_gateway_proto_rawDescGZIP(), []int{1}
}

func (x *GeneralResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *GeneralResp) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GeneralResp) GetErrorInformation() string {
	if x != nil {
		return x.ErrorInformation
	}
	return ""
}

var File_gateway_proto protoreflect.FileDescriptor

var file_gateway_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x70, 0x63, 0x22, 0xa8, 0x02, 0x0a,
	0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x43, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x3e, 0x0a, 0x10, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0x35, 0x0a, 0x0d, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x50, 0x50, 0x45, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x52, 0x45, 0x44, 0x55, 0x43, 0x45, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x56, 0x45,
	0x52, 0x52, 0x49, 0x44, 0x45, 0x10, 0x02, 0x22, 0x7b, 0x0a, 0x0b, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a, 0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x32, 0x5d, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12,
	0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gateway_proto_rawDescOnce sync.Once
	file_gateway_proto_rawDescData = file_gateway_proto_rawDesc
)

func file_gateway_proto_rawDescGZIP() []byte {
	file_gateway_proto_rawDescOnce.Do(func() {
		file_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_gateway_proto_rawDescData)
	})
	return file_gateway_proto_rawDescData
}

var file_gateway_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gateway_proto_goTypes = []interface{}{
	(UpdateGatewayListReq_OperationType)(0), // 0: gateway_rpc.UpdateGatewayListReq.OperationType
	(*UpdateGatewayListReq)(nil),            // 1: gateway_rpc.UpdateGatewayListReq
	(*GeneralResp)(nil),                     // 2: gateway_rpc.GeneralResp
	nil,                                     // 3: gateway_rpc.UpdateGatewayListReq.GatewayListEntry
}
var file_gateway_proto_depIdxs = []int32{
	0, // 0: gateway_rpc.UpdateGatewayListReq.type:type_name -> gateway_rpc.UpdateGatewayListReq.OperationType
	3, // 1: gateway_rpc.UpdateGatewayListReq.gatewayList:type_name -> gateway_rpc.UpdateGatewayListReq.GatewayListEntry
	1, // 2: gateway_rpc.Gateway.UpdateGatewayList:input_type -> gateway_rpc.UpdateGatewayListReq
	2, // 3: gateway_rpc.Gateway.UpdateGatewayList:output_type -> gateway_rpc.GeneralResp
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gateway_proto_init() }
func file_gateway_proto_init() {
	if File_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGatewayListReq); i {
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
		file_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeneralResp); i {
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
			RawDescriptor: file_gateway_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gateway_proto_goTypes,
		DependencyIndexes: file_gateway_proto_depIdxs,
		EnumInfos:         file_gateway_proto_enumTypes,
		MessageInfos:      file_gateway_proto_msgTypes,
	}.Build()
	File_gateway_proto = out.File
	file_gateway_proto_rawDesc = nil
	file_gateway_proto_goTypes = nil
	file_gateway_proto_depIdxs = nil
}
