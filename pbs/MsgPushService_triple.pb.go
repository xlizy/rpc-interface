// Code generated by protoc-gen-go-triple. DO NOT EDIT.
// versions:
// - protoc-gen-go-triple v1.0.5
// - protoc             v4.25.3
// source: MsgPushService.proto

package rpcApi

import (
	context "context"
	protocol "dubbo.apache.org/dubbo-go/v3/protocol"
	dubbo3 "dubbo.apache.org/dubbo-go/v3/protocol/dubbo3"
	invocation "dubbo.apache.org/dubbo-go/v3/protocol/invocation"
	grpc_go "github.com/dubbogo/grpc-go"
	codes "github.com/dubbogo/grpc-go/codes"
	metadata "github.com/dubbogo/grpc-go/metadata"
	status "github.com/dubbogo/grpc-go/status"
	common "github.com/dubbogo/triple/pkg/common"
	constant "github.com/dubbogo/triple/pkg/common/constant"
	triple "github.com/dubbogo/triple/pkg/triple"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc_go.SupportPackageIsVersion7

// MsgPushServiceClient is the client API for MsgPushService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgPushServiceClient interface {
	SendByDingTalkSimpleNotice(ctx context.Context, in *SendByDingTalkSimpleNoticeReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
	SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
	SendSmsBatch(ctx context.Context, in *SendSmsBatchReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
	SendMail(ctx context.Context, in *SendMailReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
	SendSmsVerCode(ctx context.Context, in *SendSmsVerCodeReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
	CheckMobileVerCode(ctx context.Context, in *CheckMobileVerCodeReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
	SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment)
}

type msgPushServiceClient struct {
	cc *triple.TripleConn
}

type MsgPushServiceClientImpl struct {
	SendByDingTalkSimpleNotice func(ctx context.Context, in *SendByDingTalkSimpleNoticeReq) (*Result, error)
	SendSms                    func(ctx context.Context, in *SendSmsReq) (*Result, error)
	SendSmsBatch               func(ctx context.Context, in *SendSmsBatchReq) (*Result, error)
	SendMail                   func(ctx context.Context, in *SendMailReq) (*Result, error)
	SendSmsVerCode             func(ctx context.Context, in *SendSmsVerCodeReq) (*Result, error)
	CheckMobileVerCode         func(ctx context.Context, in *CheckMobileVerCodeReq) (*Result, error)
	SendMessage                func(ctx context.Context, in *SendMessageReq) (*Result, error)
}

func (c *MsgPushServiceClientImpl) GetDubboStub(cc *triple.TripleConn) MsgPushServiceClient {
	return NewMsgPushServiceClient(cc)
}

func (c *MsgPushServiceClientImpl) XXX_InterfaceName() string {
	return "MsgPushService"
}

func NewMsgPushServiceClient(cc *triple.TripleConn) MsgPushServiceClient {
	return &msgPushServiceClient{cc}
}

func (c *msgPushServiceClient) SendByDingTalkSimpleNotice(ctx context.Context, in *SendByDingTalkSimpleNoticeReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SendByDingTalkSimpleNotice", in, out)
}

func (c *msgPushServiceClient) SendSms(ctx context.Context, in *SendSmsReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SendSms", in, out)
}

func (c *msgPushServiceClient) SendSmsBatch(ctx context.Context, in *SendSmsBatchReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SendSmsBatch", in, out)
}

func (c *msgPushServiceClient) SendMail(ctx context.Context, in *SendMailReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SendMail", in, out)
}

func (c *msgPushServiceClient) SendSmsVerCode(ctx context.Context, in *SendSmsVerCodeReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SendSmsVerCode", in, out)
}

func (c *msgPushServiceClient) CheckMobileVerCode(ctx context.Context, in *CheckMobileVerCodeReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/CheckMobileVerCode", in, out)
}

func (c *msgPushServiceClient) SendMessage(ctx context.Context, in *SendMessageReq, opts ...grpc_go.CallOption) (*Result, common.ErrorWithAttachment) {
	out := new(Result)
	interfaceKey := ctx.Value(constant.InterfaceKey).(string)
	return out, c.cc.Invoke(ctx, "/"+interfaceKey+"/SendMessage", in, out)
}

// MsgPushServiceServer is the server API for MsgPushService service.
// All implementations must embed UnimplementedMsgPushServiceServer
// for forward compatibility
type MsgPushServiceServer interface {
	SendByDingTalkSimpleNotice(context.Context, *SendByDingTalkSimpleNoticeReq) (*Result, error)
	SendSms(context.Context, *SendSmsReq) (*Result, error)
	SendSmsBatch(context.Context, *SendSmsBatchReq) (*Result, error)
	SendMail(context.Context, *SendMailReq) (*Result, error)
	SendSmsVerCode(context.Context, *SendSmsVerCodeReq) (*Result, error)
	CheckMobileVerCode(context.Context, *CheckMobileVerCodeReq) (*Result, error)
	SendMessage(context.Context, *SendMessageReq) (*Result, error)
	mustEmbedUnimplementedMsgPushServiceServer()
}

// UnimplementedMsgPushServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMsgPushServiceServer struct {
	proxyImpl protocol.Invoker
}

func (UnimplementedMsgPushServiceServer) SendByDingTalkSimpleNotice(context.Context, *SendByDingTalkSimpleNoticeReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendByDingTalkSimpleNotice not implemented")
}
func (UnimplementedMsgPushServiceServer) SendSms(context.Context, *SendSmsReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedMsgPushServiceServer) SendSmsBatch(context.Context, *SendSmsBatchReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsBatch not implemented")
}
func (UnimplementedMsgPushServiceServer) SendMail(context.Context, *SendMailReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMail not implemented")
}
func (UnimplementedMsgPushServiceServer) SendSmsVerCode(context.Context, *SendSmsVerCodeReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSmsVerCode not implemented")
}
func (UnimplementedMsgPushServiceServer) CheckMobileVerCode(context.Context, *CheckMobileVerCodeReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckMobileVerCode not implemented")
}
func (UnimplementedMsgPushServiceServer) SendMessage(context.Context, *SendMessageReq) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (s *UnimplementedMsgPushServiceServer) XXX_SetProxyImpl(impl protocol.Invoker) {
	s.proxyImpl = impl
}

func (s *UnimplementedMsgPushServiceServer) XXX_GetProxyImpl() protocol.Invoker {
	return s.proxyImpl
}

func (s *UnimplementedMsgPushServiceServer) XXX_ServiceDesc() *grpc_go.ServiceDesc {
	return &MsgPushService_ServiceDesc
}
func (s *UnimplementedMsgPushServiceServer) XXX_InterfaceName() string {
	return "MsgPushService"
}

func (UnimplementedMsgPushServiceServer) mustEmbedUnimplementedMsgPushServiceServer() {}

// UnsafeMsgPushServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgPushServiceServer will
// result in compilation errors.
type UnsafeMsgPushServiceServer interface {
	mustEmbedUnimplementedMsgPushServiceServer()
}

func RegisterMsgPushServiceServer(s grpc_go.ServiceRegistrar, srv MsgPushServiceServer) {
	s.RegisterService(&MsgPushService_ServiceDesc, srv)
}

func _MsgPushService_SendByDingTalkSimpleNotice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendByDingTalkSimpleNoticeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SendByDingTalkSimpleNotice", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgPushService_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SendSms", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgPushService_SendSmsBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsBatchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SendSmsBatch", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgPushService_SendMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SendMail", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgPushService_SendSmsVerCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsVerCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SendSmsVerCode", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgPushService_CheckMobileVerCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckMobileVerCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("CheckMobileVerCode", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgPushService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc_go.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	base := srv.(dubbo3.Dubbo3GrpcService)
	args := []interface{}{}
	args = append(args, in)
	md, _ := metadata.FromIncomingContext(ctx)
	invAttachment := make(map[string]interface{}, len(md))
	for k, v := range md {
		invAttachment[k] = v
	}
	invo := invocation.NewRPCInvocation("SendMessage", args, invAttachment)
	if interceptor == nil {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	info := &grpc_go.UnaryServerInfo{
		Server:     srv,
		FullMethod: ctx.Value("XXX_TRIPLE_GO_INTERFACE_NAME").(string),
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		result := base.XXX_GetProxyImpl().Invoke(ctx, invo)
		return result, result.Error()
	}
	return interceptor(ctx, in, info, handler)
}

// MsgPushService_ServiceDesc is the grpc_go.ServiceDesc for MsgPushService service.
// It's only intended for direct use with grpc_go.RegisterService,
// and not to be introspected or modified (even as a copy)
var MsgPushService_ServiceDesc = grpc_go.ServiceDesc{
	ServiceName: "MsgPushService",
	HandlerType: (*MsgPushServiceServer)(nil),
	Methods: []grpc_go.MethodDesc{
		{
			MethodName: "SendByDingTalkSimpleNotice",
			Handler:    _MsgPushService_SendByDingTalkSimpleNotice_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _MsgPushService_SendSms_Handler,
		},
		{
			MethodName: "SendSmsBatch",
			Handler:    _MsgPushService_SendSmsBatch_Handler,
		},
		{
			MethodName: "SendMail",
			Handler:    _MsgPushService_SendMail_Handler,
		},
		{
			MethodName: "SendSmsVerCode",
			Handler:    _MsgPushService_SendSmsVerCode_Handler,
		},
		{
			MethodName: "CheckMobileVerCode",
			Handler:    _MsgPushService_CheckMobileVerCode_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _MsgPushService_SendMessage_Handler,
		},
	},
	Streams:  []grpc_go.StreamDesc{},
	Metadata: "MsgPushService.proto",
}
