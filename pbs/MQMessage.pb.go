// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: MQMessage.proto

package rpc_api

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

type SendMQReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topic    string `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	ClientIp string `protobuf:"bytes,3,opt,name=clientIp,proto3" json:"clientIp,omitempty"`
	AppName  string `protobuf:"bytes,4,opt,name=appName,proto3" json:"appName,omitempty"`
}

func (x *SendMQReq) Reset() {
	*x = SendMQReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MQMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMQReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMQReq) ProtoMessage() {}

func (x *SendMQReq) ProtoReflect() protoreflect.Message {
	mi := &file_MQMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMQReq.ProtoReflect.Descriptor instead.
func (*SendMQReq) Descriptor() ([]byte, []int) {
	return file_MQMessage_proto_rawDescGZIP(), []int{0}
}

func (x *SendMQReq) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

func (x *SendMQReq) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SendMQReq) GetClientIp() string {
	if x != nil {
		return x.ClientIp
	}
	return ""
}

func (x *SendMQReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

var File_MQMessage_proto protoreflect.FileDescriptor

var file_MQMessage_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4d, 0x51, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x69, 0x0a, 0x09, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x51, 0x52, 0x65, 0x71, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x10, 0x5a, 0x0e,
	0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x3b, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MQMessage_proto_rawDescOnce sync.Once
	file_MQMessage_proto_rawDescData = file_MQMessage_proto_rawDesc
)

func file_MQMessage_proto_rawDescGZIP() []byte {
	file_MQMessage_proto_rawDescOnce.Do(func() {
		file_MQMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_MQMessage_proto_rawDescData)
	})
	return file_MQMessage_proto_rawDescData
}

var file_MQMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_MQMessage_proto_goTypes = []interface{}{
	(*SendMQReq)(nil), // 0: SendMQReq
}
var file_MQMessage_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MQMessage_proto_init() }
func file_MQMessage_proto_init() {
	if File_MQMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MQMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMQReq); i {
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
			RawDescriptor: file_MQMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MQMessage_proto_goTypes,
		DependencyIndexes: file_MQMessage_proto_depIdxs,
		MessageInfos:      file_MQMessage_proto_msgTypes,
	}.Build()
	File_MQMessage_proto = out.File
	file_MQMessage_proto_rawDesc = nil
	file_MQMessage_proto_goTypes = nil
	file_MQMessage_proto_depIdxs = nil
}
