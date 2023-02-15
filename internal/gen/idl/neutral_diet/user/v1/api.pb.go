// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: neutral_diet/user/v1/api.proto

package userv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AddFoodItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodLogItem *FoodLogItemRequest `protobuf:"bytes,1,opt,name=food_log_item,json=foodLogItem,proto3" json:"food_log_item,omitempty"`
}

func (x *AddFoodItemRequest) Reset() {
	*x = AddFoodItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFoodItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFoodItemRequest) ProtoMessage() {}

func (x *AddFoodItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFoodItemRequest.ProtoReflect.Descriptor instead.
func (*AddFoodItemRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *AddFoodItemRequest) GetFoodLogItem() *FoodLogItemRequest {
	if x != nil {
		return x.FoodLogItem
	}
	return nil
}

type AddFoodItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddFoodItemResponse) Reset() {
	*x = AddFoodItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddFoodItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddFoodItemResponse) ProtoMessage() {}

func (x *AddFoodItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddFoodItemResponse.ProtoReflect.Descriptor instead.
func (*AddFoodItemResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AddFoodItemResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteFoodItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteFoodItemRequest) Reset() {
	*x = DeleteFoodItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodItemRequest) ProtoMessage() {}

func (x *DeleteFoodItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteFoodItemRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteFoodItemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteFoodItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteFoodItemResponse) Reset() {
	*x = DeleteFoodItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodItemResponse) ProtoMessage() {}

func (x *DeleteFoodItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteFoodItemResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{3}
}

type GetFoodItemLogRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date *Date `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *GetFoodItemLogRequest) Reset() {
	*x = GetFoodItemLogRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodItemLogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodItemLogRequest) ProtoMessage() {}

func (x *GetFoodItemLogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodItemLogRequest.ProtoReflect.Descriptor instead.
func (*GetFoodItemLogRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetFoodItemLogRequest) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

type GetFoodItemLogResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodItemLog []*FoodLogItemResponse `protobuf:"bytes,1,rep,name=food_item_log,json=foodItemLog,proto3" json:"food_item_log,omitempty"`
}

func (x *GetFoodItemLogResponse) Reset() {
	*x = GetFoodItemLogResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodItemLogResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodItemLogResponse) ProtoMessage() {}

func (x *GetFoodItemLogResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodItemLogResponse.ProtoReflect.Descriptor instead.
func (*GetFoodItemLogResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{5}
}

func (x *GetFoodItemLogResponse) GetFoodItemLog() []*FoodLogItemResponse {
	if x != nil {
		return x.FoodItemLog
	}
	return nil
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirebaseUid string `protobuf:"bytes,1,opt,name=firebase_uid,json=firebaseUid,proto3" json:"firebase_uid,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{6}
}

func (x *CreateUserRequest) GetFirebaseUid() string {
	if x != nil {
		return x.FirebaseUid
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{7}
}

func (x *CreateUserResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{8}
}

type DeleteUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{9}
}

type UpdateUserSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSettings *UserSettings `protobuf:"bytes,1,opt,name=user_settings,json=userSettings,proto3" json:"user_settings,omitempty"`
}

func (x *UpdateUserSettingsRequest) Reset() {
	*x = UpdateUserSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserSettingsRequest) ProtoMessage() {}

func (x *UpdateUserSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserSettingsRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserSettingsRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateUserSettingsRequest) GetUserSettings() *UserSettings {
	if x != nil {
		return x.UserSettings
	}
	return nil
}

type UpdateUserSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserSettingsResponse) Reset() {
	*x = UpdateUserSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserSettingsResponse) ProtoMessage() {}

func (x *UpdateUserSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserSettingsResponse.ProtoReflect.Descriptor instead.
func (*UpdateUserSettingsResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{11}
}

type GetUserSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserSettingsRequest) Reset() {
	*x = GetUserSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingsRequest) ProtoMessage() {}

func (x *GetUserSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetUserSettingsRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{12}
}

type GetUserSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserSettings *UserSettings `protobuf:"bytes,1,opt,name=user_settings,json=userSettings,proto3" json:"user_settings,omitempty"`
}

func (x *GetUserSettingsResponse) Reset() {
	*x = GetUserSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_api_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserSettingsResponse) ProtoMessage() {}

