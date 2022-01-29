// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: minekube/connect/v1alpha1/tunnel_service.proto

package connectv1alpha1

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

type TunnelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is the id of the session to create the tunnel for.
	// The session id is sent by the WatchService.
	// The first TunnelRequest stream message must specify this field initially
	// and proceeding messages must only provide the data field.
	SessionId string `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	// This is a raw client bound chunk of data.
	//
	// The data is just a chunk of the stream and never
	// assumed to be a complete single packet.
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TunnelRequest) Reset() {
	*x = TunnelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelRequest) ProtoMessage() {}

func (x *TunnelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelRequest.ProtoReflect.Descriptor instead.
func (*TunnelRequest) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescGZIP(), []int{0}
}

func (x *TunnelRequest) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *TunnelRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type TunnelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// This is a raw server bound chunk of data.
	//
	// The data is just a chunk of the stream and never
	// assumed to be a complete single packet.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TunnelResponse) Reset() {
	*x = TunnelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TunnelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TunnelResponse) ProtoMessage() {}

func (x *TunnelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TunnelResponse.ProtoReflect.Descriptor instead.
func (*TunnelResponse) Descriptor() ([]byte, []int) {
	return file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescGZIP(), []int{1}
}

func (x *TunnelResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_minekube_connect_v1alpha1_tunnel_service_proto protoreflect.FileDescriptor

var file_minekube_connect_v1alpha1_tunnel_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x6e,
	0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x42, 0x0a, 0x0d, 0x54,
	0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x24, 0x0a, 0x0e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0x72, 0x0a, 0x0d, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x06, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x28, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6d, 0x69, 0x6e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0xfe, 0x01, 0x0a, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x12, 0x54, 0x75, 0x6e,
	0x6e, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a,
	0x45, 0x67, 0x6f, 0x2e, 0x6d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x69, 0x6e,
	0x65, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x4d, 0x43, 0x58, 0xaa, 0x02, 0x19, 0x4d,
	0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x19, 0x4d, 0x69, 0x6e, 0x65, 0x6b,
	0x75, 0x62, 0x65, 0x5c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x25, 0x4d, 0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x5c,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1b, 0x4d,
	0x69, 0x6e, 0x65, 0x6b, 0x75, 0x62, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescOnce sync.Once
	file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescData = file_minekube_connect_v1alpha1_tunnel_service_proto_rawDesc
)

func file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescGZIP() []byte {
	file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescOnce.Do(func() {
		file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescData)
	})
	return file_minekube_connect_v1alpha1_tunnel_service_proto_rawDescData
}

var file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_minekube_connect_v1alpha1_tunnel_service_proto_goTypes = []interface{}{
	(*TunnelRequest)(nil),  // 0: minekube.connect.v1alpha1.TunnelRequest
	(*TunnelResponse)(nil), // 1: minekube.connect.v1alpha1.TunnelResponse
}
var file_minekube_connect_v1alpha1_tunnel_service_proto_depIdxs = []int32{
	0, // 0: minekube.connect.v1alpha1.TunnelService.Tunnel:input_type -> minekube.connect.v1alpha1.TunnelRequest
	1, // 1: minekube.connect.v1alpha1.TunnelService.Tunnel:output_type -> minekube.connect.v1alpha1.TunnelResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_minekube_connect_v1alpha1_tunnel_service_proto_init() }
func file_minekube_connect_v1alpha1_tunnel_service_proto_init() {
	if File_minekube_connect_v1alpha1_tunnel_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelRequest); i {
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
		file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TunnelResponse); i {
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
			RawDescriptor: file_minekube_connect_v1alpha1_tunnel_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_minekube_connect_v1alpha1_tunnel_service_proto_goTypes,
		DependencyIndexes: file_minekube_connect_v1alpha1_tunnel_service_proto_depIdxs,
		MessageInfos:      file_minekube_connect_v1alpha1_tunnel_service_proto_msgTypes,
	}.Build()
	File_minekube_connect_v1alpha1_tunnel_service_proto = out.File
	file_minekube_connect_v1alpha1_tunnel_service_proto_rawDesc = nil
	file_minekube_connect_v1alpha1_tunnel_service_proto_goTypes = nil
	file_minekube_connect_v1alpha1_tunnel_service_proto_depIdxs = nil
}
