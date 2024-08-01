// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: trajectory/trajectory_gcs_service/protocol/v1/user_service.proto

package trajectory_gcs

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserPermission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Admin           bool `protobuf:"varint,1,opt,name=admin,proto3" json:"admin,omitempty"`
	Operator        bool `protobuf:"varint,2,opt,name=operator,proto3" json:"operator,omitempty"`
	MapSelectable   bool `protobuf:"varint,3,opt,name=map_selectable,json=mapSelectable,proto3" json:"map_selectable,omitempty"`
	SaveSetting     bool `protobuf:"varint,4,opt,name=save_setting,json=saveSetting,proto3" json:"save_setting,omitempty"`
	Movie           bool `protobuf:"varint,5,opt,name=movie,proto3" json:"movie,omitempty"`
	DroneAuth       bool `protobuf:"varint,6,opt,name=drone_auth,json=droneAuth,proto3" json:"drone_auth,omitempty"`
	SitlAuth        bool `protobuf:"varint,7,opt,name=sitl_auth,json=sitlAuth,proto3" json:"sitl_auth,omitempty"`
	LadderAuth      bool `protobuf:"varint,8,opt,name=ladder_auth,json=ladderAuth,proto3" json:"ladder_auth,omitempty"`
	TileAuth        bool `protobuf:"varint,9,opt,name=tile_auth,json=tileAuth,proto3" json:"tile_auth,omitempty"`
	AnnotationAuth  bool `protobuf:"varint,10,opt,name=annotation_auth,json=annotationAuth,proto3" json:"annotation_auth,omitempty"`
	VideoAuth       bool `protobuf:"varint,11,opt,name=video_auth,json=videoAuth,proto3" json:"video_auth,omitempty"`
	VideoUser       bool `protobuf:"varint,12,opt,name=video_user,json=videoUser,proto3" json:"video_user,omitempty"`
	SystemUser      bool `protobuf:"varint,13,opt,name=system_user,json=systemUser,proto3" json:"system_user,omitempty"`
	SystemApiAccess bool `protobuf:"varint,14,opt,name=system_api_access,json=systemApiAccess,proto3" json:"system_api_access,omitempty"`
}

func (x *UserPermission) Reset() {
	*x = UserPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPermission) ProtoMessage() {}