func (x *GetUserSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_api_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetUserSettingsResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_api_proto_rawDescGZIP(), []int{13}
}

func (x *GetUserSettingsResponse) GetUserSettings() *UserSettings {
	if x != nil {
		return x.UserSettings
	}
	return nil
}

var File_neutral_diet_user_v1_api_proto protoreflect.FileDescriptor

var file_neutral_diet_user_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f,
	0x64, 0x69, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f,
	0x6f, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x12, 0x41,
	0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x56, 0x0a, 0x0d, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x6c, 0x6f, 0x67, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x66, 0x6f,
	0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x25, 0x0a, 0x13, 0x41, 0x64, 0x64,
	0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x30, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x00, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69,
	0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22,
	0x67, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0d, 0x66, 0x6f, 0x6f,
	0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x6c, 0x6f, 0x67, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x6f,
	0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x22, 0x3f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a,
	0x0c, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x66, 0x69,
	0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x55, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6e, 0x0a, 0x19, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x51, 0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0c, 0x75, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x6c, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a,
	0x0d, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01, 0x02,
	0x10, 0x01, 0x52, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x32, 0x84, 0x06, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x64, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x28, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x6e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x64, 0x64, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f,
	0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6d, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x12, 0x2b, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61,
	0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6e, 0x65,
	0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x27, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f,
	0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28,
	0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2c, 0x2e,
	0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6e, 0x65,
	0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x12, 0x2f, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69,
	0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xdf, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e,
	0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d,
	0x30, 0x32, 0x7a, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69, 0x64,
	0x6c, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x4e, 0x55, 0x58, 0xaa, 0x02, 0x13, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65,
	0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x4e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x1f, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x55, 0x73,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x15, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x3a,
	0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_neutral_diet_user_v1_api_proto_rawDescOnce sync.Once
	file_neutral_diet_user_v1_api_proto_rawDescData = file_neutral_diet_user_v1_api_proto_rawDesc
)

func file_neutral_diet_user_v1_api_proto_rawDescGZIP() []byte {
	file_neutral_diet_user_v1_api_proto_rawDescOnce.Do(func() {
		file_neutral_diet_user_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_neutral_diet_user_v1_api_proto_rawDescData)
	})
	return file_neutral_diet_user_v1_api_proto_rawDescData
}

