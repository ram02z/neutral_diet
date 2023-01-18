// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: neutral_diet/user/v1/food_item_log.proto

package userv1

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

type FoodLogItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodItemId      int32   `protobuf:"varint,1,opt,name=food_item_id,json=foodItemId,proto3" json:"food_item_id,omitempty"`
	Weight          float64 `protobuf:"fixed64,2,opt,name=weight,proto3" json:"weight,omitempty"`
	CarbonFootprint float64 `protobuf:"fixed64,3,opt,name=carbon_footprint,json=carbonFootprint,proto3" json:"carbon_footprint,omitempty"`
}

func (x *FoodLogItem) Reset() {
	*x = FoodLogItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_food_item_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodLogItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodLogItem) ProtoMessage() {}

func (x *FoodLogItem) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_food_item_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodLogItem.ProtoReflect.Descriptor instead.
func (*FoodLogItem) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_food_item_log_proto_rawDescGZIP(), []int{0}
}

func (x *FoodLogItem) GetFoodItemId() int32 {
	if x != nil {
		return x.FoodItemId
	}
	return 0
}

func (x *FoodLogItem) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *FoodLogItem) GetCarbonFootprint() float64 {
	if x != nil {
		return x.CarbonFootprint
	}
	return 0
}

var File_neutral_diet_user_v1_food_item_log_proto protoreflect.FileDescriptor

var file_neutral_diet_user_v1_food_item_log_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x22, 0x72, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x20, 0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x72,
	0x62, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x74, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x42, 0xe7, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x75,
	0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x46, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x30, 0x32, 0x7a, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x65, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x4e, 0x55, 0x58, 0xaa, 0x02, 0x13, 0x4e, 0x65, 0x75, 0x74, 0x72,
	0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x13, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x55, 0x73, 0x65,
	0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1f, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69,
	0x65, 0x74, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c,
	0x44, 0x69, 0x65, 0x74, 0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_neutral_diet_user_v1_food_item_log_proto_rawDescOnce sync.Once
	file_neutral_diet_user_v1_food_item_log_proto_rawDescData = file_neutral_diet_user_v1_food_item_log_proto_rawDesc
)

func file_neutral_diet_user_v1_food_item_log_proto_rawDescGZIP() []byte {
	file_neutral_diet_user_v1_food_item_log_proto_rawDescOnce.Do(func() {
		file_neutral_diet_user_v1_food_item_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_neutral_diet_user_v1_food_item_log_proto_rawDescData)
	})
	return file_neutral_diet_user_v1_food_item_log_proto_rawDescData
}

var file_neutral_diet_user_v1_food_item_log_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_neutral_diet_user_v1_food_item_log_proto_goTypes = []interface{}{
	(*FoodLogItem)(nil), // 0: neutral_diet.user.v1.FoodLogItem
}
var file_neutral_diet_user_v1_food_item_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_neutral_diet_user_v1_food_item_log_proto_init() }
func file_neutral_diet_user_v1_food_item_log_proto_init() {
	if File_neutral_diet_user_v1_food_item_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_neutral_diet_user_v1_food_item_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodLogItem); i {
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
			RawDescriptor: file_neutral_diet_user_v1_food_item_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_neutral_diet_user_v1_food_item_log_proto_goTypes,
		DependencyIndexes: file_neutral_diet_user_v1_food_item_log_proto_depIdxs,
		MessageInfos:      file_neutral_diet_user_v1_food_item_log_proto_msgTypes,
	}.Build()
	File_neutral_diet_user_v1_food_item_log_proto = out.File
	file_neutral_diet_user_v1_food_item_log_proto_rawDesc = nil
	file_neutral_diet_user_v1_food_item_log_proto_goTypes = nil
	file_neutral_diet_user_v1_food_item_log_proto_depIdxs = nil
}
