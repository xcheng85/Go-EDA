// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: playerspb/api.proto

package playerspb

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

type GetPlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPlayerRequest) Reset() {
	*x = GetPlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playerspb_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerRequest) ProtoMessage() {}

func (x *GetPlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playerspb_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerRequest.ProtoReflect.Descriptor instead.
func (*GetPlayerRequest) Descriptor() ([]byte, []int) {
	return file_playerspb_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetPlayerRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *GetPlayerResponse) Reset() {
	*x = GetPlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playerspb_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerResponse) ProtoMessage() {}

func (x *GetPlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playerspb_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerResponse.ProtoReflect.Descriptor instead.
func (*GetPlayerResponse) Descriptor() ([]byte, []int) {
	return file_playerspb_api_proto_rawDescGZIP(), []int{1}
}

func (x *GetPlayerResponse) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

var File_playerspb_api_proto protoreflect.FileDescriptor

var file_playerspb_api_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2f, 0x61, 0x70, 0x69, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62,
	0x1a, 0x18, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3e,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x32, 0x5a,
	0x0a, 0x0e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x1b, 0x2e,
	0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8f, 0x01, 0x0a, 0x0d, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x42, 0x08, 0x41, 0x70,
	0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x63, 0x68, 0x65, 0x6e, 0x67, 0x38, 0x35, 0x2f, 0x47, 0x6f,
	0x2d, 0x45, 0x44, 0x41, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0xca, 0x02, 0x09, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0xe2, 0x02, 0x15, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x70, 0x62, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x09, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playerspb_api_proto_rawDescOnce sync.Once
	file_playerspb_api_proto_rawDescData = file_playerspb_api_proto_rawDesc
)

func file_playerspb_api_proto_rawDescGZIP() []byte {
	file_playerspb_api_proto_rawDescOnce.Do(func() {
		file_playerspb_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_playerspb_api_proto_rawDescData)
	})
	return file_playerspb_api_proto_rawDescData
}

var file_playerspb_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_playerspb_api_proto_goTypes = []interface{}{
	(*GetPlayerRequest)(nil),  // 0: playerspb.GetPlayerRequest
	(*GetPlayerResponse)(nil), // 1: playerspb.GetPlayerResponse
	(*Player)(nil),            // 2: playerspb.Player
}
var file_playerspb_api_proto_depIdxs = []int32{
	2, // 0: playerspb.GetPlayerResponse.player:type_name -> playerspb.Player
	0, // 1: playerspb.PlayersService.GetPlayer:input_type -> playerspb.GetPlayerRequest
	1, // 2: playerspb.PlayersService.GetPlayer:output_type -> playerspb.GetPlayerResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_playerspb_api_proto_init() }
func file_playerspb_api_proto_init() {
	if File_playerspb_api_proto != nil {
		return
	}
	file_playerspb_messages_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_playerspb_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerRequest); i {
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
		file_playerspb_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPlayerResponse); i {
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
			RawDescriptor: file_playerspb_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playerspb_api_proto_goTypes,
		DependencyIndexes: file_playerspb_api_proto_depIdxs,
		MessageInfos:      file_playerspb_api_proto_msgTypes,
	}.Build()
	File_playerspb_api_proto = out.File
	file_playerspb_api_proto_rawDesc = nil
	file_playerspb_api_proto_goTypes = nil
	file_playerspb_api_proto_depIdxs = nil
}
