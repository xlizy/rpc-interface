// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: ShortUrlService.proto

package rpcApi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_ShortUrlService_proto protoreflect.FileDescriptor

var file_ShortUrlService_proto_rawDesc = []byte{
	0x0a, 0x15, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x4d,
	0x0a, 0x0f, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x55, 0x72, 0x6c, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x3b, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_ShortUrlService_proto_goTypes = []interface{}{
	(*CreateShortUrlReq)(nil), // 0: CreateShortUrlReq
	(*CreateShortUrlRsp)(nil), // 1: CreateShortUrlRsp
}
var file_ShortUrlService_proto_depIdxs = []int32{
	0, // 0: ShortUrlService.CreateShortUrl:input_type -> CreateShortUrlReq
	1, // 1: ShortUrlService.CreateShortUrl:output_type -> CreateShortUrlRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_ShortUrlService_proto_init() }
func file_ShortUrlService_proto_init() {
	if File_ShortUrlService_proto != nil {
		return
	}
	file_ShortUrlMessage_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ShortUrlService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ShortUrlService_proto_goTypes,
		DependencyIndexes: file_ShortUrlService_proto_depIdxs,
	}.Build()
	File_ShortUrlService_proto = out.File
	file_ShortUrlService_proto_rawDesc = nil
	file_ShortUrlService_proto_goTypes = nil
	file_ShortUrlService_proto_depIdxs = nil
}
