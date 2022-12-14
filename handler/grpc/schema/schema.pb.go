// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: handler/grpc/schema/schema.proto

package schema

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

type LocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *LocationRequest) Reset() {
	*x = LocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationRequest) ProtoMessage() {}

func (x *LocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationRequest.ProtoReflect.Descriptor instead.
func (*LocationRequest) Descriptor() ([]byte, []int) {
	return file_handler_grpc_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *LocationRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type LocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip          string  `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	CountryCode string  `protobuf:"bytes,2,opt,name=countryCode,proto3" json:"countryCode,omitempty"`
	Country     string  `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	City        string  `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty"`
	Latitude    float64 `protobuf:"fixed64,5,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude   float64 `protobuf:"fixed64,6,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *LocationResponse) Reset() {
	*x = LocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_handler_grpc_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocationResponse) ProtoMessage() {}

func (x *LocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_handler_grpc_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocationResponse.ProtoReflect.Descriptor instead.
func (*LocationResponse) Descriptor() ([]byte, []int) {
	return file_handler_grpc_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *LocationResponse) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LocationResponse) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *LocationResponse) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *LocationResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *LocationResponse) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *LocationResponse) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

var File_handler_grpc_schema_schema_proto protoreflect.FileDescriptor

var file_handler_grpc_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x20, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22,
	0x21, 0x0a, 0x0f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x70, 0x22, 0xac, 0x01, 0x0a, 0x10, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74,
	0x75, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64,
	0x65, 0x32, 0x5f, 0x0a, 0x0b, 0x47, 0x65, 0x6f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x50, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x69, 0x61, 0x67, 0x6f, 0x63, 0x65, 0x73, 0x61, 0x72, 0x2f, 0x67, 0x65, 0x6f, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_handler_grpc_schema_schema_proto_rawDescOnce sync.Once
	file_handler_grpc_schema_schema_proto_rawDescData = file_handler_grpc_schema_schema_proto_rawDesc
)

func file_handler_grpc_schema_schema_proto_rawDescGZIP() []byte {
	file_handler_grpc_schema_schema_proto_rawDescOnce.Do(func() {
		file_handler_grpc_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_handler_grpc_schema_schema_proto_rawDescData)
	})
	return file_handler_grpc_schema_schema_proto_rawDescData
}

var file_handler_grpc_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_handler_grpc_schema_schema_proto_goTypes = []interface{}{
	(*LocationRequest)(nil),  // 0: grpc_server.LocationRequest
	(*LocationResponse)(nil), // 1: grpc_server.LocationResponse
}
var file_handler_grpc_schema_schema_proto_depIdxs = []int32{
	0, // 0: grpc_server.Geolocation.GetLocationData:input_type -> grpc_server.LocationRequest
	1, // 1: grpc_server.Geolocation.GetLocationData:output_type -> grpc_server.LocationResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_handler_grpc_schema_schema_proto_init() }
func file_handler_grpc_schema_schema_proto_init() {
	if File_handler_grpc_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_handler_grpc_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationRequest); i {
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
		file_handler_grpc_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LocationResponse); i {
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
			RawDescriptor: file_handler_grpc_schema_schema_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_handler_grpc_schema_schema_proto_goTypes,
		DependencyIndexes: file_handler_grpc_schema_schema_proto_depIdxs,
		MessageInfos:      file_handler_grpc_schema_schema_proto_msgTypes,
	}.Build()
	File_handler_grpc_schema_schema_proto = out.File
	file_handler_grpc_schema_schema_proto_rawDesc = nil
	file_handler_grpc_schema_schema_proto_goTypes = nil
	file_handler_grpc_schema_schema_proto_depIdxs = nil
}
