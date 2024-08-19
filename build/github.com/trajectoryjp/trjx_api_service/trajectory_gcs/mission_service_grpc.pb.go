// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: trajectory/trajectory_gcs_service/protocol/v1/mission_service.proto

package trajectory_gcs

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MissionService_CreateMission_FullMethodName        = "/trajectory.trajectory_gcs_service.protocol.v1.MissionService/CreateMission"
	MissionService_CreateMissionRecord_FullMethodName  = "/trajectory.trajectory_gcs_service.protocol.v1.MissionService/CreateMissionRecord"
	MissionService_GetMissionRecord_FullMethodName     = "/trajectory.trajectory_gcs_service.protocol.v1.MissionService/GetMissionRecord"
	MissionService_DeleteMissionRecord_FullMethodName  = "/trajectory.trajectory_gcs_service.protocol.v1.MissionService/DeleteMissionRecord"
	MissionService_UpdateMissionRecord_FullMethodName  = "/trajectory.trajectory_gcs_service.protocol.v1.MissionService/UpdateMissionRecord"
	MissionService_ListMissionIDRecords_FullMethodName = "/trajectory.trajectory_gcs_service.protocol.v1.MissionService/ListMissionIDRecords"
)

// MissionServiceClient is the client API for MissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MissionServiceClient interface {
	// CreateMission: ミッションの生成
	CreateMission(ctx context.Context, in *CreateMissionRequest, opts ...grpc.CallOption) (*CreateMissionResponse, error)
	// CreateMissionRecord: ミッションをデータベースに追加する
	CreateMissionRecord(ctx context.Context, in *MissionRecord, opts ...grpc.CallOption) (*MissionID, error)
	// GetMissionRecord: ミッションの取得
	GetMissionRecord(ctx context.Context, in *MissionID, opts ...grpc.CallOption) (*MissionRecord, error)
	// DeleteMission: ミッションの削除
	DeleteMissionRecord(ctx context.Context, in *MissionID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// UpdateMission: ミッションの更新
	UpdateMissionRecord(ctx context.Context, in *UpdateMissionRecordRequest, opts ...grpc.CallOption) (*MissionRecord, error)
	// ListMissionIDRecords: ミッションIDのリスト取得
	ListMissionIDRecords(ctx context.Context, in *ListMissionIDRecordsRequest, opts ...grpc.CallOption) (*ListMissionIDRecordsResponse, error)
}

type missionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMissionServiceClient(cc grpc.ClientConnInterface) MissionServiceClient {
	return &missionServiceClient{cc}
}

func (c *missionServiceClient) CreateMission(ctx context.Context, in *CreateMissionRequest, opts ...grpc.CallOption) (*CreateMissionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMissionResponse)
	err := c.cc.Invoke(ctx, MissionService_CreateMission_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) CreateMissionRecord(ctx context.Context, in *MissionRecord, opts ...grpc.CallOption) (*MissionID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MissionID)
	err := c.cc.Invoke(ctx, MissionService_CreateMissionRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) GetMissionRecord(ctx context.Context, in *MissionID, opts ...grpc.CallOption) (*MissionRecord, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MissionRecord)
	err := c.cc.Invoke(ctx, MissionService_GetMissionRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) DeleteMissionRecord(ctx context.Context, in *MissionID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MissionService_DeleteMissionRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) UpdateMissionRecord(ctx context.Context, in *UpdateMissionRecordRequest, opts ...grpc.CallOption) (*MissionRecord, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MissionRecord)
	err := c.cc.Invoke(ctx, MissionService_UpdateMissionRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *missionServiceClient) ListMissionIDRecords(ctx context.Context, in *ListMissionIDRecordsRequest, opts ...grpc.CallOption) (*ListMissionIDRecordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListMissionIDRecordsResponse)
	err := c.cc.Invoke(ctx, MissionService_ListMissionIDRecords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MissionServiceServer is the server API for MissionService service.
// All implementations must embed UnimplementedMissionServiceServer
// for forward compatibility
type MissionServiceServer interface {
	// CreateMission: ミッションの生成
	CreateMission(context.Context, *CreateMissionRequest) (*CreateMissionResponse, error)
	// CreateMissionRecord: ミッションをデータベースに追加する
	CreateMissionRecord(context.Context, *MissionRecord) (*MissionID, error)
	// GetMissionRecord: ミッションの取得
	GetMissionRecord(context.Context, *MissionID) (*MissionRecord, error)
	// DeleteMission: ミッションの削除
	DeleteMissionRecord(context.Context, *MissionID) (*emptypb.Empty, error)
	// UpdateMission: ミッションの更新
	UpdateMissionRecord(context.Context, *UpdateMissionRecordRequest) (*MissionRecord, error)
	// ListMissionIDRecords: ミッションIDのリスト取得
	ListMissionIDRecords(context.Context, *ListMissionIDRecordsRequest) (*ListMissionIDRecordsResponse, error)
	mustEmbedUnimplementedMissionServiceServer()
}

// UnimplementedMissionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMissionServiceServer struct {
}

