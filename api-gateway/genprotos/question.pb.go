// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: protos/question.proto

package genprotos

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

type QuestionCreateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	PollId  string `protobuf:"bytes,4,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *QuestionCreateReq) Reset() {
	*x = QuestionCreateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_question_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionCreateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionCreateReq) ProtoMessage() {}

func (x *QuestionCreateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_question_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionCreateReq.ProtoReflect.Descriptor instead.
func (*QuestionCreateReq) Descriptor() ([]byte, []int) {
	return file_protos_question_proto_rawDescGZIP(), []int{0}
}

func (x *QuestionCreateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *QuestionCreateReq) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type QuestionUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *QuestionUpdateReq) Reset() {
	*x = QuestionUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_question_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionUpdateReq) ProtoMessage() {}

func (x *QuestionUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_question_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionUpdateReq.ProtoReflect.Descriptor instead.
func (*QuestionUpdateReq) Descriptor() ([]byte, []int) {
	return file_protos_question_proto_rawDescGZIP(), []int{1}
}

func (x *QuestionUpdateReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QuestionUpdateReq) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type QuestionGetByIDRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	PollId  string `protobuf:"bytes,3,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *QuestionGetByIDRes) Reset() {
	*x = QuestionGetByIDRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_question_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionGetByIDRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionGetByIDRes) ProtoMessage() {}

func (x *QuestionGetByIDRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_question_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionGetByIDRes.ProtoReflect.Descriptor instead.
func (*QuestionGetByIDRes) Descriptor() ([]byte, []int) {
	return file_protos_question_proto_rawDescGZIP(), []int{2}
}

func (x *QuestionGetByIDRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *QuestionGetByIDRes) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *QuestionGetByIDRes) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type QuestionGetAllRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Question []*Question     `protobuf:"bytes,1,rep,name=question,proto3" json:"question,omitempty"`
	Poll     *PollGetByIDRes `protobuf:"bytes,2,opt,name=poll,proto3" json:"poll,omitempty"`
}

func (x *QuestionGetAllRes) Reset() {
	*x = QuestionGetAllRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_question_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionGetAllRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionGetAllRes) ProtoMessage() {}

func (x *QuestionGetAllRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_question_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionGetAllRes.ProtoReflect.Descriptor instead.
func (*QuestionGetAllRes) Descriptor() ([]byte, []int) {
	return file_protos_question_proto_rawDescGZIP(), []int{3}
}

func (x *QuestionGetAllRes) GetQuestion() []*Question {
	if x != nil {
		return x.Question
	}
	return nil
}

func (x *QuestionGetAllRes) GetPoll() *PollGetByIDRes {
	if x != nil {
		return x.Poll
	}
	return nil
}

type QuestionGetAllReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId string `protobuf:"bytes,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *QuestionGetAllReq) Reset() {
	*x = QuestionGetAllReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_question_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuestionGetAllReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuestionGetAllReq) ProtoMessage() {}

func (x *QuestionGetAllReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_question_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuestionGetAllReq.ProtoReflect.Descriptor instead.
func (*QuestionGetAllReq) Descriptor() ([]byte, []int) {
	return file_protos_question_proto_rawDescGZIP(), []int{4}
}

func (x *QuestionGetAllReq) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type Question struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Num     int32  `protobuf:"varint,2,opt,name=num,proto3" json:"num,omitempty"`
	Content string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	PollId  string `protobuf:"bytes,4,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *Question) Reset() {
	*x = Question{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_question_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Question) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Question) ProtoMessage() {}

func (x *Question) ProtoReflect() protoreflect.Message {
	mi := &file_protos_question_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Question.ProtoReflect.Descriptor instead.
func (*Question) Descriptor() ([]byte, []int) {
	return file_protos_question_proto_rawDescGZIP(), []int{5}
}

func (x *Question) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Question) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

func (x *Question) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Question) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

