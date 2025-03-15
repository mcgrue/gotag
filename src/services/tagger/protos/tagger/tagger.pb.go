package tagger

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type UnstructuredText struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnstructuredEntry string `protobuf:"bytes,1,opt,name=unstructured_entry,json=unstructuredEntry,proto3" json:"unstructured_entry,omitempty"`
}

func (x *UnstructuredText) Reset() {
	*x = UnstructuredText{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_tagger_tagger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnstructuredText) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnstructuredText) ProtoMessage() {}

func (x *UnstructuredText) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tagger_tagger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnstructuredText.ProtoReflect.Descriptor instead.
func (*UnstructuredText) Descriptor() ([]byte, []int) {
	return file_protos_tagger_tagger_proto_rawDescGZIP(), []int{0}
}

func (x *UnstructuredText) GetUnstructuredEntry() string {
	if x != nil {
		return x.UnstructuredEntry
	}
	return ""
}

// The response message containing the greetings
type TagReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *TagReply) Reset() {
	*x = TagReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_tagger_tagger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TagReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TagReply) ProtoMessage() {}

func (x *TagReply) ProtoReflect() protoreflect.Message {
	mi := &file_protos_tagger_tagger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TagReply.ProtoReflect.Descriptor instead.
func (*TagReply) Descriptor() ([]byte, []int) {
	return file_protos_tagger_tagger_proto_rawDescGZIP(), []int{1}
}

func (x *TagReply) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_protos_tagger_tagger_proto protoreflect.FileDescriptor

var file_protos_tagger_tagger_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2f,
	0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x10, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2d, 0x0a, 0x12, 0x75, 0x6e, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x1e, 0x0a, 0x08, 0x54, 0x61, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x32, 0x41, 0x0a, 0x06, 0x54, 0x61, 0x67, 0x67, 0x65,
	0x72, 0x12, 0x37, 0x0a, 0x07, 0x54, 0x61, 0x67, 0x54, 0x65, 0x78, 0x74, 0x12, 0x18, 0x2e, 0x74,
	0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x55, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x64, 0x54, 0x65, 0x78, 0x74, 0x1a, 0x10, 0x2e, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e,
	0x54, 0x61, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x22, 0x5a, 0x20, 0x70, 0x69,
	0x6e, 0x67, 0x70, 0x61, 0x77, 0x6e, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x74, 0x61, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_tagger_tagger_proto_rawDescOnce sync.Once
	file_protos_tagger_tagger_proto_rawDescData = file_protos_tagger_tagger_proto_rawDesc
)

func file_protos_tagger_tagger_proto_rawDescGZIP() []byte {
	file_protos_tagger_tagger_proto_rawDescOnce.Do(func() {
		file_protos_tagger_tagger_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_tagger_tagger_proto_rawDescData)
	})
	return file_protos_tagger_tagger_proto_rawDescData
}

var file_protos_tagger_tagger_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_tagger_tagger_proto_goTypes = []interface{}{
	(*UnstructuredText)(nil), // 0: tagger.UnstructuredText
	(*TagReply)(nil),         // 1: tagger.TagReply
}
var file_protos_tagger_tagger_proto_depIdxs = []int32{
	0, // 0: tagger.Tagger.TagText:input_type -> tagger.UnstructuredText
	1, // 1: tagger.Tagger.TagText:output_type -> tagger.TagReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_tagger_tagger_proto_init() }
func file_protos_tagger_tagger_proto_init() {
	if File_protos_tagger_tagger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_tagger_tagger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnstructuredText); i {
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
		file_protos_tagger_tagger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TagReply); i {
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
			RawDescriptor: file_protos_tagger_tagger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_tagger_tagger_proto_goTypes,
		DependencyIndexes: file_protos_tagger_tagger_proto_depIdxs,
		MessageInfos:      file_protos_tagger_tagger_proto_msgTypes,
	}.Build()
	File_protos_tagger_tagger_proto = out.File
	file_protos_tagger_tagger_proto_rawDesc = nil
	file_protos_tagger_tagger_proto_goTypes = nil
	file_protos_tagger_tagger_proto_depIdxs = nil
}
