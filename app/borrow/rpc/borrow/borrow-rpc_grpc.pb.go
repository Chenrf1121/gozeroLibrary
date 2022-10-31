// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: borrow-rpc.proto

package borrow

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SearchAllClient is the client API for SearchAll service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SearchAllClient interface {
	SearchAll(ctx context.Context, in *BorrwoIdReq, opts ...grpc.CallOption) (*BorrowListResp, error)
}

type searchAllClient struct {
	cc grpc.ClientConnInterface
}

func NewSearchAllClient(cc grpc.ClientConnInterface) SearchAllClient {
	return &searchAllClient{cc}
}

func (c *searchAllClient) SearchAll(ctx context.Context, in *BorrwoIdReq, opts ...grpc.CallOption) (*BorrowListResp, error) {
	out := new(BorrowListResp)
	err := c.cc.Invoke(ctx, "/borrow.SearchAll/searchAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchAllServer is the server API for SearchAll service.
// All implementations must embed UnimplementedSearchAllServer
// for forward compatibility
type SearchAllServer interface {
	SearchAll(context.Context, *BorrwoIdReq) (*BorrowListResp, error)
	mustEmbedUnimplementedSearchAllServer()
}

// UnimplementedSearchAllServer must be embedded to have forward compatible implementations.
type UnimplementedSearchAllServer struct {
}

func (UnimplementedSearchAllServer) SearchAll(context.Context, *BorrwoIdReq) (*BorrowListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchAll not implemented")
}
func (UnimplementedSearchAllServer) mustEmbedUnimplementedSearchAllServer() {}

// UnsafeSearchAllServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SearchAllServer will
// result in compilation errors.
type UnsafeSearchAllServer interface {
	mustEmbedUnimplementedSearchAllServer()
}

func RegisterSearchAllServer(s grpc.ServiceRegistrar, srv SearchAllServer) {
	s.RegisterService(&SearchAll_ServiceDesc, srv)
}

func _SearchAll_SearchAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BorrwoIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchAllServer).SearchAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/borrow.SearchAll/searchAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchAllServer).SearchAll(ctx, req.(*BorrwoIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SearchAll_ServiceDesc is the grpc.ServiceDesc for SearchAll service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SearchAll_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "borrow.SearchAll",
	HandlerType: (*SearchAllServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "searchAll",
			Handler:    _SearchAll_SearchAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "borrow-rpc.proto",
}