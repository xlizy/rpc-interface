// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: MsgPushService.proto

package rpc_api

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

var File_MsgPushService_proto protoreflect.FileDescriptor

var file_MsgPushService_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4d, 0x73, 0x67, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x4d, 0x73, 0x67, 0x50, 0x75, 0x73, 0x68, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe3, 0x02, 0x0a, 0x0e, 0x4d,
	0x73, 0x67, 0x50, 0x75, 0x73, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a,
	0x1a, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x79, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x53,
	0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12, 0x1e, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x79, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x53, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x21, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d,
	0x73, 0x12, 0x0b, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x07,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x0c, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x10, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x53, 0x6d, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x23, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61,
	0x69, 0x6c, 0x12, 0x0c, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x12, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x12,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x16, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65,
	0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x29, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x0f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00,
	0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x3b, 0x72, 0x70, 0x63, 0x5f, 0x61,
	0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_MsgPushService_proto_goTypes = []interface{}{
	(*SendByDingTalkSimpleNoticeReq)(nil), // 0: SendByDingTalkSimpleNoticeReq
	(*SendSmsReq)(nil),                    // 1: SendSmsReq
	(*SendSmsBatchReq)(nil),               // 2: SendSmsBatchReq
	(*SendMailReq)(nil),                   // 3: SendMailReq
	(*SendSmsVerCodeReq)(nil),             // 4: SendSmsVerCodeReq
	(*CheckMobileVerCodeReq)(nil),         // 5: CheckMobileVerCodeReq
	(*SendMessageReq)(nil),                // 6: SendMessageReq
	(*Result)(nil),                        // 7: Result
}
var file_MsgPushService_proto_depIdxs = []int32{
	0, // 0: MsgPushService.SendByDingTalkSimpleNotice:input_type -> SendByDingTalkSimpleNoticeReq
	1, // 1: MsgPushService.SendSms:input_type -> SendSmsReq
	2, // 2: MsgPushService.SendSmsBatch:input_type -> SendSmsBatchReq
	3, // 3: MsgPushService.SendMail:input_type -> SendMailReq
	4, // 4: MsgPushService.SendSmsVerCode:input_type -> SendSmsVerCodeReq
	5, // 5: MsgPushService.CheckMobileVerCode:input_type -> CheckMobileVerCodeReq
	6, // 6: MsgPushService.SendMessage:input_type -> SendMessageReq
	7, // 7: MsgPushService.SendByDingTalkSimpleNotice:output_type -> Result
	7, // 8: MsgPushService.SendSms:output_type -> Result
	7, // 9: MsgPushService.SendSmsBatch:output_type -> Result
	7, // 10: MsgPushService.SendMail:output_type -> Result
	7, // 11: MsgPushService.SendSmsVerCode:output_type -> Result
	7, // 12: MsgPushService.CheckMobileVerCode:output_type -> Result
	7, // 13: MsgPushService.SendMessage:output_type -> Result
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_MsgPushService_proto_init() }
func file_MsgPushService_proto_init() {
	if File_MsgPushService_proto != nil {
		return
	}
	file_MsgPushMessage_proto_init()
	file_Result_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_MsgPushService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_MsgPushService_proto_goTypes,
		DependencyIndexes: file_MsgPushService_proto_depIdxs,
	}.Build()
	File_MsgPushService_proto = out.File
	file_MsgPushService_proto_rawDesc = nil
	file_MsgPushService_proto_goTypes = nil
	file_MsgPushService_proto_depIdxs = nil
}
