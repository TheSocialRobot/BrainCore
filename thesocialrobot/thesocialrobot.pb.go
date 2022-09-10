// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: thesocialrobot/thesocialrobot.proto

package thesocialrobot

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

// event from robot (the body) containing current state
type BodyEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BodyEvent) Reset() {
	*x = BodyEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyEvent) ProtoMessage() {}

func (x *BodyEvent) ProtoReflect() protoreflect.Message {
	mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyEvent.ProtoReflect.Descriptor instead.
func (*BodyEvent) Descriptor() ([]byte, []int) {
	return file_thesocialrobot_thesocialrobot_proto_rawDescGZIP(), []int{0}
}

func (x *BodyEvent) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Say struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Say) Reset() {
	*x = Say{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Say) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Say) ProtoMessage() {}

func (x *Say) ProtoReflect() protoreflect.Message {
	mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Say.ProtoReflect.Descriptor instead.
func (*Say) Descriptor() ([]byte, []int) {
	return file_thesocialrobot_thesocialrobot_proto_rawDescGZIP(), []int{1}
}

func (x *Say) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delay in milliseconds before triggering the action
	// could use the Duration type here, but don't think we need nanosecond precision and the second/nanosecond split complicates things
	Delay int32 `protobuf:"varint,1,opt,name=delay,proto3" json:"delay,omitempty"`
	// Types that are assignable to Action:
	//
	//	*Action_Say
	Action isAction_Action `protobuf_oneof:"action"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_thesocialrobot_thesocialrobot_proto_rawDescGZIP(), []int{2}
}

func (x *Action) GetDelay() int32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (m *Action) GetAction() isAction_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (x *Action) GetSay() *Say {
	if x, ok := x.GetAction().(*Action_Say); ok {
		return x.Say
	}
	return nil
}

type isAction_Action interface {
	isAction_Action()
}

type Action_Say struct {
	Say *Say `protobuf:"bytes,2,opt,name=say,proto3,oneof"`
}

func (*Action_Say) isAction_Action() {}

// message from brain to body, instructing the body do take one or more actions
type BodyCommand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Actions []*Action `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *BodyCommand) Reset() {
	*x = BodyCommand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BodyCommand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BodyCommand) ProtoMessage() {}

func (x *BodyCommand) ProtoReflect() protoreflect.Message {
	mi := &file_thesocialrobot_thesocialrobot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BodyCommand.ProtoReflect.Descriptor instead.
func (*BodyCommand) Descriptor() ([]byte, []int) {
	return file_thesocialrobot_thesocialrobot_proto_rawDescGZIP(), []int{3}
}

func (x *BodyCommand) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BodyCommand) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_thesocialrobot_thesocialrobot_proto protoreflect.FileDescriptor

var file_thesocialrobot_thesocialrobot_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x72, 0x6f, 0x62, 0x6f, 0x74,
	0x2f, 0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c,
	0x72, 0x6f, 0x62, 0x6f, 0x74, 0x22, 0x1b, 0x0a, 0x09, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x19, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x51, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x12, 0x27, 0x0a,
	0x03, 0x73, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x74, 0x68, 0x65,
	0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x53, 0x61, 0x79, 0x48,
	0x00, 0x52, 0x03, 0x73, 0x61, 0x79, 0x42, 0x08, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x4f, 0x0a, 0x0b, 0x42, 0x6f, 0x64, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x30, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x16, 0x2e, 0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x72, 0x6f, 0x62, 0x6f,
	0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x32, 0x5d, 0x0a, 0x0e, 0x54, 0x68, 0x65, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x12, 0x19, 0x2e, 0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x72, 0x6f,
	0x62, 0x6f, 0x74, 0x2e, 0x42, 0x6f, 0x64, 0x79, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x1b, 0x2e,
	0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2e, 0x42,
	0x6f, 0x64, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54,
	0x68, 0x65, 0x53, 0x6f, 0x63, 0x69, 0x61, 0x6c, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x2f, 0x42, 0x72,
	0x61, 0x69, 0x6e, 0x43, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x68, 0x65, 0x73, 0x6f, 0x63, 0x69, 0x61,
	0x6c, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_thesocialrobot_thesocialrobot_proto_rawDescOnce sync.Once
	file_thesocialrobot_thesocialrobot_proto_rawDescData = file_thesocialrobot_thesocialrobot_proto_rawDesc
)

func file_thesocialrobot_thesocialrobot_proto_rawDescGZIP() []byte {
	file_thesocialrobot_thesocialrobot_proto_rawDescOnce.Do(func() {
		file_thesocialrobot_thesocialrobot_proto_rawDescData = protoimpl.X.CompressGZIP(file_thesocialrobot_thesocialrobot_proto_rawDescData)
	})
	return file_thesocialrobot_thesocialrobot_proto_rawDescData
}

var file_thesocialrobot_thesocialrobot_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_thesocialrobot_thesocialrobot_proto_goTypes = []interface{}{
	(*BodyEvent)(nil),   // 0: thesocialrobot.BodyEvent
	(*Say)(nil),         // 1: thesocialrobot.Say
	(*Action)(nil),      // 2: thesocialrobot.Action
	(*BodyCommand)(nil), // 3: thesocialrobot.BodyCommand
}
var file_thesocialrobot_thesocialrobot_proto_depIdxs = []int32{
	1, // 0: thesocialrobot.Action.say:type_name -> thesocialrobot.Say
	2, // 1: thesocialrobot.BodyCommand.actions:type_name -> thesocialrobot.Action
	0, // 2: thesocialrobot.TheSocialRobot.EventStream:input_type -> thesocialrobot.BodyEvent
	3, // 3: thesocialrobot.TheSocialRobot.EventStream:output_type -> thesocialrobot.BodyCommand
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_thesocialrobot_thesocialrobot_proto_init() }
func file_thesocialrobot_thesocialrobot_proto_init() {
	if File_thesocialrobot_thesocialrobot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_thesocialrobot_thesocialrobot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyEvent); i {
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
		file_thesocialrobot_thesocialrobot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Say); i {
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
		file_thesocialrobot_thesocialrobot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_thesocialrobot_thesocialrobot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BodyCommand); i {
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
	file_thesocialrobot_thesocialrobot_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Action_Say)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_thesocialrobot_thesocialrobot_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_thesocialrobot_thesocialrobot_proto_goTypes,
		DependencyIndexes: file_thesocialrobot_thesocialrobot_proto_depIdxs,
		MessageInfos:      file_thesocialrobot_thesocialrobot_proto_msgTypes,
	}.Build()
	File_thesocialrobot_thesocialrobot_proto = out.File
	file_thesocialrobot_thesocialrobot_proto_rawDesc = nil
	file_thesocialrobot_thesocialrobot_proto_goTypes = nil
	file_thesocialrobot_thesocialrobot_proto_depIdxs = nil
}
