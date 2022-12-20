// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.11
// source: netcity.proto

package netcity_pb

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

type AuthParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cid   int32  `protobuf:"varint,1,opt,name=Cid,proto3" json:"Cid,omitempty"`
	Sid   int32  `protobuf:"varint,2,opt,name=Sid,proto3" json:"Sid,omitempty"`
	Pid   int32  `protobuf:"varint,3,opt,name=Pid,proto3" json:"Pid,omitempty"`
	Cn    int32  `protobuf:"varint,4,opt,name=Cn,proto3" json:"Cn,omitempty"`
	Scid  int32  `protobuf:"varint,5,opt,name=Scid,proto3" json:"Scid,omitempty"`
	UN    string `protobuf:"bytes,6,opt,name=UN,proto3" json:"UN,omitempty"`
	PW    string `protobuf:"bytes,7,opt,name=PW,proto3" json:"PW,omitempty"`
	Sft   int32  `protobuf:"varint,8,opt,name=Sft,proto3" json:"Sft,omitempty"`
	UrlId uint32 `protobuf:"varint,9,opt,name=UrlId,proto3" json:"UrlId,omitempty"`
}

func (x *AuthParam) Reset() {
	*x = AuthParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netcity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthParam) ProtoMessage() {}

func (x *AuthParam) ProtoReflect() protoreflect.Message {
	mi := &file_netcity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthParam.ProtoReflect.Descriptor instead.
func (*AuthParam) Descriptor() ([]byte, []int) {
	return file_netcity_proto_rawDescGZIP(), []int{0}
}

func (x *AuthParam) GetCid() int32 {
	if x != nil {
		return x.Cid
	}
	return 0
}

func (x *AuthParam) GetSid() int32 {
	if x != nil {
		return x.Sid
	}
	return 0
}

func (x *AuthParam) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *AuthParam) GetCn() int32 {
	if x != nil {
		return x.Cn
	}
	return 0
}

func (x *AuthParam) GetScid() int32 {
	if x != nil {
		return x.Scid
	}
	return 0
}

func (x *AuthParam) GetUN() string {
	if x != nil {
		return x.UN
	}
	return ""
}

func (x *AuthParam) GetPW() string {
	if x != nil {
		return x.PW
	}
	return ""
}

func (x *AuthParam) GetSft() int32 {
	if x != nil {
		return x.Sft
	}
	return 0
}

func (x *AuthParam) GetUrlId() uint32 {
	if x != nil {
		return x.UrlId
	}
	return 0
}

type UserConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSubscribeAssignments bool `protobuf:"varint,1,opt,name=is_subscribe_assignments,json=isSubscribeAssignments,proto3" json:"is_subscribe_assignments,omitempty"`
	IsSubscribeMarks       bool `protobuf:"varint,2,opt,name=is_subscribe_marks,json=isSubscribeMarks,proto3" json:"is_subscribe_marks,omitempty"`
}

func (x *UserConfig) Reset() {
	*x = UserConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_netcity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserConfig) ProtoMessage() {}

func (x *UserConfig) ProtoReflect() protoreflect.Message {
	mi := &file_netcity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserConfig.ProtoReflect.Descriptor instead.
func (*UserConfig) Descriptor() ([]byte, []int) {
	return file_netcity_proto_rawDescGZIP(), []int{1}
}

func (x *UserConfig) GetIsSubscribeAssignments() bool {
	if x != nil {
		return x.IsSubscribeAssignments
	}
	return false
}

func (x *UserConfig) GetIsSubscribeMarks() bool {
	if x != nil {
		return x.IsSubscribeMarks
	}
	return false
}

var File_netcity_proto protoreflect.FileDescriptor

var file_netcity_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6e, 0x65, 0x74, 0x63, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6e, 0x65, 0x74, 0x63, 0x69, 0x74, 0x79, 0x22, 0xad, 0x01, 0x0a, 0x09, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x43, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x50, 0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x43, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x43, 0x6e, 0x12, 0x12, 0x0a, 0x04,
	0x53, 0x63, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x53, 0x63, 0x69, 0x64,
	0x12, 0x0e, 0x0a, 0x02, 0x55, 0x4e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x55, 0x4e,
	0x12, 0x0e, 0x0a, 0x02, 0x50, 0x57, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x50, 0x57,
	0x12, 0x10, 0x0a, 0x03, 0x53, 0x66, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x53,
	0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x55, 0x72, 0x6c, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x05, 0x55, 0x72, 0x6c, 0x49, 0x64, 0x22, 0x74, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x38, 0x0a, 0x18, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x69, 0x73, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x2c, 0x0a, 0x12, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x73,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4d, 0x61, 0x72, 0x6b, 0x73, 0x42, 0x2f,
	0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6d, 0x6c,
	0x65, 0x62, 0x65, 0x64, 0x65, 0x76, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x69, 0x74, 0x79, 0x62, 0x6f,
	0x74, 0x2f, 0x70, 0x62, 0x2f, 0x6e, 0x65, 0x74, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_netcity_proto_rawDescOnce sync.Once
	file_netcity_proto_rawDescData = file_netcity_proto_rawDesc
)

func file_netcity_proto_rawDescGZIP() []byte {
	file_netcity_proto_rawDescOnce.Do(func() {
		file_netcity_proto_rawDescData = protoimpl.X.CompressGZIP(file_netcity_proto_rawDescData)
	})
	return file_netcity_proto_rawDescData
}

var file_netcity_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_netcity_proto_goTypes = []interface{}{
	(*AuthParam)(nil),  // 0: netcity.AuthParam
	(*UserConfig)(nil), // 1: netcity.UserConfig
}
var file_netcity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_netcity_proto_init() }
func file_netcity_proto_init() {
	if File_netcity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_netcity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthParam); i {
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
		file_netcity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserConfig); i {
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
			RawDescriptor: file_netcity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_netcity_proto_goTypes,
		DependencyIndexes: file_netcity_proto_depIdxs,
		MessageInfos:      file_netcity_proto_msgTypes,
	}.Build()
	File_netcity_proto = out.File
	file_netcity_proto_rawDesc = nil
	file_netcity_proto_goTypes = nil
	file_netcity_proto_depIdxs = nil
}
