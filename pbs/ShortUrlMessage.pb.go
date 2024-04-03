// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: ShortUrlMessage.proto

package rpcApi

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

type CreateShortUrlReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       int32  `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty"`
	OriginUrl  string `protobuf:"bytes,2,opt,name=originUrl,proto3" json:"originUrl,omitempty"`
	ExpireTime string `protobuf:"bytes,3,opt,name=expireTime,proto3" json:"expireTime,omitempty"`
	Times      int32  `protobuf:"varint,4,opt,name=times,proto3" json:"times,omitempty"`
}

func (x *CreateShortUrlReq) Reset() {
	*x = CreateShortUrlReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShortUrlMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlReq) ProtoMessage() {}

func (x *CreateShortUrlReq) ProtoReflect() protoreflect.Message {
	mi := &file_ShortUrlMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlReq.ProtoReflect.Descriptor instead.
func (*CreateShortUrlReq) Descriptor() ([]byte, []int) {
	return file_ShortUrlMessage_proto_rawDescGZIP(), []int{0}
}

func (x *CreateShortUrlReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *CreateShortUrlReq) GetOriginUrl() string {
	if x != nil {
		return x.OriginUrl
	}
	return ""
}

func (x *CreateShortUrlReq) GetExpireTime() string {
	if x != nil {
		return x.ExpireTime
	}
	return ""
}

func (x *CreateShortUrlReq) GetTimes() int32 {
	if x != nil {
		return x.Times
	}
	return 0
}

type CreateShortUrlInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *CreateShortUrlInfo) Reset() {
	*x = CreateShortUrlInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShortUrlMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlInfo) ProtoMessage() {}

func (x *CreateShortUrlInfo) ProtoReflect() protoreflect.Message {
	mi := &file_ShortUrlMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlInfo.ProtoReflect.Descriptor instead.
func (*CreateShortUrlInfo) Descriptor() ([]byte, []int) {
	return file_ShortUrlMessage_proto_rawDescGZIP(), []int{1}
}

func (x *CreateShortUrlInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type CreateShortUrlRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result             `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Data   *CreateShortUrlInfo `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CreateShortUrlRsp) Reset() {
	*x = CreateShortUrlRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ShortUrlMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateShortUrlRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateShortUrlRsp) ProtoMessage() {}

func (x *CreateShortUrlRsp) ProtoReflect() protoreflect.Message {
	mi := &file_ShortUrlMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateShortUrlRsp.ProtoReflect.Descriptor instead.
func (*CreateShortUrlRsp) Descriptor() ([]byte, []int) {
	return file_ShortUrlMessage_proto_rawDescGZIP(), []int{2}
}

func (x *CreateShortUrlRsp) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CreateShortUrlRsp) GetData() *CreateShortUrlInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_ShortUrlMessage_proto protoreflect.FileDescriptor

var file_ShortUrlMessage_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x55, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x22, 0x26, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x5d, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x12,
	0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2e, 0x2f,
	0x70, 0x62, 0x73, 0x3b, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_ShortUrlMessage_proto_rawDescOnce sync.Once
	file_ShortUrlMessage_proto_rawDescData = file_ShortUrlMessage_proto_rawDesc
)

func file_ShortUrlMessage_proto_rawDescGZIP() []byte {
	file_ShortUrlMessage_proto_rawDescOnce.Do(func() {
		file_ShortUrlMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_ShortUrlMessage_proto_rawDescData)
	})
	return file_ShortUrlMessage_proto_rawDescData
}

var file_ShortUrlMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ShortUrlMessage_proto_goTypes = []interface{}{
	(*CreateShortUrlReq)(nil),  // 0: CreateShortUrlReq
	(*CreateShortUrlInfo)(nil), // 1: CreateShortUrlInfo
	(*CreateShortUrlRsp)(nil),  // 2: CreateShortUrlRsp
	(*Result)(nil),             // 3: Result
}
var file_ShortUrlMessage_proto_depIdxs = []int32{
	3, // 0: CreateShortUrlRsp.result:type_name -> Result
	1, // 1: CreateShortUrlRsp.data:type_name -> CreateShortUrlInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ShortUrlMessage_proto_init() }
func file_ShortUrlMessage_proto_init() {
	if File_ShortUrlMessage_proto != nil {
		return
	}
	file_Result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ShortUrlMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlReq); i {
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
		file_ShortUrlMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlInfo); i {
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
		file_ShortUrlMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateShortUrlRsp); i {
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
			RawDescriptor: file_ShortUrlMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ShortUrlMessage_proto_goTypes,
		DependencyIndexes: file_ShortUrlMessage_proto_depIdxs,
		MessageInfos:      file_ShortUrlMessage_proto_msgTypes,
	}.Build()
	File_ShortUrlMessage_proto = out.File
	file_ShortUrlMessage_proto_rawDesc = nil
	file_ShortUrlMessage_proto_goTypes = nil
	file_ShortUrlMessage_proto_depIdxs = nil
}
