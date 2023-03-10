// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: neutral_diet/user/v1/food_item_log.proto

package userv1

import (
	v1 "github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1"
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

type Unit int32

const (
	// Kilogram
	Unit_UNIT_UNSPECIFIED Unit = 0
	Unit_UNIT_GRAM        Unit = 1
	Unit_UNIT_OUNCE       Unit = 2
	Unit_UNIT_POUND       Unit = 3
	Unit_UNIT_LITRE       Unit = 4
	Unit_UNIT_MILLILITRE  Unit = 5
	Unit_UNIT_GALLON      Unit = 6
	Unit_UNIT_PINT        Unit = 7
)

// Enum value maps for Unit.
var (
	Unit_name = map[int32]string{
		0: "UNIT_UNSPECIFIED",
		1: "UNIT_GRAM",
		2: "UNIT_OUNCE",
		3: "UNIT_POUND",
		4: "UNIT_LITRE",
		5: "UNIT_MILLILITRE",
		6: "UNIT_GALLON",
		7: "UNIT_PINT",
	}
	Unit_value = map[string]int32{
		"UNIT_UNSPECIFIED": 0,
		"UNIT_GRAM":        1,
		"UNIT_OUNCE":       2,
		"UNIT_POUND":       3,
		"UNIT_LITRE":       4,
		"UNIT_MILLILITRE":  5,
		"UNIT_GALLON":      6,
		"UNIT_PINT":        7,
	}
)

func (x Unit) Enum() *Unit {
	p := new(Unit)
	*p = x
	return p
}

func (x Unit) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Unit) Descriptor() protoreflect.EnumDescriptor {
	return file_neutral_diet_user_v1_food_item_log_proto_enumTypes[0].Descriptor()
}

func (Unit) Type() protoreflect.EnumType {
	return &file_neutral_diet_user_v1_food_item_log_proto_enumTypes[0]
}

func (x Unit) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Unit.Descriptor instead.
func (Unit) EnumDescriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_food_item_log_proto_rawDescGZIP(), []int{0}
}

type Meal int32

const (
	Meal_MEAL_UNSPECIFIED Meal = 0
	Meal_MEAL_BREAKFAST   Meal = 1
	Meal_MEAL_LUNCH       Meal = 2
	Meal_MEAL_DINNER      Meal = 3
)

// Enum value maps for Meal.
var (
	Meal_name = map[int32]string{
		0: "MEAL_UNSPECIFIED",
		1: "MEAL_BREAKFAST",
		2: "MEAL_LUNCH",
		3: "MEAL_DINNER",
	}
	Meal_value = map[string]int32{
		"MEAL_UNSPECIFIED": 0,
		"MEAL_BREAKFAST":   1,
		"MEAL_LUNCH":       2,
		"MEAL_DINNER":      3,
	}
)

func (x Meal) Enum() *Meal {
	p := new(Meal)
	*p = x
	return p
}

func (x Meal) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Meal) Descriptor() protoreflect.EnumDescriptor {
	return file_neutral_diet_user_v1_food_item_log_proto_enumTypes[1].Descriptor()
}

func (Meal) Type() protoreflect.EnumType {
	return &file_neutral_diet_user_v1_food_item_log_proto_enumTypes[1]
}

func (x Meal) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Meal.Descriptor instead.
func (Meal) EnumDescriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_food_item_log_proto_rawDescGZIP(), []int{1}
}

type FoodLogItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FoodItemId int32     `protobuf:"varint,1,opt,name=food_item_id,json=foodItemId,proto3" json:"food_item_id,omitempty"`
	Quantity   float64   `protobuf:"fixed64,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Date       *Date     `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Unit       Unit      `protobuf:"varint,4,opt,name=unit,proto3,enum=neutral_diet.user.v1.Unit" json:"unit,omitempty"`
	Region     v1.Region `protobuf:"varint,5,opt,name=region,proto3,enum=neutral_diet.food.v1.Region" json:"region,omitempty"`
	Meal       Meal      `protobuf:"varint,6,opt,name=meal,proto3,enum=neutral_diet.user.v1.Meal" json:"meal,omitempty"`
}

func (x *FoodLogItemRequest) Reset() {
	*x = FoodLogItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_food_item_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodLogItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodLogItemRequest) ProtoMessage() {}

func (x *FoodLogItemRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use FoodLogItemRequest.ProtoReflect.Descriptor instead.
func (*FoodLogItemRequest) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_food_item_log_proto_rawDescGZIP(), []int{0}
}

func (x *FoodLogItemRequest) GetFoodItemId() int32 {
	if x != nil {
		return x.FoodItemId
	}
	return 0
}

func (x *FoodLogItemRequest) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *FoodLogItemRequest) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *FoodLogItemRequest) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNIT_UNSPECIFIED
}