func (UnimplementedMissionServiceServer) CreateMission(context.Context, *CreateMissionRequest) (*CreateMissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMission not implemented")
}
func (UnimplementedMissionServiceServer) CreateMissionRecord(context.Context, *MissionRecord) (*MissionID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMissionRecord not implemented")
}
func (UnimplementedMissionServiceServer) GetMissionRecord(context.Context, *MissionID) (*MissionRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMissionRecord not implemented")
}
func (UnimplementedMissionServiceServer) DeleteMissionRecord(context.Context, *MissionID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMissionRecord not implemented")
}
func (UnimplementedMissionServiceServer) UpdateMissionRecord(context.Context, *UpdateMissionRecordRequest) (*MissionRecord, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMissionRecord not implemented")
}
func (UnimplementedMissionServiceServer) ListMissionIDRecords(context.Context, *ListMissionIDRecordsRequest) (*ListMissionIDRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMissionIDRecords not implemented")
}
func (UnimplementedMissionServiceServer) mustEmbedUnimplementedMissionServiceServer() {}

// UnsafeMissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MissionServiceServer will
// result in compilation errors.
type UnsafeMissionServiceServer interface {
	mustEmbedUnimplementedMissionServiceServer()
}

func RegisterMissionServiceServer(s grpc.ServiceRegistrar, srv MissionServiceServer) {
	s.RegisterService(&MissionService_ServiceDesc, srv)
}

func _MissionService_CreateMission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).CreateMission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_CreateMission_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).CreateMission(ctx, req.(*CreateMissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_CreateMissionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MissionRecord)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).CreateMissionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_CreateMissionRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).CreateMissionRecord(ctx, req.(*MissionRecord))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_GetMissionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MissionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).GetMissionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_GetMissionRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).GetMissionRecord(ctx, req.(*MissionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_DeleteMissionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MissionID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).DeleteMissionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_DeleteMissionRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).DeleteMissionRecord(ctx, req.(*MissionID))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_UpdateMissionRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMissionRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).UpdateMissionRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_UpdateMissionRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).UpdateMissionRecord(ctx, req.(*UpdateMissionRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MissionService_ListMissionIDRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMissionIDRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MissionServiceServer).ListMissionIDRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MissionService_ListMissionIDRecords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MissionServiceServer).ListMissionIDRecords(ctx, req.(*ListMissionIDRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MissionService_ServiceDesc is the grpc.ServiceDesc for MissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trajectory.trajectory_gcs_service.protocol.v1.MissionService",
	HandlerType: (*MissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMission",
			Handler:    _MissionService_CreateMission_Handler,
		},
		{
			MethodName: "CreateMissionRecord",
			Handler:    _MissionService_CreateMissionRecord_Handler,
		},
		{
			MethodName: "GetMissionRecord",
			Handler:    _MissionService_GetMissionRecord_Handler,
		},
		{
			MethodName: "DeleteMissionRecord",
			Handler:    _MissionService_DeleteMissionRecord_Handler,
		},
		{
			MethodName: "UpdateMissionRecord",
			Handler:    _MissionService_UpdateMissionRecord_Handler,
		},
		{
			MethodName: "ListMissionIDRecords",
			Handler:    _MissionService_ListMissionIDRecords_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trajectory/trajectory_gcs_service/protocol/v1/mission_service.proto",
}
