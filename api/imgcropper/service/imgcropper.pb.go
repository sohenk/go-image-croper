// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.2
// source: imgcropper.proto

package service

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CropImgRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url   string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Width int64  `protobuf:"varint,2,opt,name=width,proto3" json:"width,omitempty"`
}

func (x *CropImgRequest) Reset() {
	*x = CropImgRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imgcropper_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CropImgRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CropImgRequest) ProtoMessage() {}

func (x *CropImgRequest) ProtoReflect() protoreflect.Message {
	mi := &file_imgcropper_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CropImgRequest.ProtoReflect.Descriptor instead.
func (*CropImgRequest) Descriptor() ([]byte, []int) {
	return file_imgcropper_proto_rawDescGZIP(), []int{0}
}

func (x *CropImgRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CropImgRequest) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

type CropImgReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imgname   string `protobuf:"bytes,1,opt,name=imgname,proto3" json:"imgname,omitempty"`
	Imagetype string `protobuf:"bytes,2,opt,name=imagetype,proto3" json:"imagetype,omitempty"`
	Imgdata   []byte `protobuf:"bytes,3,opt,name=imgdata,proto3" json:"imgdata,omitempty"`
}

func (x *CropImgReply) Reset() {
	*x = CropImgReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_imgcropper_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CropImgReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CropImgReply) ProtoMessage() {}

func (x *CropImgReply) ProtoReflect() protoreflect.Message {
	mi := &file_imgcropper_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CropImgReply.ProtoReflect.Descriptor instead.
func (*CropImgReply) Descriptor() ([]byte, []int) {
	return file_imgcropper_proto_rawDescGZIP(), []int{1}
}

func (x *CropImgReply) GetImgname() string {
	if x != nil {
		return x.Imgname
	}
	return ""
}

func (x *CropImgReply) GetImagetype() string {
	if x != nil {
		return x.Imagetype
	}
	return ""
}

func (x *CropImgReply) GetImgdata() []byte {
	if x != nil {
		return x.Imgdata
	}
	return nil
}

var File_imgcropper_proto protoreflect.FileDescriptor

var file_imgcropper_proto_rawDesc = []byte{
	0x0a, 0x10, 0x69, 0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70,
	0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x4a, 0x0a, 0x0e, 0x43, 0x72, 0x6f, 0x70, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x1d,
	0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa,
	0x42, 0x04, 0x22, 0x02, 0x28, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x60, 0x0a,
	0x0c, 0x43, 0x72, 0x6f, 0x70, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x69, 0x6d, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x69, 0x6d, 0x67, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x6d, 0x67, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x69, 0x6d, 0x67, 0x64, 0x61, 0x74, 0x61, 0x32,
	0x7a, 0x0a, 0x0a, 0x49, 0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x12, 0x6c, 0x0a,
	0x07, 0x43, 0x72, 0x6f, 0x70, 0x49, 0x6d, 0x67, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69,
	0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x72, 0x6f, 0x70, 0x49, 0x6d, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65,
	0x72, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x6f, 0x70, 0x49, 0x6d,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b,
	0x2f, 0x67, 0x65, 0x74, 0x63, 0x72, 0x6f, 0x70, 0x69, 0x6d, 0x67, 0x42, 0x42, 0x0a, 0x16, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x01, 0x5a, 0x26, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x67, 0x63, 0x72, 0x6f, 0x70, 0x70, 0x65, 0x72, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_imgcropper_proto_rawDescOnce sync.Once
	file_imgcropper_proto_rawDescData = file_imgcropper_proto_rawDesc
)

func file_imgcropper_proto_rawDescGZIP() []byte {
	file_imgcropper_proto_rawDescOnce.Do(func() {
		file_imgcropper_proto_rawDescData = protoimpl.X.CompressGZIP(file_imgcropper_proto_rawDescData)
	})
	return file_imgcropper_proto_rawDescData
}

var file_imgcropper_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_imgcropper_proto_goTypes = []interface{}{
	(*CropImgRequest)(nil), // 0: api.imgcropper.service.CropImgRequest
	(*CropImgReply)(nil),   // 1: api.imgcropper.service.CropImgReply
}
var file_imgcropper_proto_depIdxs = []int32{
	0, // 0: api.imgcropper.service.Imgcropper.CropImg:input_type -> api.imgcropper.service.CropImgRequest
	1, // 1: api.imgcropper.service.Imgcropper.CropImg:output_type -> api.imgcropper.service.CropImgReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_imgcropper_proto_init() }
func file_imgcropper_proto_init() {
	if File_imgcropper_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_imgcropper_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CropImgRequest); i {
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
		file_imgcropper_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CropImgReply); i {
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
			RawDescriptor: file_imgcropper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_imgcropper_proto_goTypes,
		DependencyIndexes: file_imgcropper_proto_depIdxs,
		MessageInfos:      file_imgcropper_proto_msgTypes,
	}.Build()
	File_imgcropper_proto = out.File
	file_imgcropper_proto_rawDesc = nil
	file_imgcropper_proto_goTypes = nil
	file_imgcropper_proto_depIdxs = nil
}