var file_neutral_diet_user_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_neutral_diet_user_v1_api_proto_goTypes = []interface{}{
	(*AddFoodItemRequest)(nil),         // 0: neutral_diet.user.v1.AddFoodItemRequest
	(*AddFoodItemResponse)(nil),        // 1: neutral_diet.user.v1.AddFoodItemResponse
	(*DeleteFoodItemRequest)(nil),      // 2: neutral_diet.user.v1.DeleteFoodItemRequest
	(*DeleteFoodItemResponse)(nil),     // 3: neutral_diet.user.v1.DeleteFoodItemResponse
	(*GetFoodItemLogRequest)(nil),      // 4: neutral_diet.user.v1.GetFoodItemLogRequest
	(*GetFoodItemLogResponse)(nil),     // 5: neutral_diet.user.v1.GetFoodItemLogResponse
	(*CreateUserRequest)(nil),          // 6: neutral_diet.user.v1.CreateUserRequest
	(*CreateUserResponse)(nil),         // 7: neutral_diet.user.v1.CreateUserResponse
	(*DeleteUserRequest)(nil),          // 8: neutral_diet.user.v1.DeleteUserRequest
	(*DeleteUserResponse)(nil),         // 9: neutral_diet.user.v1.DeleteUserResponse
	(*UpdateUserSettingsRequest)(nil),  // 10: neutral_diet.user.v1.UpdateUserSettingsRequest
	(*UpdateUserSettingsResponse)(nil), // 11: neutral_diet.user.v1.UpdateUserSettingsResponse
	(*GetUserSettingsRequest)(nil),     // 12: neutral_diet.user.v1.GetUserSettingsRequest
	(*GetUserSettingsResponse)(nil),    // 13: neutral_diet.user.v1.GetUserSettingsResponse
	(*FoodLogItemRequest)(nil),         // 14: neutral_diet.user.v1.FoodLogItemRequest
	(*Date)(nil),                       // 15: neutral_diet.user.v1.Date
	(*FoodLogItemResponse)(nil),        // 16: neutral_diet.user.v1.FoodLogItemResponse
	(*UserSettings)(nil),               // 17: neutral_diet.user.v1.UserSettings
}
var file_neutral_diet_user_v1_api_proto_depIdxs = []int32{
	14, // 0: neutral_diet.user.v1.AddFoodItemRequest.food_log_item:type_name -> neutral_diet.user.v1.FoodLogItemRequest
	15, // 1: neutral_diet.user.v1.GetFoodItemLogRequest.date:type_name -> neutral_diet.user.v1.Date
	16, // 2: neutral_diet.user.v1.GetFoodItemLogResponse.food_item_log:type_name -> neutral_diet.user.v1.FoodLogItemResponse
	17, // 3: neutral_diet.user.v1.UpdateUserSettingsRequest.user_settings:type_name -> neutral_diet.user.v1.UserSettings
	17, // 4: neutral_diet.user.v1.GetUserSettingsResponse.user_settings:type_name -> neutral_diet.user.v1.UserSettings
	0,  // 5: neutral_diet.user.v1.UserService.AddFoodItem:input_type -> neutral_diet.user.v1.AddFoodItemRequest
	2,  // 6: neutral_diet.user.v1.UserService.DeleteFoodItem:input_type -> neutral_diet.user.v1.DeleteFoodItemRequest
	4,  // 7: neutral_diet.user.v1.UserService.GetFoodItemLog:input_type -> neutral_diet.user.v1.GetFoodItemLogRequest
	6,  // 8: neutral_diet.user.v1.UserService.CreateUser:input_type -> neutral_diet.user.v1.CreateUserRequest
	8,  // 9: neutral_diet.user.v1.UserService.DeleteUser:input_type -> neutral_diet.user.v1.DeleteUserRequest
	12, // 10: neutral_diet.user.v1.UserService.GetUserSettings:input_type -> neutral_diet.user.v1.GetUserSettingsRequest
	10, // 11: neutral_diet.user.v1.UserService.UpdateUserSettings:input_type -> neutral_diet.user.v1.UpdateUserSettingsRequest
	1,  // 12: neutral_diet.user.v1.UserService.AddFoodItem:output_type -> neutral_diet.user.v1.AddFoodItemResponse
	3,  // 13: neutral_diet.user.v1.UserService.DeleteFoodItem:output_type -> neutral_diet.user.v1.DeleteFoodItemResponse
	5,  // 14: neutral_diet.user.v1.UserService.GetFoodItemLog:output_type -> neutral_diet.user.v1.GetFoodItemLogResponse
	7,  // 15: neutral_diet.user.v1.UserService.CreateUser:output_type -> neutral_diet.user.v1.CreateUserResponse
	9,  // 16: neutral_diet.user.v1.UserService.DeleteUser:output_type -> neutral_diet.user.v1.DeleteUserResponse
	13, // 17: neutral_diet.user.v1.UserService.GetUserSettings:output_type -> neutral_diet.user.v1.GetUserSettingsResponse
	11, // 18: neutral_diet.user.v1.UserService.UpdateUserSettings:output_type -> neutral_diet.user.v1.UpdateUserSettingsResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_neutral_diet_user_v1_api_proto_init() }
func file_neutral_diet_user_v1_api_proto_init() {
	if File_neutral_diet_user_v1_api_proto != nil {
		return
	}
	file_neutral_diet_user_v1_date_proto_init()
	file_neutral_diet_user_v1_food_item_log_proto_init()
	file_neutral_diet_user_v1_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_neutral_diet_user_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFoodItemRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddFoodItemResponse); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodItemRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodItemResponse); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodItemLogRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodItemLogResponse); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserResponse); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserSettingsRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserSettingsResponse); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSettingsRequest); i {
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
		file_neutral_diet_user_v1_api_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserSettingsResponse); i {
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
			RawDescriptor: file_neutral_diet_user_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neutral_diet_user_v1_api_proto_goTypes,
		DependencyIndexes: file_neutral_diet_user_v1_api_proto_depIdxs,
		MessageInfos:      file_neutral_diet_user_v1_api_proto_msgTypes,
	}.Build()
	File_neutral_diet_user_v1_api_proto = out.File
	file_neutral_diet_user_v1_api_proto_rawDesc = nil
	file_neutral_diet_user_v1_api_proto_goTypes = nil
	file_neutral_diet_user_v1_api_proto_depIdxs = nil
}
