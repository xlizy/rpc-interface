// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: AuthMessage.proto

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

type AnalysisTokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *AnalysisTokenReq) Reset() {
	*x = AnalysisTokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisTokenReq) ProtoMessage() {}

func (x *AnalysisTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_AuthMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisTokenReq.ProtoReflect.Descriptor instead.
func (*AnalysisTokenReq) Descriptor() ([]byte, []int) {
	return file_AuthMessage_proto_rawDescGZIP(), []int{0}
}

func (x *AnalysisTokenReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *AnalysisTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type AnalysisTokenRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Roles           []string `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles,omitempty"`
	NewAccessToken  string   `protobuf:"bytes,3,opt,name=newAccessToken,proto3" json:"newAccessToken,omitempty"`
	NewRefreshToken string   `protobuf:"bytes,4,opt,name=newRefreshToken,proto3" json:"newRefreshToken,omitempty"`
}

func (x *AnalysisTokenRes) Reset() {
	*x = AnalysisTokenRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisTokenRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisTokenRes) ProtoMessage() {}

func (x *AnalysisTokenRes) ProtoReflect() protoreflect.Message {
	mi := &file_AuthMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisTokenRes.ProtoReflect.Descriptor instead.
func (*AnalysisTokenRes) Descriptor() ([]byte, []int) {
	return file_AuthMessage_proto_rawDescGZIP(), []int{1}
}

func (x *AnalysisTokenRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AnalysisTokenRes) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *AnalysisTokenRes) GetNewAccessToken() string {
	if x != nil {
		return x.NewAccessToken
	}
	return ""
}

func (x *AnalysisTokenRes) GetNewRefreshToken() string {
	if x != nil {
		return x.NewRefreshToken
	}
	return ""
}

type AnalysisTokenRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result           `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Data   *AnalysisTokenRes `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *AnalysisTokenRsp) Reset() {
	*x = AnalysisTokenRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnalysisTokenRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnalysisTokenRsp) ProtoMessage() {}

func (x *AnalysisTokenRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AuthMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnalysisTokenRsp.ProtoReflect.Descriptor instead.
func (*AnalysisTokenRsp) Descriptor() ([]byte, []int) {
	return file_AuthMessage_proto_rawDescGZIP(), []int{2}
}

func (x *AnalysisTokenRsp) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *AnalysisTokenRsp) GetData() *AnalysisTokenRes {
	if x != nil {
		return x.Data
	}
	return nil
}

type CheckPermissionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
	Url          string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	HttpMethod   string `protobuf:"bytes,4,opt,name=httpMethod,proto3" json:"httpMethod,omitempty"`
}

func (x *CheckPermissionReq) Reset() {
	*x = CheckPermissionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthMessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionReq) ProtoMessage() {}

func (x *CheckPermissionReq) ProtoReflect() protoreflect.Message {
	mi := &file_AuthMessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionReq.ProtoReflect.Descriptor instead.
func (*CheckPermissionReq) Descriptor() ([]byte, []int) {
	return file_AuthMessage_proto_rawDescGZIP(), []int{3}
}

func (x *CheckPermissionReq) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *CheckPermissionReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *CheckPermissionReq) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CheckPermissionReq) GetHttpMethod() string {
	if x != nil {
		return x.HttpMethod
	}
	return ""
}

type CheckPermissionRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	HavePermission  bool   `protobuf:"varint,2,opt,name=havePermission,proto3" json:"havePermission,omitempty"`
	NewAccessToken  string `protobuf:"bytes,3,opt,name=newAccessToken,proto3" json:"newAccessToken,omitempty"`
	NewRefreshToken string `protobuf:"bytes,4,opt,name=newRefreshToken,proto3" json:"newRefreshToken,omitempty"`
}

func (x *CheckPermissionRes) Reset() {
	*x = CheckPermissionRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthMessage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRes) ProtoMessage() {}

