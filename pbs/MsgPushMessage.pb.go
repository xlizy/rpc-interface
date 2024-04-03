// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.3
// source: MsgPushMessage.proto

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

type SendByDingTalkSimpleNoticeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async         bool              `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`
	SourceSystem  string            `protobuf:"bytes,2,opt,name=sourceSystem,proto3" json:"sourceSystem,omitempty"`
	Target        string            `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	InnerCode     string            `protobuf:"bytes,4,opt,name=innerCode,proto3" json:"innerCode,omitempty"`
	TemplateParam map[string]string `protobuf:"bytes,5,rep,name=templateParam,proto3" json:"templateParam,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendByDingTalkSimpleNoticeReq) Reset() {
	*x = SendByDingTalkSimpleNoticeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendByDingTalkSimpleNoticeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendByDingTalkSimpleNoticeReq) ProtoMessage() {}

func (x *SendByDingTalkSimpleNoticeReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendByDingTalkSimpleNoticeReq.ProtoReflect.Descriptor instead.
func (*SendByDingTalkSimpleNoticeReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{0}
}

func (x *SendByDingTalkSimpleNoticeReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendByDingTalkSimpleNoticeReq) GetSourceSystem() string {
	if x != nil {
		return x.SourceSystem
	}
	return ""
}

func (x *SendByDingTalkSimpleNoticeReq) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *SendByDingTalkSimpleNoticeReq) GetInnerCode() string {
	if x != nil {
		return x.InnerCode
	}
	return ""
}

func (x *SendByDingTalkSimpleNoticeReq) GetTemplateParam() map[string]string {
	if x != nil {
		return x.TemplateParam
	}
	return nil
}

type SendSmsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async         bool              `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`
	SourceSystem  string            `protobuf:"bytes,2,opt,name=sourceSystem,proto3" json:"sourceSystem,omitempty"`
	InnerCode     string            `protobuf:"bytes,3,opt,name=innerCode,proto3" json:"innerCode,omitempty"`
	Mobiles       []string          `protobuf:"bytes,4,rep,name=mobiles,proto3" json:"mobiles,omitempty"`
	TemplateParam map[string]string `protobuf:"bytes,5,rep,name=templateParam,proto3" json:"templateParam,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendSmsReq) Reset() {
	*x = SendSmsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsReq) ProtoMessage() {}

func (x *SendSmsReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsReq.ProtoReflect.Descriptor instead.
func (*SendSmsReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{1}
}

func (x *SendSmsReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendSmsReq) GetSourceSystem() string {
	if x != nil {
		return x.SourceSystem
	}
	return ""
}

func (x *SendSmsReq) GetInnerCode() string {
	if x != nil {
		return x.InnerCode
	}
	return ""
}

func (x *SendSmsReq) GetMobiles() []string {
	if x != nil {
		return x.Mobiles
	}
	return nil
}

func (x *SendSmsReq) GetTemplateParam() map[string]string {
	if x != nil {
		return x.TemplateParam
	}
	return nil
}

type SMSBatchSend struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile        string            `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	TemplateParam map[string]string `protobuf:"bytes,2,rep,name=templateParam,proto3" json:"templateParam,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SMSBatchSend) Reset() {
	*x = SMSBatchSend{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SMSBatchSend) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SMSBatchSend) ProtoMessage() {}

func (x *SMSBatchSend) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SMSBatchSend.ProtoReflect.Descriptor instead.
func (*SMSBatchSend) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{2}
}

func (x *SMSBatchSend) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *SMSBatchSend) GetTemplateParam() map[string]string {
	if x != nil {
		return x.TemplateParam
	}
	return nil
}

type SendSmsBatchReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async        bool            `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`
	SourceSystem string          `protobuf:"bytes,2,opt,name=sourceSystem,proto3" json:"sourceSystem,omitempty"`
	InnerCode    string          `protobuf:"bytes,3,opt,name=innerCode,proto3" json:"innerCode,omitempty"`
	Data         []*SMSBatchSend `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *SendSmsBatchReq) Reset() {
	*x = SendSmsBatchReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsBatchReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsBatchReq) ProtoMessage() {}