func (x *UserPermission) ProtoReflect() protoreflect.Message {
	mi := &file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPermission.ProtoReflect.Descriptor instead.
func (*UserPermission) Descriptor() ([]byte, []int) {
	return file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *UserPermission) GetAdmin() bool {
	if x != nil {
		return x.Admin
	}
	return false
}

func (x *UserPermission) GetOperator() bool {
	if x != nil {
		return x.Operator
	}
	return false
}

func (x *UserPermission) GetMapSelectable() bool {
	if x != nil {
		return x.MapSelectable
	}
	return false
}

func (x *UserPermission) GetSaveSetting() bool {
	if x != nil {
		return x.SaveSetting
	}
	return false
}

func (x *UserPermission) GetMovie() bool {
	if x != nil {
		return x.Movie
	}
	return false
}

func (x *UserPermission) GetDroneAuth() bool {
	if x != nil {
		return x.DroneAuth
	}
	return false
}

func (x *UserPermission) GetSitlAuth() bool {
	if x != nil {
		return x.SitlAuth
	}
	return false
}

func (x *UserPermission) GetLadderAuth() bool {
	if x != nil {
		return x.LadderAuth
	}
	return false
}

func (x *UserPermission) GetTileAuth() bool {
	if x != nil {
		return x.TileAuth
	}
	return false
}

func (x *UserPermission) GetAnnotationAuth() bool {
	if x != nil {
		return x.AnnotationAuth
	}
	return false
}

func (x *UserPermission) GetVideoAuth() bool {
	if x != nil {
		return x.VideoAuth
	}
	return false
}

func (x *UserPermission) GetVideoUser() bool {
	if x != nil {
		return x.VideoUser
	}
	return false
}

func (x *UserPermission) GetSystemUser() bool {
	if x != nil {
		return x.SystemUser
	}
	return false
}

func (x *UserPermission) GetSystemApiAccess() bool {
	if x != nil {
		return x.SystemApiAccess
	}
	return false
}

type UserPermissionWithName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId          string          `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User_Permission *UserPermission `protobuf:"bytes,2,opt,name=user_Permission,json=userPermission,proto3" json:"user_Permission,omitempty"`
}

func (x *UserPermissionWithName) Reset() {
	*x = UserPermissionWithName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPermissionWithName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPermissionWithName) ProtoMessage() {}

func (x *UserPermissionWithName) ProtoReflect() protoreflect.Message {
	mi := &file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPermissionWithName.ProtoReflect.Descriptor instead.
func (*UserPermissionWithName) Descriptor() ([]byte, []int) {
	return file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *UserPermissionWithName) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserPermissionWithName) GetUser_Permission() *UserPermission {
	if x != nil {
		return x.User_Permission
	}
	return nil
}

type ListUserPermissionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserPermissionWithNameList []*UserPermission `protobuf:"bytes,1,rep,name=user_permission_with_name_list,json=userPermissionWithNameList,proto3" json:"user_permission_with_name_list,omitempty"`
}

func (x *ListUserPermissionsResponse) Reset() {
	*x = ListUserPermissionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUserPermissionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUserPermissionsResponse) ProtoMessage() {}

func (x *ListUserPermissionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUserPermissionsResponse.ProtoReflect.Descriptor instead.
func (*ListUserPermissionsResponse) Descriptor() ([]byte, []int) {
	return file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListUserPermissionsResponse) GetUserPermissionWithNameList() []*UserPermission {
	if x != nil {
		return x.UserPermissionWithNameList
	}
	return nil
}

var File_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto protoreflect.FileDescriptor

var file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDesc = []byte{
	0x0a, 0x40, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x74, 0x72, 0x61,
	0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2d, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x74,
	0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0,
	0x03, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x6d, 0x61, 0x70, 0x5f, 0x73, 0x65, 0x6c, 0x65, 0x63,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6d, 0x61, 0x70,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x61,
	0x76, 0x65, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0b, 0x73, 0x61, 0x76, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6d, 0x6f,
	0x76, 0x69, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x5f, 0x61, 0x75, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x64, 0x72, 0x6f, 0x6e, 0x65, 0x41, 0x75,
	0x74, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x69, 0x74, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x69, 0x74, 0x6c, 0x41, 0x75, 0x74, 0x68, 0x12,
	0x1f, 0x0a, 0x0b, 0x6c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6c, 0x61, 0x64, 0x64, 0x65, 0x72, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x1b, 0x0a, 0x09, 0x74, 0x69, 0x6c, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x74, 0x69, 0x6c, 0x65, 0x41, 0x75, 0x74, 0x68, 0x12, 0x27, 0x0a,
	0x0f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x75, 0x74, 0x68,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f,
	0x61, 0x75, 0x74, 0x68, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x55, 0x73, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f,
	0x61, 0x70, 0x69, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x70, 0x69, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x22, 0x99, 0x01, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x66, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d,
	0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6a,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x75,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xa1, 0x01,
	0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x81, 0x01,
	0x0a, 0x1e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67,
	0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x1a, 0x75, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x69, 0x73,
	0x74, 0x32, 0xa8, 0x03, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x3d, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x7b, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x4a, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63,
	0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0xa6, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x2e, 0x74,
	0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x45, 0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79,
	0x2e, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63, 0x73, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x57, 0x69, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x00, 0x42, 0x39, 0x5a, 0x37,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6a, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x79, 0x6a, 0x70, 0x2f, 0x74, 0x72, 0x6a, 0x78, 0x5f, 0x61, 0x70, 0x69,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x79, 0x5f, 0x67, 0x63, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescOnce sync.Once
	file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescData = file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDesc
)

func file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescGZIP() []byte {
	file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescOnce.Do(func() {
		file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescData)
	})
	return file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDescData
}

var file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_goTypes = []any{
	(*UserPermission)(nil),              // 0: trajectory.trajectory_gcs_service.protocol.v1.UserPermission
	(*UserPermissionWithName)(nil),      // 1: trajectory.trajectory_gcs_service.protocol.v1.UserPermissionWithName
	(*ListUserPermissionsResponse)(nil), // 2: trajectory.trajectory_gcs_service.protocol.v1.ListUserPermissionsResponse
	(*emptypb.Empty)(nil),               // 3: google.protobuf.Empty
}
var file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_depIdxs = []int32{
	0, // 0: trajectory.trajectory_gcs_service.protocol.v1.UserPermissionWithName.user_Permission:type_name -> trajectory.trajectory_gcs_service.protocol.v1.UserPermission
	0, // 1: trajectory.trajectory_gcs_service.protocol.v1.ListUserPermissionsResponse.user_permission_with_name_list:type_name -> trajectory.trajectory_gcs_service.protocol.v1.UserPermission
	3, // 2: trajectory.trajectory_gcs_service.protocol.v1.UserAccountService.GetUserPermission:input_type -> google.protobuf.Empty
	3, // 3: trajectory.trajectory_gcs_service.protocol.v1.UserAccountService.ListUserPermissions:input_type -> google.protobuf.Empty
	1, // 4: trajectory.trajectory_gcs_service.protocol.v1.UserAccountService.UpdateUserPermission:input_type -> trajectory.trajectory_gcs_service.protocol.v1.UserPermissionWithName
	0, // 5: trajectory.trajectory_gcs_service.protocol.v1.UserAccountService.GetUserPermission:output_type -> trajectory.trajectory_gcs_service.protocol.v1.UserPermission
	2, // 6: trajectory.trajectory_gcs_service.protocol.v1.UserAccountService.ListUserPermissions:output_type -> trajectory.trajectory_gcs_service.protocol.v1.ListUserPermissionsResponse
	1, // 7: trajectory.trajectory_gcs_service.protocol.v1.UserAccountService.UpdateUserPermission:output_type -> trajectory.trajectory_gcs_service.protocol.v1.UserPermissionWithName
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_init() }
func file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_init() {
	if File_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserPermission); i {
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
		file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*UserPermissionWithName); i {
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
		file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListUserPermissionsResponse); i {
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
			RawDescriptor: file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_goTypes,
		DependencyIndexes: file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_depIdxs,
		MessageInfos:      file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_msgTypes,
	}.Build()
	File_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto = out.File
	file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_rawDesc = nil
	file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_goTypes = nil
	file_trajectory_trajectory_gcs_service_protocol_v1_user_service_proto_depIdxs = nil
}