func (x *FoodLogItemRequest) GetRegion() v1.Region {
	if x != nil {
		return x.Region
	}
	return v1.Region(0)
}

func (x *FoodLogItemRequest) GetMeal() Meal {
	if x != nil {
		return x.Meal
	}
	return Meal_MEAL_UNSPECIFIED
}

type FoodLogItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FoodItemId      int32     `protobuf:"varint,2,opt,name=food_item_id,json=foodItemId,proto3" json:"food_item_id,omitempty"`
	Name            string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Quantity        float64   `protobuf:"fixed64,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	CarbonFootprint float64   `protobuf:"fixed64,5,opt,name=carbon_footprint,json=carbonFootprint,proto3" json:"carbon_footprint,omitempty"`
	Date            *Date     `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Unit            Unit      `protobuf:"varint,7,opt,name=unit,proto3,enum=neutral_diet.user.v1.Unit" json:"unit,omitempty"`
	Region          v1.Region `protobuf:"varint,8,opt,name=region,proto3,enum=neutral_diet.food.v1.Region" json:"region,omitempty"`
	Meal            Meal      `protobuf:"varint,9,opt,name=meal,proto3,enum=neutral_diet.user.v1.Meal" json:"meal,omitempty"`
}

func (x *FoodLogItemResponse) Reset() {
	*x = FoodLogItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_neutral_diet_user_v1_food_item_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FoodLogItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FoodLogItemResponse) ProtoMessage() {}

func (x *FoodLogItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_neutral_diet_user_v1_food_item_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FoodLogItemResponse.ProtoReflect.Descriptor instead.
func (*FoodLogItemResponse) Descriptor() ([]byte, []int) {
	return file_neutral_diet_user_v1_food_item_log_proto_rawDescGZIP(), []int{1}
}

func (x *FoodLogItemResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FoodLogItemResponse) GetFoodItemId() int32 {
	if x != nil {
		return x.FoodItemId
	}
	return 0
}

func (x *FoodLogItemResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FoodLogItemResponse) GetQuantity() float64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *FoodLogItemResponse) GetCarbonFootprint() float64 {
	if x != nil {
		return x.CarbonFootprint
	}
	return 0
}

func (x *FoodLogItemResponse) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *FoodLogItemResponse) GetUnit() Unit {
	if x != nil {
		return x.Unit
	}
	return Unit_UNIT_UNSPECIFIED
}

func (x *FoodLogItemResponse) GetRegion() v1.Region {
	if x != nil {
		return x.Region
	}
	return v1.Region(0)
}

func (x *FoodLogItemResponse) GetMeal() Meal {
	if x != nil {
		return x.Meal
	}
	return Meal_MEAL_UNSPECIFIED
}

var File_neutral_diet_user_v1_food_item_log_proto protoreflect.FileDescriptor