func (x *CheckPermissionRes) ProtoReflect() protoreflect.Message {
	mi := &file_AuthMessage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRes.ProtoReflect.Descriptor instead.
func (*CheckPermissionRes) Descriptor() ([]byte, []int) {
	return file_AuthMessage_proto_rawDescGZIP(), []int{4}
}

func (x *CheckPermissionRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckPermissionRes) GetHavePermission() bool {
	if x != nil {
		return x.HavePermission
	}
	return false
}

func (x *CheckPermissionRes) GetNewAccessToken() string {
	if x != nil {
		return x.NewAccessToken
	}
	return ""
}

func (x *CheckPermissionRes) GetNewRefreshToken() string {
	if x != nil {
		return x.NewRefreshToken
	}
	return ""
}

type CheckPermissionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *Result             `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Data   *CheckPermissionRes `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *CheckPermissionRsp) Reset() {
	*x = CheckPermissionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_AuthMessage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPermissionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPermissionRsp) ProtoMessage() {}

func (x *CheckPermissionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_AuthMessage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPermissionRsp.ProtoReflect.Descriptor instead.
func (*CheckPermissionRsp) Descriptor() ([]byte, []int) {
	return file_AuthMessage_proto_rawDescGZIP(), []int{5}
}

func (x *CheckPermissionRsp) GetResult() *Result {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CheckPermissionRsp) GetData() *CheckPermissionRes {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_AuthMessage_proto protoreflect.FileDescriptor

var file_AuthMessage_proto_rawDesc = []byte{
	0x0a, 0x11, 0x41, 0x75, 0x74, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x58, 0x0a, 0x10, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x10,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12, 0x26,
	0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x5a, 0x0a, 0x10, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x41, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x8c, 0x01, 0x0a,
	0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x68,
	0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x68, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0xa6, 0x01, 0x0a, 0x12,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x68, 0x61,
	0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0e, 0x68, 0x61, 0x76, 0x65, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x65, 0x77, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x65, 0x77, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6e, 0x65,
	0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x6e, 0x65, 0x77, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x5e, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x3b, 0x72,
	0x70, 0x63, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_AuthMessage_proto_rawDescOnce sync.Once
	file_AuthMessage_proto_rawDescData = file_AuthMessage_proto_rawDesc
)

func file_AuthMessage_proto_rawDescGZIP() []byte {
	file_AuthMessage_proto_rawDescOnce.Do(func() {
		file_AuthMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_AuthMessage_proto_rawDescData)
	})
	return file_AuthMessage_proto_rawDescData
}

var file_AuthMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_AuthMessage_proto_goTypes = []interface{}{
	(*AnalysisTokenReq)(nil),   // 0: AnalysisTokenReq
	(*AnalysisTokenRes)(nil),   // 1: AnalysisTokenRes
	(*AnalysisTokenRsp)(nil),   // 2: AnalysisTokenRsp
	(*CheckPermissionReq)(nil), // 3: CheckPermissionReq
	(*CheckPermissionRes)(nil), // 4: CheckPermissionRes
	(*CheckPermissionRsp)(nil), // 5: CheckPermissionRsp
	(*Result)(nil),             // 6: Result
}
var file_AuthMessage_proto_depIdxs = []int32{
	6, // 0: AnalysisTokenRsp.result:type_name -> Result
	1, // 1: AnalysisTokenRsp.data:type_name -> AnalysisTokenRes
	6, // 2: CheckPermissionRsp.result:type_name -> Result
	4, // 3: CheckPermissionRsp.data:type_name -> CheckPermissionRes
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_AuthMessage_proto_init() }
func file_AuthMessage_proto_init() {
	if File_AuthMessage_proto != nil {
		return
	}
	file_Result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_AuthMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisTokenReq); i {
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
		file_AuthMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisTokenRes); i {
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
		file_AuthMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnalysisTokenRsp); i {
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
		file_AuthMessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionReq); i {
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
		file_AuthMessage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRes); i {
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
		file_AuthMessage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPermissionRsp); i {
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
			RawDescriptor: file_AuthMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_AuthMessage_proto_goTypes,
		DependencyIndexes: file_AuthMessage_proto_depIdxs,
		MessageInfos:      file_AuthMessage_proto_msgTypes,
	}.Build()
	File_AuthMessage_proto = out.File
	file_AuthMessage_proto_rawDesc = nil
	file_AuthMessage_proto_goTypes = nil
	file_AuthMessage_proto_depIdxs = nil
}
