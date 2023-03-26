// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: neutral_diet/job/v1/api.proto

package jobv1

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

type SendGoalNotificationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendGoalNotificationsRequest) Reset() {
	*x = SendGoalNotificationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_job_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendGoalNotificationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendGoalNotificationsRequest) ProtoMessage() {}

func (x *SendGoalNotificationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_job_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendGoalNotificationsRequest.ProtoReflect.Descriptor instead.
func (*SendGoalNotificationsRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_job_v1_api_proto_rawDescGZIP(), []int{0}
}

type SendGoalNotificationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendGoalNotificationsResponse) Reset() {
	*x = SendGoalNotificationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_job_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendGoalNotificationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendGoalNotificationsResponse) ProtoMessage() {}

func (x *SendGoalNotificationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_job_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendGoalNotificationsResponse.ProtoReflect.Descriptor instead.
func (*SendGoalNotificationsResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_job_v1_api_proto_rawDescGZIP(), []int{1}
}

type MarkCompletedGoalsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkCompletedGoalsRequest) Reset() {
	*x = MarkCompletedGoalsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_job_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkCompletedGoalsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkCompletedGoalsRequest) ProtoMessage() {}

func (x *MarkCompletedGoalsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_job_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkCompletedGoalsRequest.ProtoReflect.Descriptor instead.
func (*MarkCompletedGoalsRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_job_v1_api_proto_rawDescGZIP(), []int{2}
}

type MarkCompletedGoalsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *MarkCompletedGoalsResponse) Reset() {
	*x = MarkCompletedGoalsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_job_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkCompletedGoalsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkCompletedGoalsResponse) ProtoMessage() {}

func (x *MarkCompletedGoalsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_job_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkCompletedGoalsResponse.ProtoReflect.Descriptor instead.
func (*MarkCompletedGoalsResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_job_v1_api_proto_rawDescGZIP(), []int{3}
}

var File_neutral_diet_job_v1_api_proto protoreflect.FileDescriptor

var file_neutral_diet_job_v1_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x6a,
	0x6f, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x6a, 0x6f,
	0x62, 0x2e, 0x76, 0x31, 0x22, 0x1e, 0x0a, 0x1c, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x61, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x61, 0x6c,
	0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x0a, 0x19, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0x88, 0x02, 0x0a, 0x0a, 0x4a, 0x6f, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x80, 0x01, 0x0a, 0x15, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x2e, 0x6e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6e,
	0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x47, 0x6f, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x77, 0x0a, 0x12, 0x4d, 0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x47, 0x6f, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x61, 0x72, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x47, 0x6f, 0x61, 0x6c,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0xd8, 0x01, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74,
	0x2e, 0x6a, 0x6f, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x61, 0x6d, 0x30, 0x32, 0x7a, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2f, 0x6a, 0x6f, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x6a, 0x6f, 0x62, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x4e, 0x4a, 0x58, 0xaa, 0x02, 0x12, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69,
	0x65, 0x74, 0x2e, 0x4a, 0x6f, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x12, 0x4e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x4a, 0x6f, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02,
	0x1e, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x4a, 0x6f, 0x62,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x14, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x3a, 0x3a, 0x4a,
	0x6f, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neutral_diet_job_v1_api_proto_rawDescOnce sync.Once
	file_neutral_diet_job_v1_api_proto_rawDescData = file_neutral_diet_job_v1_api_proto_rawDesc
)

func file_neutral_diet_job_v1_api_proto_rawDescGZIP() []byte {
	file_neutral_diet_job_v1_api_proto_rawDescOnce.Do(func() {
		file_neutral_diet_job_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_neutral_diet_job_v1_api_proto_rawDescData)
	})
	return file_neutral_diet_job_v1_api_proto_rawDescData
}

var file_neutral_diet_job_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_neutral_diet_job_v1_api_proto_goTypes = []interface{}{
	(*SendGoalNotificationsRequest)(nil),  // 0: neutral_diet.job.v1.SendGoalNotificationsRequest
	(*SendGoalNotificationsResponse)(nil), // 1: neutral_diet.job.v1.SendGoalNotificationsResponse
	(*MarkCompletedGoalsRequest)(nil),     // 2: neutral_diet.job.v1.MarkCompletedGoalsRequest
	(*MarkCompletedGoalsResponse)(nil),    // 3: neutral_diet.job.v1.MarkCompletedGoalsResponse
}
var file_neutral_diet_job_v1_api_proto_depIdxs = []int32{
	0, // 0: neutral_diet.job.v1.JobService.SendGoalNotifications:input_type -> neutral_diet.job.v1.SendGoalNotificationsRequest
	2, // 1: neutral_diet.job.v1.JobService.MarkCompletedGoals:input_type -> neutral_diet.job.v1.MarkCompletedGoalsRequest
	1, // 2: neutral_diet.job.v1.JobService.SendGoalNotifications:output_type -> neutral_diet.job.v1.SendGoalNotificationsResponse
	3, // 3: neutral_diet.job.v1.JobService.MarkCompletedGoals:output_type -> neutral_diet.job.v1.MarkCompletedGoalsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_neutral_diet_job_v1_api_proto_init() }
func file_neutral_diet_job_v1_api_proto_init() {
	if File_neutral_diet_job_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neutral_diet_job_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendGoalNotificationsRequest); i {
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
		file_neutral_diet_job_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendGoalNotificationsResponse); i {
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
		file_neutral_diet_job_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkCompletedGoalsRequest); i {
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
		file_neutral_diet_job_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkCompletedGoalsResponse); i {
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
			RawDescriptor: file_neutral_diet_job_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_neutral_diet_job_v1_api_proto_goTypes,
		DependencyIndexes: file_neutral_diet_job_v1_api_proto_depIdxs,
		MessageInfos:      file_neutral_diet_job_v1_api_proto_msgTypes,
	}.Build()
	File_neutral_diet_job_v1_api_proto = out.File
	file_neutral_diet_job_v1_api_proto_rawDesc = nil
	file_neutral_diet_job_v1_api_proto_goTypes = nil
	file_neutral_diet_job_v1_api_proto_depIdxs = nil
}
