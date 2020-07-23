// Code generated by protoc-gen-go. DO NOT EDIT.
// source: copy.proto

package copy

import (
	context "context"
	fmt "fmt"
	_ "git.zapa.cloud/merchant-tools/protobuf/middleware/report"
	_ "git.zapa.cloud/merchant-tools/protobuf/report-service"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	core_service "github.com/tvducmt/protoc-gen-copy/protobuf-file/core-service"
	middleware "github.com/tvducmt/protoc-gen-copy/protobuf-file/middleware"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("copy.proto", fileDescriptor_8b8564a7da00bcc7) }

var fileDescriptor_8b8564a7da00bcc7 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xbf, 0x4e, 0x04, 0x21,
	0x10, 0xc6, 0x2d, 0x2e, 0x16, 0x5b, 0x58, 0x6c, 0x61, 0xcc, 0xda, 0xd9, 0xc3, 0x26, 0xfa, 0x08,
	0xfe, 0x89, 0x31, 0x17, 0x5f, 0xc1, 0xb0, 0xc3, 0x1c, 0x4b, 0x02, 0x37, 0x04, 0x86, 0x33, 0xe7,
	0xd3, 0x1b, 0x0e, 0xbc, 0x5d, 0x3b, 0xad, 0x66, 0x3e, 0x86, 0xf9, 0x4d, 0xbe, 0xaf, 0xeb, 0x80,
	0xc2, 0x51, 0x86, 0x48, 0x4c, 0xfd, 0xa6, 0xf4, 0xc3, 0xab, 0xb1, 0x3c, 0xe7, 0x49, 0x02, 0xf9,
	0x91, 0x0f, 0x3a, 0x83, 0xe7, 0xf1, 0x34, 0x07, 0x61, 0x70, 0x2f, 0xca, 0x9f, 0xaa, 0xa7, 0xbc,
	0x13, 0x3b, 0xeb, 0x70, 0x04, 0x8a, 0x28, 0x12, 0xc6, 0x83, 0x85, 0x2a, 0x2a, 0x6f, 0xd8, 0xfe,
	0x9b, 0xe4, 0xad, 0xd6, 0x0e, 0x3f, 0x55, 0x5c, 0xb7, 0x8d, 0xf6, 0x6c, 0x2c, 0xcb, 0x2f, 0x15,
	0x94, 0x04, 0x47, 0x59, 0x8f, 0x1e, 0x23, 0xcc, 0x6a, 0xcf, 0x82, 0x89, 0x5c, 0x3a, 0x83, 0xd6,
	0x8c, 0x88, 0x81, 0x22, 0xb7, 0xd2, 0x30, 0x6f, 0x7f, 0xc4, 0xd4, 0xa5, 0xb3, 0xad, 0xdf, 0xb2,
	0xb1, 0x6e, 0x0d, 0x91, 0x71, 0xb8, 0x2c, 0xa1, 0x0f, 0xdc, 0xd2, 0xbc, 0xdf, 0x76, 0x9b, 0x47,
	0x0a, 0xc7, 0xfe, 0xa9, 0xbb, 0x2a, 0xf5, 0x1d, 0x13, 0xa3, 0x7e, 0xb1, 0x69, 0xee, 0xaf, 0xe5,
	0xca, 0xdc, 0xf2, 0x3e, 0xdc, 0xc8, 0x12, 0xde, 0xc7, 0xcf, 0x8d, 0x65, 0x72, 0x77, 0x31, 0x5d,
	0x9e, 0xa0, 0x0f, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2c, 0xc1, 0x8b, 0xdb, 0xb0, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CopyClient is the client API for Copy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CopyClient interface {
	// rpc ListCITransactionResponse (report.service.ListTransactionResponse) returns (api.report.ListTransLogResponse){};
	// rpc CopyFish(middleware.Fish) returns (core_service.Fish) {}
	CopyNestedFish(ctx context.Context, in *middleware.NestedFish, opts ...grpc.CallOption) (*core_service.NestedFish, error)
}

type copyClient struct {
	cc *grpc.ClientConn
}

func NewCopyClient(cc *grpc.ClientConn) CopyClient {
	return &copyClient{cc}
}

func (c *copyClient) CopyNestedFish(ctx context.Context, in *middleware.NestedFish, opts ...grpc.CallOption) (*core_service.NestedFish, error) {
	out := new(core_service.NestedFish)
	err := c.cc.Invoke(ctx, "/copy.Copy/CopyNestedFish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CopyServer is the server API for Copy service.
type CopyServer interface {
	// rpc ListCITransactionResponse (report.service.ListTransactionResponse) returns (api.report.ListTransLogResponse){};
	// rpc CopyFish(middleware.Fish) returns (core_service.Fish) {}
	CopyNestedFish(context.Context, *middleware.NestedFish) (*core_service.NestedFish, error)
}

// UnimplementedCopyServer can be embedded to have forward compatible implementations.
type UnimplementedCopyServer struct {
}

func (*UnimplementedCopyServer) CopyNestedFish(ctx context.Context, req *middleware.NestedFish) (*core_service.NestedFish, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CopyNestedFish not implemented")
}

func RegisterCopyServer(s *grpc.Server, srv CopyServer) {
	s.RegisterService(&_Copy_serviceDesc, srv)
}

func _Copy_CopyNestedFish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(middleware.NestedFish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CopyServer).CopyNestedFish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/copy.Copy/CopyNestedFish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CopyServer).CopyNestedFish(ctx, req.(*middleware.NestedFish))
	}
	return interceptor(ctx, in, info, handler)
}

var _Copy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "copy.Copy",
	HandlerType: (*CopyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CopyNestedFish",
			Handler:    _Copy_CopyNestedFish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "copy.proto",
}
