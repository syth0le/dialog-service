// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: internalapi/dialog_service.proto

package internalapi

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateDialogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	ParticipantsIds []string `protobuf:"bytes,2,rep,name=participantsIds,proto3" json:"participantsIds,omitempty"`
}

func (x *CreateDialogRequest) Reset() {
	*x = CreateDialogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_dialog_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDialogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDialogRequest) ProtoMessage() {}

func (x *CreateDialogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_dialog_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDialogRequest.ProtoReflect.Descriptor instead.
func (*CreateDialogRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_dialog_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateDialogRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateDialogRequest) GetParticipantsIds() []string {
	if x != nil {
		return x.ParticipantsIds
	}
	return nil
}

type CreateDialogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogId        string   `protobuf:"bytes,1,opt,name=dialogId,proto3" json:"dialogId,omitempty"`
	ParticipantsIds []string `protobuf:"bytes,2,rep,name=participantsIds,proto3" json:"participantsIds,omitempty"`
}

func (x *CreateDialogResponse) Reset() {
	*x = CreateDialogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_dialog_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDialogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDialogResponse) ProtoMessage() {}

func (x *CreateDialogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_dialog_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDialogResponse.ProtoReflect.Descriptor instead.
func (*CreateDialogResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_dialog_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDialogResponse) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *CreateDialogResponse) GetParticipantsIds() []string {
	if x != nil {
		return x.ParticipantsIds
	}
	return nil
}

type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DialogId string `protobuf:"bytes,1,opt,name=dialogId,proto3" json:"dialogId,omitempty"`
	SenderId string `protobuf:"bytes,2,opt,name=senderId,proto3" json:"senderId,omitempty"`
	Text     string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_dialog_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_dialog_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_dialog_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateMessageRequest) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *CreateMessageRequest) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *CreateMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetDialogMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	DialogId string `protobuf:"bytes,2,opt,name=dialogId,proto3" json:"dialogId,omitempty"`
}

func (x *GetDialogMessagesRequest) Reset() {
	*x = GetDialogMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_dialog_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialogMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialogMessagesRequest) ProtoMessage() {}

func (x *GetDialogMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_dialog_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialogMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetDialogMessagesRequest) Descriptor() ([]byte, []int) {
	return file_internalapi_dialog_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetDialogMessagesRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetDialogMessagesRequest) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	DialogId string `protobuf:"bytes,2,opt,name=dialogId,proto3" json:"dialogId,omitempty"`
	SenderId string `protobuf:"bytes,3,opt,name=senderId,proto3" json:"senderId,omitempty"`
	Text     string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_dialog_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_dialog_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_internalapi_dialog_service_proto_rawDescGZIP(), []int{4}
}

func (x *Message) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Message) GetDialogId() string {
	if x != nil {
		return x.DialogId
	}
	return ""
}

func (x *Message) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetDialogMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *GetDialogMessagesResponse) Reset() {
	*x = GetDialogMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internalapi_dialog_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDialogMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDialogMessagesResponse) ProtoMessage() {}

func (x *GetDialogMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internalapi_dialog_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDialogMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetDialogMessagesResponse) Descriptor() ([]byte, []int) {
	return file_internalapi_dialog_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetDialogMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

var File_internalapi_dialog_service_proto protoreflect.FileDescriptor

var file_internalapi_dialog_service_proto_rawDesc = []byte{
	0x0a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1a, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x1a, 0x1b,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74,
	0x73, 0x49, 0x64, 0x73, 0x22, 0x5c, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x49,
	0x64, 0x73, 0x22, 0x62, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61,
	0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69,
	0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x5c, 0x0a,
	0x19, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x73,
	0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x32, 0xe6, 0x02, 0x0a, 0x0d,
	0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x12, 0x2f, 0x2e,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x5b, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x30, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x82, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x34, 0x2e, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61,
	0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c, 0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x6f,
	0x63, 0x69, 0x61, 0x6c, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x69, 0x61, 0x6c,
	0x6f, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x79, 0x74, 0x68, 0x30, 0x6c, 0x65, 0x2f, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x2d, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internalapi_dialog_service_proto_rawDescOnce sync.Once
	file_internalapi_dialog_service_proto_rawDescData = file_internalapi_dialog_service_proto_rawDesc
)

func file_internalapi_dialog_service_proto_rawDescGZIP() []byte {
	file_internalapi_dialog_service_proto_rawDescOnce.Do(func() {
		file_internalapi_dialog_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_internalapi_dialog_service_proto_rawDescData)
	})
	return file_internalapi_dialog_service_proto_rawDescData
}

var file_internalapi_dialog_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internalapi_dialog_service_proto_goTypes = []interface{}{
	(*CreateDialogRequest)(nil),       // 0: social_network.internalapi.CreateDialogRequest
	(*CreateDialogResponse)(nil),      // 1: social_network.internalapi.CreateDialogResponse
	(*CreateMessageRequest)(nil),      // 2: social_network.internalapi.CreateMessageRequest
	(*GetDialogMessagesRequest)(nil),  // 3: social_network.internalapi.GetDialogMessagesRequest
	(*Message)(nil),                   // 4: social_network.internalapi.Message
	(*GetDialogMessagesResponse)(nil), // 5: social_network.internalapi.GetDialogMessagesResponse
	(*emptypb.Empty)(nil),             // 6: google.protobuf.Empty
}
var file_internalapi_dialog_service_proto_depIdxs = []int32{
	4, // 0: social_network.internalapi.GetDialogMessagesResponse.messages:type_name -> social_network.internalapi.Message
	0, // 1: social_network.internalapi.DialogService.CreateDialog:input_type -> social_network.internalapi.CreateDialogRequest
	2, // 2: social_network.internalapi.DialogService.CreateMessage:input_type -> social_network.internalapi.CreateMessageRequest
	3, // 3: social_network.internalapi.DialogService.GetDialogMessages:input_type -> social_network.internalapi.GetDialogMessagesRequest
	1, // 4: social_network.internalapi.DialogService.CreateDialog:output_type -> social_network.internalapi.CreateDialogResponse
	6, // 5: social_network.internalapi.DialogService.CreateMessage:output_type -> google.protobuf.Empty
	5, // 6: social_network.internalapi.DialogService.GetDialogMessages:output_type -> social_network.internalapi.GetDialogMessagesResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internalapi_dialog_service_proto_init() }
func file_internalapi_dialog_service_proto_init() {
	if File_internalapi_dialog_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internalapi_dialog_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDialogRequest); i {
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
		file_internalapi_dialog_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDialogResponse); i {
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
		file_internalapi_dialog_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMessageRequest); i {
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
		file_internalapi_dialog_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialogMessagesRequest); i {
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
		file_internalapi_dialog_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_internalapi_dialog_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDialogMessagesResponse); i {
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
			RawDescriptor: file_internalapi_dialog_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internalapi_dialog_service_proto_goTypes,
		DependencyIndexes: file_internalapi_dialog_service_proto_depIdxs,
		MessageInfos:      file_internalapi_dialog_service_proto_msgTypes,
	}.Build()
	File_internalapi_dialog_service_proto = out.File
	file_internalapi_dialog_service_proto_rawDesc = nil
	file_internalapi_dialog_service_proto_goTypes = nil
	file_internalapi_dialog_service_proto_depIdxs = nil
}
