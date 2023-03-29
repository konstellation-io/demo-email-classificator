// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: public_input.proto

package proto

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

type ClassificatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename  string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	BatchSize int32  `protobuf:"varint,2,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *ClassificatorRequest) Reset() {
	*x = ClassificatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificatorRequest) ProtoMessage() {}

func (x *ClassificatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificatorRequest.ProtoReflect.Descriptor instead.
func (*ClassificatorRequest) Descriptor() ([]byte, []int) {
	return file_public_input_proto_rawDescGZIP(), []int{0}
}

func (x *ClassificatorRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *ClassificatorRequest) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type BatchClassificatorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename  string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	BatchSize int32  `protobuf:"varint,2,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
}

func (x *BatchClassificatorRequest) Reset() {
	*x = BatchClassificatorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_input_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchClassificatorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchClassificatorRequest) ProtoMessage() {}

func (x *BatchClassificatorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_public_input_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchClassificatorRequest.ProtoReflect.Descriptor instead.
func (*BatchClassificatorRequest) Descriptor() ([]byte, []int) {
	return file_public_input_proto_rawDescGZIP(), []int{1}
}

func (x *BatchClassificatorRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *BatchClassificatorRequest) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_public_input_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_public_input_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_public_input_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_public_input_proto protoreflect.FileDescriptor

var file_public_input_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x51, 0x0a, 0x14, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x56, 0x0a,
	0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63,
	0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x94, 0x01, 0x0a, 0x0a,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x3d, 0x0a, 0x0d, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x1a, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x12, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x1f, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_public_input_proto_rawDescOnce sync.Once
	file_public_input_proto_rawDescData = file_public_input_proto_rawDesc
)

func file_public_input_proto_rawDescGZIP() []byte {
	file_public_input_proto_rawDescOnce.Do(func() {
		file_public_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_public_input_proto_rawDescData)
	})
	return file_public_input_proto_rawDescData
}

var file_public_input_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_public_input_proto_goTypes = []interface{}{
	(*ClassificatorRequest)(nil),      // 0: main.ClassificatorRequest
	(*BatchClassificatorRequest)(nil), // 1: main.BatchClassificatorRequest
	(*Response)(nil),                  // 2: main.Response
}
var file_public_input_proto_depIdxs = []int32{
	0, // 0: main.Entrypoint.Classificator:input_type -> main.ClassificatorRequest
	1, // 1: main.Entrypoint.BatchClassificator:input_type -> main.BatchClassificatorRequest
	2, // 2: main.Entrypoint.Classificator:output_type -> main.Response
	2, // 3: main.Entrypoint.BatchClassificator:output_type -> main.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_public_input_proto_init() }
func file_public_input_proto_init() {
	if File_public_input_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_public_input_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificatorRequest); i {
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
		file_public_input_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchClassificatorRequest); i {
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
		file_public_input_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_public_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_public_input_proto_goTypes,
		DependencyIndexes: file_public_input_proto_depIdxs,
		MessageInfos:      file_public_input_proto_msgTypes,
	}.Build()
	File_public_input_proto = out.File
	file_public_input_proto_rawDesc = nil
	file_public_input_proto_goTypes = nil
	file_public_input_proto_depIdxs = nil
}