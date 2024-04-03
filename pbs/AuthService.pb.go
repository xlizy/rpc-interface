// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: AuthService.proto

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

var File_AuthService_proto protoreflect.FileDescriptor

var file_AuthService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x85, 0x01, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0d, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x11, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73,
	0x69, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x41, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12,
	0x3d, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0f,
	0x5a, 0x0d, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x3b, 0x72, 0x70, 0x63, 0x41, 0x70, 0x69, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_AuthService_proto_goTypes = []interface{}{
	(*AnalysisTokenReq)(nil),   // 0: AnalysisTokenReq
	(*CheckPermissionReq)(nil), // 1: CheckPermissionReq
	(*AnalysisTokenRsp)(nil),   // 2: AnalysisTokenRsp
	(*CheckPermissionRsp)(nil), // 3: CheckPermissionRsp
}
var file_AuthService_proto_depIdxs = []int32{
	0, // 0: AuthService.AnalysisToken:input_type -> AnalysisTokenReq
	1, // 1: AuthService.CheckPermission:input_type -> CheckPermissionReq
	2, // 2: AuthService.AnalysisToken:output_type -> AnalysisTokenRsp
	3, // 3: AuthService.CheckPermission:output_type -> CheckPermissionRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_AuthService_proto_init() }
func file_AuthService_proto_init() {
	if File_AuthService_proto != nil {
		return
	}
	file_AuthMessage_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_AuthService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_AuthService_proto_goTypes,
		DependencyIndexes: file_AuthService_proto_depIdxs,
	}.Build()
	File_AuthService_proto = out.File
	file_AuthService_proto_rawDesc = nil
	file_AuthService_proto_goTypes = nil
	file_AuthService_proto_depIdxs = nil
}