func (x *SendSmsBatchReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsBatchReq.ProtoReflect.Descriptor instead.
func (*SendSmsBatchReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{3}
}

func (x *SendSmsBatchReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendSmsBatchReq) GetSourceSystem() string {
	if x != nil {
		return x.SourceSystem
	}
	return ""
}

func (x *SendSmsBatchReq) GetInnerCode() string {
	if x != nil {
		return x.InnerCode
	}
	return ""
}

func (x *SendSmsBatchReq) GetData() []*SMSBatchSend {
	if x != nil {
		return x.Data
	}
	return nil
}

type MailFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url         string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *MailFile) Reset() {
	*x = MailFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MailFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MailFile) ProtoMessage() {}

func (x *MailFile) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MailFile.ProtoReflect.Descriptor instead.
func (*MailFile) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{4}
}

func (x *MailFile) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *MailFile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MailFile) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type SendMailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async         bool              `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`
	SourceSystem  string            `protobuf:"bytes,2,opt,name=sourceSystem,proto3" json:"sourceSystem,omitempty"`
	InnerCode     string            `protobuf:"bytes,3,opt,name=innerCode,proto3" json:"innerCode,omitempty"`
	Subject       string            `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	To            []string          `protobuf:"bytes,5,rep,name=to,proto3" json:"to,omitempty"`
	Cc            []string          `protobuf:"bytes,6,rep,name=cc,proto3" json:"cc,omitempty"`
	Bcc           []string          `protobuf:"bytes,7,rep,name=bcc,proto3" json:"bcc,omitempty"`
	TemplateParam map[string]string `protobuf:"bytes,8,rep,name=templateParam,proto3" json:"templateParam,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SubjectParam  map[string]string `protobuf:"bytes,9,rep,name=subjectParam,proto3" json:"subjectParam,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Files         []*MailFile       `protobuf:"bytes,10,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *SendMailReq) Reset() {
	*x = SendMailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMailReq) ProtoMessage() {}

func (x *SendMailReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMailReq.ProtoReflect.Descriptor instead.
func (*SendMailReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{5}
}

func (x *SendMailReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendMailReq) GetSourceSystem() string {
	if x != nil {
		return x.SourceSystem
	}
	return ""
}

func (x *SendMailReq) GetInnerCode() string {
	if x != nil {
		return x.InnerCode
	}
	return ""
}

func (x *SendMailReq) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *SendMailReq) GetTo() []string {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *SendMailReq) GetCc() []string {
	if x != nil {
		return x.Cc
	}
	return nil
}

func (x *SendMailReq) GetBcc() []string {
	if x != nil {
		return x.Bcc
	}
	return nil
}

func (x *SendMailReq) GetTemplateParam() map[string]string {
	if x != nil {
		return x.TemplateParam
	}
	return nil
}

func (x *SendMailReq) GetSubjectParam() map[string]string {
	if x != nil {
		return x.SubjectParam
	}
	return nil
}

func (x *SendMailReq) GetFiles() []*MailFile {
	if x != nil {
		return x.Files
	}
	return nil
}

type SendSmsVerCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async          bool   `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`
	SourceSystem   string `protobuf:"bytes,2,opt,name=sourceSystem,proto3" json:"sourceSystem,omitempty"`
	InnerCode      string `protobuf:"bytes,3,opt,name=innerCode,proto3" json:"innerCode,omitempty"`
	Mobile         int64  `protobuf:"varint,4,opt,name=mobile,proto3" json:"mobile,omitempty"`
	MobileAreaCode int32  `protobuf:"varint,5,opt,name=mobileAreaCode,proto3" json:"mobileAreaCode,omitempty"`
}

func (x *SendSmsVerCodeReq) Reset() {
	*x = SendSmsVerCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSmsVerCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSmsVerCodeReq) ProtoMessage() {}

func (x *SendSmsVerCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSmsVerCodeReq.ProtoReflect.Descriptor instead.
func (*SendSmsVerCodeReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{6}
}

func (x *SendSmsVerCodeReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendSmsVerCodeReq) GetSourceSystem() string {
	if x != nil {
		return x.SourceSystem
	}
	return ""
}

func (x *SendSmsVerCodeReq) GetInnerCode() string {
	if x != nil {
		return x.InnerCode
	}
	return ""
}

func (x *SendSmsVerCodeReq) GetMobile() int64 {
	if x != nil {
		return x.Mobile
	}
	return 0
}

func (x *SendSmsVerCodeReq) GetMobileAreaCode() int32 {
	if x != nil {
		return x.MobileAreaCode
	}
	return 0
}

type CheckMobileVerCodeReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InnerCode      string `protobuf:"bytes,1,opt,name=innerCode,proto3" json:"innerCode,omitempty"`
	Mobile         int64  `protobuf:"varint,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	MobileAreaCode int32  `protobuf:"varint,3,opt,name=mobileAreaCode,proto3" json:"mobileAreaCode,omitempty"`
	VerCodeValue   string `protobuf:"bytes,4,opt,name=verCodeValue,proto3" json:"verCodeValue,omitempty"`
}

