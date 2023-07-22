// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: message.proto

package goshare

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

type MessageType int32

const (
	// 订阅行情
	MessageType_REQ_SUBSCRIBE_MARKET_DATA MessageType = 0
	// 订阅结果
	MessageType_RSP_SUBSCRIBE_MARKET_DATA MessageType = 1
	// 有新行情
	MessageType_RTN_MARKET_DATA_UPDATE MessageType = 2
	// 订阅行情
	MessageType_REQ_UNSUBSCRIBE_MARKET_DATA MessageType = 3
	// 订阅结果
	MessageType_RSP_UNSUBSCRIBE_MARKET_DATA MessageType = 4
	// 心跳
	MessageType_HEATBEAT MessageType = 5
	// 上传tick
	MessageType_UPLOAD_TICK MessageType = 6
	// 策略委托
	MessageType_SOPT_STRATEGY_ORDER MessageType = 7
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "REQ_SUBSCRIBE_MARKET_DATA",
		1: "RSP_SUBSCRIBE_MARKET_DATA",
		2: "RTN_MARKET_DATA_UPDATE",
		3: "REQ_UNSUBSCRIBE_MARKET_DATA",
		4: "RSP_UNSUBSCRIBE_MARKET_DATA",
		5: "HEATBEAT",
		6: "UPLOAD_TICK",
		7: "SOPT_STRATEGY_ORDER",
	}
	MessageType_value = map[string]int32{
		"REQ_SUBSCRIBE_MARKET_DATA":   0,
		"RSP_SUBSCRIBE_MARKET_DATA":   1,
		"RTN_MARKET_DATA_UPDATE":      2,
		"REQ_UNSUBSCRIBE_MARKET_DATA": 3,
		"RSP_UNSUBSCRIBE_MARKET_DATA": 4,
		"HEATBEAT":                    5,
		"UPLOAD_TICK":                 6,
		"SOPT_STRATEGY_ORDER":         7,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_message_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_message_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

// Message 消息
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=goshare.MessageType" json:"type,omitempty"`
	Target string      `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Data   []byte      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_REQ_SUBSCRIBE_MARKET_DATA
}

func (x *Message) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *Message) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// MessageList 消息list
type MessageList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Message `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *MessageList) Reset() {
	*x = MessageList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageList) ProtoMessage() {}

func (x *MessageList) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageList.ProtoReflect.Descriptor instead.
func (*MessageList) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *MessageList) GetList() []*Message {
	if x != nil {
		return x.List
	}
	return nil
}

type StreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=goshare.MessageType" json:"type,omitempty"`
	Target string      `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Data   string      `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *StreamMessage) Reset() {
	*x = StreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessage) ProtoMessage() {}

func (x *StreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessage.ProtoReflect.Descriptor instead.
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *StreamMessage) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_REQ_SUBSCRIBE_MARKET_DATA
}

func (x *StreamMessage) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

func (x *StreamMessage) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x67, 0x6f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x22, 0x5f, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x0b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x65,
	0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x2a, 0xe1, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x51, 0x5f, 0x53, 0x55, 0x42,
	0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x5f, 0x44, 0x41,
	0x54, 0x41, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x53, 0x50, 0x5f, 0x53, 0x55, 0x42, 0x53,
	0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x10, 0x01, 0x12, 0x1a, 0x0a, 0x16, 0x52, 0x54, 0x4e, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45,
	0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x02, 0x12,
	0x1f, 0x0a, 0x1b, 0x52, 0x45, 0x51, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52, 0x49,
	0x42, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10, 0x03,
	0x12, 0x1f, 0x0a, 0x1b, 0x52, 0x53, 0x50, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x42, 0x53, 0x43, 0x52,
	0x49, 0x42, 0x45, 0x5f, 0x4d, 0x41, 0x52, 0x4b, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x10,
	0x04, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x45, 0x41, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x05, 0x12,
	0x0f, 0x0a, 0x0b, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x54, 0x49, 0x43, 0x4b, 0x10, 0x06,
	0x12, 0x17, 0x0a, 0x13, 0x53, 0x4f, 0x50, 0x54, 0x5f, 0x53, 0x54, 0x52, 0x41, 0x54, 0x45, 0x47,
	0x59, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x07, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x61, 0x76, 0x65, 0x6e, 0x48, 0x75, 0x6f,
	0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x6f, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_message_proto_goTypes = []interface{}{
	(MessageType)(0),      // 0: goshare.MessageType
	(*Message)(nil),       // 1: goshare.Message
	(*MessageList)(nil),   // 2: goshare.MessageList
	(*StreamMessage)(nil), // 3: goshare.StreamMessage
}
var file_message_proto_depIdxs = []int32{
	0, // 0: goshare.Message.type:type_name -> goshare.MessageType
	1, // 1: goshare.MessageList.list:type_name -> goshare.Message
	0, // 2: goshare.StreamMessage.type:type_name -> goshare.MessageType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageList); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessage); i {
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
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		EnumInfos:         file_message_proto_enumTypes,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}
