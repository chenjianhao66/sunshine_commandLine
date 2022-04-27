// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: internal/proto/BlogServer.proto

package grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// 请求参数
type BlogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Body []byte `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *BlogReq) Reset() {
	*x = BlogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_BlogServer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogReq) ProtoMessage() {}

func (x *BlogReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_BlogServer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogReq.ProtoReflect.Descriptor instead.
func (*BlogReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_BlogServer_proto_rawDescGZIP(), []int{0}
}

func (x *BlogReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BlogReq) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

// 返回参数
type BlogResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *BlogResp) Reset() {
	*x = BlogResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_BlogServer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlogResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlogResp) ProtoMessage() {}

func (x *BlogResp) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_BlogServer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlogResp.ProtoReflect.Descriptor instead.
func (*BlogResp) Descriptor() ([]byte, []int) {
	return file_internal_proto_BlogServer_proto_rawDescGZIP(), []int{1}
}

func (x *BlogResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BlogResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_proto_BlogServer_proto protoreflect.FileDescriptor

var file_internal_proto_BlogServer_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x42, 0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x31, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x67, 0x52,
	0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x22, 0x38, 0x0a, 0x08, 0x42, 0x6c,
	0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0x37, 0x0a, 0x04, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x2f, 0x0a, 0x0a,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x6c, 0x6f, 0x67, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x42, 0x6c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x28, 0x01, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_BlogServer_proto_rawDescOnce sync.Once
	file_internal_proto_BlogServer_proto_rawDescData = file_internal_proto_BlogServer_proto_rawDesc
)

func file_internal_proto_BlogServer_proto_rawDescGZIP() []byte {
	file_internal_proto_BlogServer_proto_rawDescOnce.Do(func() {
		file_internal_proto_BlogServer_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_BlogServer_proto_rawDescData)
	})
	return file_internal_proto_BlogServer_proto_rawDescData
}

var file_internal_proto_BlogServer_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_proto_BlogServer_proto_goTypes = []interface{}{
	(*BlogReq)(nil),  // 0: grpc.BlogReq
	(*BlogResp)(nil), // 1: grpc.BlogResp
}
var file_internal_proto_BlogServer_proto_depIdxs = []int32{
	0, // 0: grpc.Blog.uploadBlog:input_type -> grpc.BlogReq
	1, // 1: grpc.Blog.uploadBlog:output_type -> grpc.BlogResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_BlogServer_proto_init() }
func file_internal_proto_BlogServer_proto_init() {
	if File_internal_proto_BlogServer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_BlogServer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogReq); i {
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
		file_internal_proto_BlogServer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlogResp); i {
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
			RawDescriptor: file_internal_proto_BlogServer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_BlogServer_proto_goTypes,
		DependencyIndexes: file_internal_proto_BlogServer_proto_depIdxs,
		MessageInfos:      file_internal_proto_BlogServer_proto_msgTypes,
	}.Build()
	File_internal_proto_BlogServer_proto = out.File
	file_internal_proto_BlogServer_proto_rawDesc = nil
	file_internal_proto_BlogServer_proto_goTypes = nil
	file_internal_proto_BlogServer_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlogClient is the client API for Blog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogClient interface {
	UploadBlog(ctx context.Context, opts ...grpc.CallOption) (Blog_UploadBlogClient, error)
}

type blogClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogClient(cc grpc.ClientConnInterface) BlogClient {
	return &blogClient{cc}
}

func (c *blogClient) UploadBlog(ctx context.Context, opts ...grpc.CallOption) (Blog_UploadBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Blog_serviceDesc.Streams[0], "/grpc.Blog/uploadBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogUploadBlogClient{stream}
	return x, nil
}

type Blog_UploadBlogClient interface {
	Send(*BlogReq) error
	CloseAndRecv() (*BlogResp, error)
	grpc.ClientStream
}

type blogUploadBlogClient struct {
	grpc.ClientStream
}

func (x *blogUploadBlogClient) Send(m *BlogReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blogUploadBlogClient) CloseAndRecv() (*BlogResp, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BlogResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServer is the server API for Blog service.
type BlogServer interface {
	UploadBlog(Blog_UploadBlogServer) error
}

// UnimplementedBlogServer can be embedded to have forward compatible implementations.
type UnimplementedBlogServer struct {
}

func (*UnimplementedBlogServer) UploadBlog(Blog_UploadBlogServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadBlog not implemented")
}

func RegisterBlogServer(s *grpc.Server, srv BlogServer) {
	s.RegisterService(&_Blog_serviceDesc, srv)
}

func _Blog_UploadBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlogServer).UploadBlog(&blogUploadBlogServer{stream})
}

type Blog_UploadBlogServer interface {
	SendAndClose(*BlogResp) error
	Recv() (*BlogReq, error)
	grpc.ServerStream
}

type blogUploadBlogServer struct {
	grpc.ServerStream
}

func (x *blogUploadBlogServer) SendAndClose(m *BlogResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blogUploadBlogServer) Recv() (*BlogReq, error) {
	m := new(BlogReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Blog_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Blog",
	HandlerType: (*BlogServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "uploadBlog",
			Handler:       _Blog_UploadBlog_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "internal/proto/BlogServer.proto",
}