func (x *CheckMobileVerCodeReq) Reset() {
	*x = CheckMobileVerCodeReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckMobileVerCodeReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckMobileVerCodeReq) ProtoMessage() {}

func (x *CheckMobileVerCodeReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckMobileVerCodeReq.ProtoReflect.Descriptor instead.
func (*CheckMobileVerCodeReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{7}
}

func (x *CheckMobileVerCodeReq) GetInnerCode() string {
	if x != nil {
		return x.InnerCode
	}
	return ""
}

func (x *CheckMobileVerCodeReq) GetMobile() int64 {
	if x != nil {
		return x.Mobile
	}
	return 0
}

func (x *CheckMobileVerCodeReq) GetMobileAreaCode() int32 {
	if x != nil {
		return x.MobileAreaCode
	}
	return 0
}

func (x *CheckMobileVerCodeReq) GetVerCodeValue() string {
	if x != nil {
		return x.VerCodeValue
	}
	return ""
}

type SendMessageReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Async        bool              `protobuf:"varint,1,opt,name=async,proto3" json:"async,omitempty"`
	SourceSystem string            `protobuf:"bytes,2,opt,name=sourceSystem,proto3" json:"sourceSystem,omitempty"`
	UserId       string            `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
	MsgType      int32             `protobuf:"varint,4,opt,name=msgType,proto3" json:"msgType,omitempty"`
	Manual       int32             `protobuf:"varint,5,opt,name=manual,proto3" json:"manual,omitempty"`
	Title        string            `protobuf:"bytes,6,opt,name=title,proto3" json:"title,omitempty"`
	Content      map[string]string `protobuf:"bytes,7,rep,name=content,proto3" json:"content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SendMessageReq) Reset() {
	*x = SendMessageReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MsgPushMessage_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageReq) ProtoMessage() {}