var File_protos_question_proto protoreflect.FileDescriptor

var file_protos_question_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x3d,
	0x0a, 0x11, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x57, 0x0a,
	0x12, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x04, 0x70, 0x6f,
	0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x52, 0x04, 0x70, 0x6f, 0x6c, 0x6c, 0x22, 0x2c, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07,
	0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x08, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x6e, 0x75, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x32, 0x9c, 0x02, 0x0a, 0x0f, 0x51, 0x75, 0x65, 0x73, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x51,
	0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x33, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x6f, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0d,
	0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x0d, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x07,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x2e, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x1b, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x1a, 0x2e,
	0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x51, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_question_proto_rawDescOnce sync.Once
	file_protos_question_proto_rawDescData = file_protos_question_proto_rawDesc
)

func file_protos_question_proto_rawDescGZIP() []byte {
	file_protos_question_proto_rawDescOnce.Do(func() {
		file_protos_question_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_question_proto_rawDescData)
	})
	return file_protos_question_proto_rawDescData
}

var file_protos_question_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_question_proto_goTypes = []any{
	(*QuestionCreateReq)(nil),  // 0: polling.QuestionCreateReq
	(*QuestionUpdateReq)(nil),  // 1: polling.QuestionUpdateReq
	(*QuestionGetByIDRes)(nil), // 2: polling.QuestionGetByIDRes
	(*QuestionGetAllRes)(nil),  // 3: polling.QuestionGetAllRes
	(*QuestionGetAllReq)(nil),  // 4: polling.QuestionGetAllReq
	(*Question)(nil),           // 5: polling.Question
	(*PollGetByIDRes)(nil),     // 6: polling.PollGetByIDRes
	(*ByID)(nil),               // 7: polling.ByID
	(*Void)(nil),               // 8: polling.Void
}
var file_protos_question_proto_depIdxs = []int32{
	5, // 0: polling.QuestionGetAllRes.question:type_name -> polling.Question
	6, // 1: polling.QuestionGetAllRes.poll:type_name -> polling.PollGetByIDRes
	0, // 2: polling.QuestionService.Create:input_type -> polling.QuestionCreateReq
	1, // 3: polling.QuestionService.Update:input_type -> polling.QuestionUpdateReq
	7, // 4: polling.QuestionService.Delete:input_type -> polling.ByID
	7, // 5: polling.QuestionService.GetByID:input_type -> polling.ByID
	4, // 6: polling.QuestionService.GetAll:input_type -> polling.QuestionGetAllReq
	8, // 7: polling.QuestionService.Create:output_type -> polling.Void
	8, // 8: polling.QuestionService.Update:output_type -> polling.Void
	8, // 9: polling.QuestionService.Delete:output_type -> polling.Void
	2, // 10: polling.QuestionService.GetByID:output_type -> polling.QuestionGetByIDRes
	3, // 11: polling.QuestionService.GetAll:output_type -> polling.QuestionGetAllRes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_question_proto_init() }
func file_protos_question_proto_init() {
	if File_protos_question_proto != nil {
		return
	}
	file_protos_void_proto_init()
	file_protos_poll_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_question_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*QuestionCreateReq); i {
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
		file_protos_question_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*QuestionUpdateReq); i {
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
		file_protos_question_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*QuestionGetByIDRes); i {
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
		file_protos_question_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*QuestionGetAllRes); i {
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
		file_protos_question_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*QuestionGetAllReq); i {
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
		file_protos_question_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*Question); i {
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
			RawDescriptor: file_protos_question_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_question_proto_goTypes,
		DependencyIndexes: file_protos_question_proto_depIdxs,
		MessageInfos:      file_protos_question_proto_msgTypes,
	}.Build()
	File_protos_question_proto = out.File
	file_protos_question_proto_rawDesc = nil
	file_protos_question_proto_goTypes = nil
	file_protos_question_proto_depIdxs = nil
}
