// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.0
// source: message.proto

package pb

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type MMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgType   MSG_TYPE      `protobuf:"varint,1,opt,name=msgType,proto3,enum=matchproto.MSG_TYPE" json:"msgType,omitempty"`      //消息类型
	MsgCode   GAME_MSG_CODE `protobuf:"varint,2,opt,name=msgCode,proto3,enum=matchproto.GAME_MSG_CODE" json:"msgCode,omitempty"` //消息码 用于表示具体业务类型 参照msgenum.proto中的GAME_MSG_CODE枚举
	SessionId int32         `protobuf:"varint,3,opt,name=sessionId,proto3" json:"sessionId,omitempty"`                           //会话ID
	SeqId     int32         `protobuf:"varint,4,opt,name=seqId,proto3" json:"seqId,omitempty"`                                   //消息的编号
	Notify    *Notify       `protobuf:"bytes,5,opt,name=notify,proto3" json:"notify,omitempty"`                                  //通知类型
	Request   *Request      `protobuf:"bytes,6,opt,name=request,proto3" json:"request,omitempty"`                                //请求类型
	Response  *Response     `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty"`                              //答复类型
	SendTime  int64         `protobuf:"varint,8,opt,name=sendTime,proto3" json:"sendTime,omitempty"`                             //消息发送时间，用于测试网络时延
}

func (x *MMessage) Reset() {
	*x = MMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MMessage) ProtoMessage() {}

func (x *MMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MMessage.ProtoReflect.Descriptor instead.
func (*MMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *MMessage) GetMsgType() MSG_TYPE {
	if x != nil {
		return x.MsgType
	}
	return MSG_TYPE_NOTIFY
}

func (x *MMessage) GetMsgCode() GAME_MSG_CODE {
	if x != nil {
		return x.MsgCode
	}
	return GAME_MSG_CODE_PLAYER_MATCHING_REQUEST
}

func (x *MMessage) GetSessionId() int32 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *MMessage) GetSeqId() int32 {
	if x != nil {
		return x.SeqId
	}
	return 0
}

func (x *MMessage) GetNotify() *Notify {
	if x != nil {
		return x.Notify
	}
	return nil
}

func (x *MMessage) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *MMessage) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *MMessage) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x6d, 0x73, 0x67, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x08, 0x4d, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4d, 0x53, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x52, 0x07, 0x6d, 0x73, 0x67,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x4d, 0x53, 0x47, 0x5f, 0x43, 0x4f, 0x44, 0x45,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x65, 0x71, 0x49, 0x64, 0x12, 0x2a, 0x0a,
	0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x12, 0x2d, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_message_proto_goTypes = []interface{}{
	(*MMessage)(nil),   // 0: matchproto.MMessage
	(MSG_TYPE)(0),      // 1: matchproto.MSG_TYPE
	(GAME_MSG_CODE)(0), // 2: matchproto.GAME_MSG_CODE
	(*Notify)(nil),     // 3: matchproto.Notify
	(*Request)(nil),    // 4: matchproto.Request
	(*Response)(nil),   // 5: matchproto.Response
}
var file_message_proto_depIdxs = []int32{
	1, // 0: matchproto.MMessage.msgType:type_name -> matchproto.MSG_TYPE
	2, // 1: matchproto.MMessage.msgCode:type_name -> matchproto.GAME_MSG_CODE
	3, // 2: matchproto.MMessage.notify:type_name -> matchproto.Notify
	4, // 3: matchproto.MMessage.request:type_name -> matchproto.Request
	5, // 4: matchproto.MMessage.response:type_name -> matchproto.Response
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	file_notify_proto_init()
	file_request_proto_init()
	file_response_proto_init()
	file_msgenum_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MMessage); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