func (x *SendMessageReq) ProtoReflect() protoreflect.Message {
	mi := &file_MsgPushMessage_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageReq.ProtoReflect.Descriptor instead.
func (*SendMessageReq) Descriptor() ([]byte, []int) {
	return file_MsgPushMessage_proto_rawDescGZIP(), []int{8}
}

func (x *SendMessageReq) GetAsync() bool {
	if x != nil {
		return x.Async
	}
	return false
}

func (x *SendMessageReq) GetSourceSystem() string {
	if x != nil {
		return x.SourceSystem
	}
	return ""
}

func (x *SendMessageReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SendMessageReq) GetMsgType() int32 {
	if x != nil {
		return x.MsgType
	}
	return 0
}

func (x *SendMessageReq) GetManual() int32 {
	if x != nil {
		return x.Manual
	}
	return 0
}

func (x *SendMessageReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendMessageReq) GetContent() map[string]string {
	if x != nil {
		return x.Content
	}
	return nil
}

var File_MsgPushMessage_proto protoreflect.FileDescriptor

var file_MsgPushMessage_proto_rawDesc = []byte{
	0x0a, 0x14, 0x4d, 0x73, 0x67, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaa, 0x02, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x42,
	0x79, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4e,
	0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x22,
	0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x57, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x79, 0x44, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x6c, 0x6b,
	0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x2e,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x86, 0x02, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f,
	0x62, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x53, 0x6d, 0x73, 0x52, 0x65, 0x71, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb0, 0x01, 0x0a,
	0x0c, 0x53, 0x4d, 0x53, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x53,
	0x4d, 0x53, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x40, 0x0a,
	0x12, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x8c, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d, 0x73, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x53, 0x4d, 0x53, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x52,
	0x0a, 0x08, 0x4d, 0x61, 0x69, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0xe0, 0x03, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x63, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x02, 0x63, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x63, 0x63, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x03, 0x62, 0x63, 0x63, 0x12, 0x45, 0x0a, 0x0d, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x42, 0x0a,
	0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x09, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x61, 0x69, 0x6c, 0x52, 0x65,
	0x71, 0x2e, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x12, 0x1f, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x09, 0x2e, 0x4d, 0x61, 0x69, 0x6c, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x1a, 0x40, 0x0a, 0x12, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xab, 0x01, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6d,
	0x73, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61,
	0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e,
	0x63, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x65, 0x61, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x99, 0x01, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a,
	0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x41, 0x72, 0x65,
	0x61, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x6d, 0x6f, 0x62,
	0x69, 0x6c, 0x65, 0x41, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x76,
	0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x9e, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x05, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x73, 0x3b, 0x72, 0x70, 0x63, 0x41, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_MsgPushMessage_proto_rawDescOnce sync.Once
	file_MsgPushMessage_proto_rawDescData = file_MsgPushMessage_proto_rawDesc
)

func file_MsgPushMessage_proto_rawDescGZIP() []byte {
	file_MsgPushMessage_proto_rawDescOnce.Do(func() {
		file_MsgPushMessage_proto_rawDescData = protoimpl.X.CompressGZIP(file_MsgPushMessage_proto_rawDescData)
	})
	return file_MsgPushMessage_proto_rawDescData
}

var file_MsgPushMessage_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_MsgPushMessage_proto_goTypes = []interface{}{
	(*SendByDingTalkSimpleNoticeReq)(nil), // 0: SendByDingTalkSimpleNoticeReq
	(*SendSmsReq)(nil),                    // 1: SendSmsReq
	(*SMSBatchSend)(nil),                  // 2: SMSBatchSend
	(*SendSmsBatchReq)(nil),               // 3: SendSmsBatchReq
	(*MailFile)(nil),                      // 4: MailFile
	(*SendMailReq)(nil),                   // 5: SendMailReq
	(*SendSmsVerCodeReq)(nil),             // 6: SendSmsVerCodeReq
	(*CheckMobileVerCodeReq)(nil),         // 7: CheckMobileVerCodeReq
	(*SendMessageReq)(nil),                // 8: SendMessageReq
	nil,                                   // 9: SendByDingTalkSimpleNoticeReq.TemplateParamEntry
	nil,                                   // 10: SendSmsReq.TemplateParamEntry
	nil,                                   // 11: SMSBatchSend.TemplateParamEntry
	nil,                                   // 12: SendMailReq.TemplateParamEntry
	nil,                                   // 13: SendMailReq.SubjectParamEntry
	nil,                                   // 14: SendMessageReq.ContentEntry
}
var file_MsgPushMessage_proto_depIdxs = []int32{
	9,  // 0: SendByDingTalkSimpleNoticeReq.templateParam:type_name -> SendByDingTalkSimpleNoticeReq.TemplateParamEntry
	10, // 1: SendSmsReq.templateParam:type_name -> SendSmsReq.TemplateParamEntry
	11, // 2: SMSBatchSend.templateParam:type_name -> SMSBatchSend.TemplateParamEntry
	2,  // 3: SendSmsBatchReq.data:type_name -> SMSBatchSend
	12, // 4: SendMailReq.templateParam:type_name -> SendMailReq.TemplateParamEntry
	13, // 5: SendMailReq.subjectParam:type_name -> SendMailReq.SubjectParamEntry
	4,  // 6: SendMailReq.files:type_name -> MailFile
	14, // 7: SendMessageReq.content:type_name -> SendMessageReq.ContentEntry
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_MsgPushMessage_proto_init() }
func file_MsgPushMessage_proto_init() {
	if File_MsgPushMessage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_MsgPushMessage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendByDingTalkSimpleNoticeReq); i {
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
		file_MsgPushMessage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsReq); i {
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
		file_MsgPushMessage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SMSBatchSend); i {
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
		file_MsgPushMessage_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsBatchReq); i {
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
		file_MsgPushMessage_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MailFile); i {
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
		file_MsgPushMessage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMailReq); i {
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
		file_MsgPushMessage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSmsVerCodeReq); i {
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
		file_MsgPushMessage_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckMobileVerCodeReq); i {
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
		file_MsgPushMessage_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageReq); i {
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
			RawDescriptor: file_MsgPushMessage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MsgPushMessage_proto_goTypes,
		DependencyIndexes: file_MsgPushMessage_proto_depIdxs,
		MessageInfos:      file_MsgPushMessage_proto_msgTypes,
	}.Build()
	File_MsgPushMessage_proto = out.File
	file_MsgPushMessage_proto_rawDesc = nil
	file_MsgPushMessage_proto_goTypes = nil
	file_MsgPushMessage_proto_depIdxs = nil
}
