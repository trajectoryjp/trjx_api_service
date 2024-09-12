// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: uas/generic/v2/sdsp_g.proto

// プロダクト名：Generic API
// パッケージ名：generic.v2
// インタフェース名：generic.v2.GenericAPI

package spatial

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GenericAPI_ConnectServer_FullMethodName = "/generic.v2.GenericAPI/ConnectServer"
)

// GenericAPIClient is the client API for GenericAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GenericAPIClient interface {
	// サーバー接続（ログイン）
	// ユーザIDとパスワードでトークンを得る
	ConnectServer(ctx context.Context, in *ConnectServerRequest, opts ...grpc.CallOption) (*Token, error)
}

type genericAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewGenericAPIClient(cc grpc.ClientConnInterface) GenericAPIClient {
	return &genericAPIClient{cc}
}

func (c *genericAPIClient) ConnectServer(ctx context.Context, in *ConnectServerRequest, opts ...grpc.CallOption) (*Token, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Token)
	err := c.cc.Invoke(ctx, GenericAPI_ConnectServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenericAPIServer is the server API for GenericAPI service.
// All implementations must embed UnimplementedGenericAPIServer
// for forward compatibility
type GenericAPIServer interface {
	// サーバー接続（ログイン）
	// ユーザIDとパスワードでトークンを得る
	ConnectServer(context.Context, *ConnectServerRequest) (*Token, error)
	mustEmbedUnimplementedGenericAPIServer()
}

// UnimplementedGenericAPIServer must be embedded to have forward compatible implementations.
type UnimplementedGenericAPIServer struct {
}

func (UnimplementedGenericAPIServer) ConnectServer(context.Context, *ConnectServerRequest) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectServer not implemented")
}
func (UnimplementedGenericAPIServer) mustEmbedUnimplementedGenericAPIServer() {}

// UnsafeGenericAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GenericAPIServer will
// result in compilation errors.
type UnsafeGenericAPIServer interface {
	mustEmbedUnimplementedGenericAPIServer()
}

func RegisterGenericAPIServer(s grpc.ServiceRegistrar, srv GenericAPIServer) {
	s.RegisterService(&GenericAPI_ServiceDesc, srv)
}

func _GenericAPI_ConnectServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenericAPIServer).ConnectServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GenericAPI_ConnectServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenericAPIServer).ConnectServer(ctx, req.(*ConnectServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GenericAPI_ServiceDesc is the grpc.ServiceDesc for GenericAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GenericAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generic.v2.GenericAPI",
	HandlerType: (*GenericAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ConnectServer",
			Handler:    _GenericAPI_ConnectServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "uas/generic/v2/sdsp_g.proto",
}
