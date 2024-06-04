// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.12.2
// source: other/types.proto

package other

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

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val *int64 `protobuf:"varint,1,opt,name=val" json:"val,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_other_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_other_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_other_types_proto_rawDescGZIP(), []int{0}
}

func (x *Test) GetVal() int64 {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return 0
}

var File_other_types_proto protoreflect.FileDescriptor

var file_other_types_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x22, 0x18, 0x0a, 0x04, 0x54, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x66, 0x68, 0x65, 0x72, 0x6d, 0x65, 0x6e, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2d, 0x63, 0x61, 0x63, 0x68, 0x65, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72,
}

var (
	file_other_types_proto_rawDescOnce sync.Once
	file_other_types_proto_rawDescData = file_other_types_proto_rawDesc
)

func file_other_types_proto_rawDescGZIP() []byte {
	file_other_types_proto_rawDescOnce.Do(func() {
		file_other_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_other_types_proto_rawDescData)
	})
	return file_other_types_proto_rawDescData
}

var file_other_types_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_other_types_proto_goTypes = []interface{}{
	(*Test)(nil), // 0: other.Test
}
var file_other_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_other_types_proto_init() }
func file_other_types_proto_init() {
	if File_other_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_other_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
			RawDescriptor: file_other_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_other_types_proto_goTypes,
		DependencyIndexes: file_other_types_proto_depIdxs,
		MessageInfos:      file_other_types_proto_msgTypes,
	}.Build()
	File_other_types_proto = out.File
	file_other_types_proto_rawDesc = nil
	file_other_types_proto_goTypes = nil
	file_other_types_proto_depIdxs = nil
}