var file_neutral_diet_user_v1_food_item_log_proto_rawDesc = []byte{
	0x0a, 0x28, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x5f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x6e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x21, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f, 0x66,
	0x6f, 0x6f, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x02, 0x0a, 0x12, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x67,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x66,
	0x6f, 0x6f, 0x64, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61,
	0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61,
	0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x6e, 0x69, 0x74, 0x52, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x75, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e,
	0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x61, 0x6c, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6c, 0x22,
	0xe8, 0x02, 0x0a, 0x13, 0x46, 0x6f, 0x6f, 0x64, 0x4c, 0x6f, 0x67, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x66, 0x6f, 0x6f, 0x64, 0x5f,
	0x69, 0x74, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66,
	0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x61, 0x72,
	0x62, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x6f, 0x74, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0f, 0x63, 0x61, 0x72, 0x62, 0x6f, 0x6e, 0x46, 0x6f, 0x6f, 0x74, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x04,
	0x75, 0x6e, 0x69, 0x74, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64,
	0x69, 0x65, 0x74, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65,
	0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72,
	0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x61, 0x6c, 0x52, 0x04, 0x6d, 0x65, 0x61, 0x6c, 0x2a, 0x90, 0x01, 0x0a, 0x04, 0x55,
	0x6e, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x10, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x49,
	0x54, 0x5f, 0x47, 0x52, 0x41, 0x4d, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x4f, 0x55, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x50, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x03, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x4c, 0x49, 0x54, 0x52, 0x45, 0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x55, 0x4e, 0x49, 0x54,
	0x5f, 0x4d, 0x49, 0x4c, 0x4c, 0x49, 0x4c, 0x49, 0x54, 0x52, 0x45, 0x10, 0x05, 0x12, 0x0f, 0x0a,
	0x0b, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x47, 0x41, 0x4c, 0x4c, 0x4f, 0x4e, 0x10, 0x06, 0x12, 0x0d,
	0x0a, 0x09, 0x55, 0x4e, 0x49, 0x54, 0x5f, 0x50, 0x49, 0x4e, 0x54, 0x10, 0x07, 0x2a, 0x51, 0x0a,
	0x04, 0x4d, 0x65, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x10, 0x4d, 0x45, 0x41, 0x4c, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x4d,
	0x45, 0x41, 0x4c, 0x5f, 0x42, 0x52, 0x45, 0x41, 0x4b, 0x46, 0x41, 0x53, 0x54, 0x10, 0x01, 0x12,
	0x0e, 0x0a, 0x0a, 0x4d, 0x45, 0x41, 0x4c, 0x5f, 0x4c, 0x55, 0x4e, 0x43, 0x48, 0x10, 0x02, 0x12,
	0x0f, 0x0a, 0x0b, 0x4d, 0x45, 0x41, 0x4c, 0x5f, 0x44, 0x49, 0x4e, 0x4e, 0x45, 0x52, 0x10, 0x03,
	0x42, 0xe7, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x64, 0x69, 0x65, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x10, 0x46,
	0x6f, 0x6f, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x4c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x6d, 0x30, 0x32, 0x7a, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65,
	0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x69,
	0x64, 0x6c, 0x2f, 0x6e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x64, 0x69, 0x65, 0x74, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x4e, 0x55, 0x58, 0xaa, 0x02, 0x13, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69,
	0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x13, 0x4e, 0x65, 0x75,
	0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x1f, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74, 0x5c, 0x55,
	0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x15, 0x4e, 0x65, 0x75, 0x74, 0x72, 0x61, 0x6c, 0x44, 0x69, 0x65, 0x74,
	0x3a, 0x3a, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
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

var file_neutral_diet_user_v1_food_item_log_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_neutral_diet_user_v1_food_item_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_neutral_diet_user_v1_food_item_log_proto_goTypes = []interface{}{
	(Unit)(0),                   // 0: neutral_diet.user.v1.Unit
	(Meal)(0),                   // 1: neutral_diet.user.v1.Meal
	(*FoodLogItemRequest)(nil),  // 2: neutral_diet.user.v1.FoodLogItemRequest
	(*FoodLogItemResponse)(nil), // 3: neutral_diet.user.v1.FoodLogItemResponse
	(*Date)(nil),                // 4: neutral_diet.user.v1.Date
	(v1.Region)(0),              // 5: neutral_diet.food.v1.Region
}
var file_neutral_diet_user_v1_food_item_log_proto_depIdxs = []int32{
	4, // 0: neutral_diet.user.v1.FoodLogItemRequest.date:type_name -> neutral_diet.user.v1.Date
	0, // 1: neutral_diet.user.v1.FoodLogItemRequest.unit:type_name -> neutral_diet.user.v1.Unit
	5, // 2: neutral_diet.user.v1.FoodLogItemRequest.region:type_name -> neutral_diet.food.v1.Region
	1, // 3: neutral_diet.user.v1.FoodLogItemRequest.meal:type_name -> neutral_diet.user.v1.Meal
	4, // 4: neutral_diet.user.v1.FoodLogItemResponse.date:type_name -> neutral_diet.user.v1.Date
	0, // 5: neutral_diet.user.v1.FoodLogItemResponse.unit:type_name -> neutral_diet.user.v1.Unit
	5, // 6: neutral_diet.user.v1.FoodLogItemResponse.region:type_name -> neutral_diet.food.v1.Region
	1, // 7: neutral_diet.user.v1.FoodLogItemResponse.meal:type_name -> neutral_diet.user.v1.Meal
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_neutral_diet_user_v1_food_item_log_proto_init() }
func file_neutral_diet_user_v1_food_item_log_proto_init() {
	if File_neutral_diet_user_v1_food_item_log_proto != nil {
		return
	}
	file_neutral_diet_user_v1_date_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_neutral_diet_user_v1_food_item_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodLogItemRequest); i {
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
		file_neutral_diet_user_v1_food_item_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FoodLogItemResponse); i {
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
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_neutral_diet_user_v1_food_item_log_proto_goTypes,
		DependencyIndexes: file_neutral_diet_user_v1_food_item_log_proto_depIdxs,
		EnumInfos:         file_neutral_diet_user_v1_food_item_log_proto_enumTypes,
		MessageInfos:      file_neutral_diet_user_v1_food_item_log_proto_msgTypes,
	}.Build()
	File_neutral_diet_user_v1_food_item_log_proto = out.File
	file_neutral_diet_user_v1_food_item_log_proto_rawDesc = nil
	file_neutral_diet_user_v1_food_item_log_proto_goTypes = nil
	file_neutral_diet_user_v1_food_item_log_proto_depIdxs = nil
}
