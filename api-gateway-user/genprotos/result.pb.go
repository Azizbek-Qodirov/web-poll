// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/result.proto

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

type CreateResultReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PollId string `protobuf:"bytes,2,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *CreateResultReq) Reset() {
	*x = CreateResultReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResultReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResultReq) ProtoMessage() {}

func (x *CreateResultReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResultReq.ProtoReflect.Descriptor instead.
func (*CreateResultReq) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{0}
}

func (x *CreateResultReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateResultReq) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

type SavePollAnswerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultId    string `protobuf:"bytes,1,opt,name=result_id,json=resultId,proto3" json:"result_id,omitempty"`
	QuestionNum int32  `protobuf:"varint,2,opt,name=question_num,json=questionNum,proto3" json:"question_num,omitempty"`
	Answer      int32  `protobuf:"varint,3,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *SavePollAnswerReq) Reset() {
	*x = SavePollAnswerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SavePollAnswerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SavePollAnswerReq) ProtoMessage() {}

func (x *SavePollAnswerReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SavePollAnswerReq.ProtoReflect.Descriptor instead.
func (*SavePollAnswerReq) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{1}
}

func (x *SavePollAnswerReq) GetResultId() string {
	if x != nil {
		return x.ResultId
	}
	return ""
}

func (x *SavePollAnswerReq) GetQuestionNum() int32 {
	if x != nil {
		return x.QuestionNum
	}
	return 0
}

func (x *SavePollAnswerReq) GetAnswer() int32 {
	if x != nil {
		return x.Answer
	}
	return 0
}

type GetPoll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId string `protobuf:"bytes,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *GetPoll) Reset() {
	*x = GetPoll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_result_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPoll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPoll) ProtoMessage() {}

func (x *GetPoll) ProtoReflect() protoreflect.Message {
	mi := &file_protos_result_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPoll.ProtoReflect.Descriptor instead.
func (*GetPoll) Descriptor() ([]byte, []int) {
	return file_protos_result_proto_rawDescGZIP(), []int{2}
}

func (x *GetPoll) GetPollId() string {
	if x != nil {
		return x.PollId
	}
	return ""
}

var File_protos_result_proto protoreflect.FileDescriptor

var file_protos_result_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x1a, 0x11,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x6f, 0x69, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x43, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x22, 0x6b, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f,
	0x6c, 0x6c, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6e, 0x73,
	0x77, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x17,
	0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x32, 0x85, 0x01, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x6f, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f,
	0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x53, 0x61, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x12, 0x1a, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53,
	0x61, 0x76, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x0d, 0x2e, 0x70, 0x6f, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x42,
	0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_result_proto_rawDescOnce sync.Once
	file_protos_result_proto_rawDescData = file_protos_result_proto_rawDesc
)

func file_protos_result_proto_rawDescGZIP() []byte {
	file_protos_result_proto_rawDescOnce.Do(func() {
		file_protos_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_result_proto_rawDescData)
	})
	return file_protos_result_proto_rawDescData
}

var file_protos_result_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_result_proto_goTypes = []any{
	(*CreateResultReq)(nil),   // 0: polling.CreateResultReq
	(*SavePollAnswerReq)(nil), // 1: polling.SavePollAnswerReq
	(*GetPoll)(nil),           // 2: polling.GetPoll
	(*Void)(nil),              // 3: polling.Void
}
var file_protos_result_proto_depIdxs = []int32{
	0, // 0: polling.ResultService.CreateResult:input_type -> polling.CreateResultReq
	1, // 1: polling.ResultService.SavePollAnswer:input_type -> polling.SavePollAnswerReq
	3, // 2: polling.ResultService.CreateResult:output_type -> polling.Void
	3, // 3: polling.ResultService.SavePollAnswer:output_type -> polling.Void
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_result_proto_init() }
func file_protos_result_proto_init() {
	if File_protos_result_proto != nil {
		return
	}
	file_protos_void_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_result_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResultReq); i {
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
		file_protos_result_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SavePollAnswerReq); i {
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
		file_protos_result_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetPoll); i {
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
			RawDescriptor: file_protos_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_result_proto_goTypes,
		DependencyIndexes: file_protos_result_proto_depIdxs,
		MessageInfos:      file_protos_result_proto_msgTypes,
	}.Build()
	File_protos_result_proto = out.File
	file_protos_result_proto_rawDesc = nil
	file_protos_result_proto_goTypes = nil
	file_protos_result_proto_depIdxs = nil
}
