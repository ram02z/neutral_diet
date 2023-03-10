// @generated by protoc-gen-es v0.5.0 with parameter "target=ts"
// @generated from file neutral_diet/user/v1/api.proto (package neutral_diet.user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { FoodLogItemRequest, FoodLogItemResponse, Meal, Unit } from "./food_item_log_pb.js";
import { Region } from "../../food/v1/region_pb.js";
import { Date } from "./date_pb.js";
import { UserSettings } from "./user_pb.js";

/**
 * @generated from message neutral_diet.user.v1.AddFoodItemRequest
 */
export class AddFoodItemRequest extends Message<AddFoodItemRequest> {
  /**
   * @generated from field: neutral_diet.user.v1.FoodLogItemRequest food_log_item = 1;
   */
  foodLogItem?: FoodLogItemRequest;

  constructor(data?: PartialMessage<AddFoodItemRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.AddFoodItemRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "food_log_item", kind: "message", T: FoodLogItemRequest },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddFoodItemRequest {
    return new AddFoodItemRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddFoodItemRequest {
    return new AddFoodItemRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddFoodItemRequest {
    return new AddFoodItemRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddFoodItemRequest | PlainMessage<AddFoodItemRequest> | undefined, b: AddFoodItemRequest | PlainMessage<AddFoodItemRequest> | undefined): boolean {
    return proto3.util.equals(AddFoodItemRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.AddFoodItemResponse
 */
export class AddFoodItemResponse extends Message<AddFoodItemResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: double carbon_footprint = 2;
   */
  carbonFootprint = 0;

  constructor(data?: PartialMessage<AddFoodItemResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.AddFoodItemResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "carbon_footprint", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddFoodItemResponse {
    return new AddFoodItemResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddFoodItemResponse {
    return new AddFoodItemResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddFoodItemResponse {
    return new AddFoodItemResponse().fromJsonString(jsonString, options);
  }

  static equals(a: AddFoodItemResponse | PlainMessage<AddFoodItemResponse> | undefined, b: AddFoodItemResponse | PlainMessage<AddFoodItemResponse> | undefined): boolean {
    return proto3.util.equals(AddFoodItemResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.UpdateFoodItemRequest
 */
export class UpdateFoodItemRequest extends Message<UpdateFoodItemRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: double quantity = 2;
   */
  quantity = 0;

  /**
   * @generated from field: neutral_diet.user.v1.Unit unit = 3;
   */
  unit = Unit.UNSPECIFIED;

  /**
   * @generated from field: neutral_diet.food.v1.Region region = 4;
   */
  region = Region.UNSPECIFIED;

  /**
   * @generated from field: neutral_diet.user.v1.Meal meal = 5;
   */
  meal = Meal.UNSPECIFIED;

  constructor(data?: PartialMessage<UpdateFoodItemRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.UpdateFoodItemRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "quantity", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 3, name: "unit", kind: "enum", T: proto3.getEnumType(Unit) },
    { no: 4, name: "region", kind: "enum", T: proto3.getEnumType(Region) },
    { no: 5, name: "meal", kind: "enum", T: proto3.getEnumType(Meal) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateFoodItemRequest {
    return new UpdateFoodItemRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateFoodItemRequest {
    return new UpdateFoodItemRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateFoodItemRequest {
    return new UpdateFoodItemRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateFoodItemRequest | PlainMessage<UpdateFoodItemRequest> | undefined, b: UpdateFoodItemRequest | PlainMessage<UpdateFoodItemRequest> | undefined): boolean {
    return proto3.util.equals(UpdateFoodItemRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.UpdateFoodItemResponse
 */
export class UpdateFoodItemResponse extends Message<UpdateFoodItemResponse> {
  /**
   * @generated from field: double carbon_footprint = 1;
   */
  carbonFootprint = 0;

  constructor(data?: PartialMessage<UpdateFoodItemResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.UpdateFoodItemResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "carbon_footprint", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateFoodItemResponse {
    return new UpdateFoodItemResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateFoodItemResponse {
    return new UpdateFoodItemResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateFoodItemResponse {
    return new UpdateFoodItemResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateFoodItemResponse | PlainMessage<UpdateFoodItemResponse> | undefined, b: UpdateFoodItemResponse | PlainMessage<UpdateFoodItemResponse> | undefined): boolean {
    return proto3.util.equals(UpdateFoodItemResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.DeleteFoodItemRequest
 */
export class DeleteFoodItemRequest extends Message<DeleteFoodItemRequest> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<DeleteFoodItemRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.DeleteFoodItemRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteFoodItemRequest {
    return new DeleteFoodItemRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteFoodItemRequest {
    return new DeleteFoodItemRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteFoodItemRequest {
    return new DeleteFoodItemRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteFoodItemRequest | PlainMessage<DeleteFoodItemRequest> | undefined, b: DeleteFoodItemRequest | PlainMessage<DeleteFoodItemRequest> | undefined): boolean {
    return proto3.util.equals(DeleteFoodItemRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.DeleteFoodItemResponse
 */
export class DeleteFoodItemResponse extends Message<DeleteFoodItemResponse> {
  constructor(data?: PartialMessage<DeleteFoodItemResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.DeleteFoodItemResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteFoodItemResponse {
    return new DeleteFoodItemResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteFoodItemResponse {
    return new DeleteFoodItemResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteFoodItemResponse {
    return new DeleteFoodItemResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteFoodItemResponse | PlainMessage<DeleteFoodItemResponse> | undefined, b: DeleteFoodItemResponse | PlainMessage<DeleteFoodItemResponse> | undefined): boolean {
    return proto3.util.equals(DeleteFoodItemResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetFoodItemLogRequest
 */
export class GetFoodItemLogRequest extends Message<GetFoodItemLogRequest> {
  /**
   * @generated from field: neutral_diet.user.v1.Date date = 1;
   */
  date?: Date;

  constructor(data?: PartialMessage<GetFoodItemLogRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetFoodItemLogRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "date", kind: "message", T: Date },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetFoodItemLogRequest {
    return new GetFoodItemLogRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetFoodItemLogRequest {
    return new GetFoodItemLogRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetFoodItemLogRequest {
    return new GetFoodItemLogRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetFoodItemLogRequest | PlainMessage<GetFoodItemLogRequest> | undefined, b: GetFoodItemLogRequest | PlainMessage<GetFoodItemLogRequest> | undefined): boolean {
    return proto3.util.equals(GetFoodItemLogRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetFoodItemLogResponse
 */
export class GetFoodItemLogResponse extends Message<GetFoodItemLogResponse> {
  /**
   * @generated from field: repeated neutral_diet.user.v1.FoodLogItemResponse food_item_log = 1;
   */
  foodItemLog: FoodLogItemResponse[] = [];

  constructor(data?: PartialMessage<GetFoodItemLogResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetFoodItemLogResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "food_item_log", kind: "message", T: FoodLogItemResponse, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetFoodItemLogResponse {
    return new GetFoodItemLogResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetFoodItemLogResponse {
    return new GetFoodItemLogResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetFoodItemLogResponse {
    return new GetFoodItemLogResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetFoodItemLogResponse | PlainMessage<GetFoodItemLogResponse> | undefined, b: GetFoodItemLogResponse | PlainMessage<GetFoodItemLogResponse> | undefined): boolean {
    return proto3.util.equals(GetFoodItemLogResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetFoodItemLogDaysRequest
 */
export class GetFoodItemLogDaysRequest extends Message<GetFoodItemLogDaysRequest> {
  /**
   * @generated from field: int32 month = 1;
   */
  month = 0;

  /**
   * @generated from field: int32 year = 2;
   */
  year = 0;

  constructor(data?: PartialMessage<GetFoodItemLogDaysRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetFoodItemLogDaysRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "month", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "year", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetFoodItemLogDaysRequest {
    return new GetFoodItemLogDaysRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetFoodItemLogDaysRequest {
    return new GetFoodItemLogDaysRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetFoodItemLogDaysRequest {
    return new GetFoodItemLogDaysRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetFoodItemLogDaysRequest | PlainMessage<GetFoodItemLogDaysRequest> | undefined, b: GetFoodItemLogDaysRequest | PlainMessage<GetFoodItemLogDaysRequest> | undefined): boolean {
    return proto3.util.equals(GetFoodItemLogDaysRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetFoodItemLogDaysResponse
 */
export class GetFoodItemLogDaysResponse extends Message<GetFoodItemLogDaysResponse> {
  /**
   * @generated from field: repeated int32 days = 1;
   */
  days: number[] = [];

  constructor(data?: PartialMessage<GetFoodItemLogDaysResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetFoodItemLogDaysResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "days", kind: "scalar", T: 5 /* ScalarType.INT32 */, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetFoodItemLogDaysResponse {
    return new GetFoodItemLogDaysResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetFoodItemLogDaysResponse {
    return new GetFoodItemLogDaysResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetFoodItemLogDaysResponse {
    return new GetFoodItemLogDaysResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetFoodItemLogDaysResponse | PlainMessage<GetFoodItemLogDaysResponse> | undefined, b: GetFoodItemLogDaysResponse | PlainMessage<GetFoodItemLogDaysResponse> | undefined): boolean {
    return proto3.util.equals(GetFoodItemLogDaysResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.CreateUserRequest
 */
export class CreateUserRequest extends Message<CreateUserRequest> {
  /**
   * @generated from field: string firebase_uid = 1;
   */
  firebaseUid = "";

  constructor(data?: PartialMessage<CreateUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.CreateUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "firebase_uid", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined, b: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined): boolean {
    return proto3.util.equals(CreateUserRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.CreateUserResponse
 */
export class CreateUserResponse extends Message<CreateUserResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<CreateUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.CreateUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserResponse {
    return new CreateUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined, b: CreateUserResponse | PlainMessage<CreateUserResponse> | undefined): boolean {
    return proto3.util.equals(CreateUserResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.DeleteUserRequest
 */
export class DeleteUserRequest extends Message<DeleteUserRequest> {
  constructor(data?: PartialMessage<DeleteUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.DeleteUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserRequest {
    return new DeleteUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserRequest {
    return new DeleteUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserRequest {
    return new DeleteUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteUserRequest | PlainMessage<DeleteUserRequest> | undefined, b: DeleteUserRequest | PlainMessage<DeleteUserRequest> | undefined): boolean {
    return proto3.util.equals(DeleteUserRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.DeleteUserResponse
 */
export class DeleteUserResponse extends Message<DeleteUserResponse> {
  constructor(data?: PartialMessage<DeleteUserResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.DeleteUserResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): DeleteUserResponse {
    return new DeleteUserResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): DeleteUserResponse {
    return new DeleteUserResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): DeleteUserResponse {
    return new DeleteUserResponse().fromJsonString(jsonString, options);
  }

  static equals(a: DeleteUserResponse | PlainMessage<DeleteUserResponse> | undefined, b: DeleteUserResponse | PlainMessage<DeleteUserResponse> | undefined): boolean {
    return proto3.util.equals(DeleteUserResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.UpdateUserSettingsRequest
 */
export class UpdateUserSettingsRequest extends Message<UpdateUserSettingsRequest> {
  /**
   * @generated from field: neutral_diet.user.v1.UserSettings user_settings = 1;
   */
  userSettings?: UserSettings;

  constructor(data?: PartialMessage<UpdateUserSettingsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.UpdateUserSettingsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_settings", kind: "message", T: UserSettings },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserSettingsRequest {
    return new UpdateUserSettingsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserSettingsRequest {
    return new UpdateUserSettingsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserSettingsRequest {
    return new UpdateUserSettingsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserSettingsRequest | PlainMessage<UpdateUserSettingsRequest> | undefined, b: UpdateUserSettingsRequest | PlainMessage<UpdateUserSettingsRequest> | undefined): boolean {
    return proto3.util.equals(UpdateUserSettingsRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.UpdateUserSettingsResponse
 */
export class UpdateUserSettingsResponse extends Message<UpdateUserSettingsResponse> {
  constructor(data?: PartialMessage<UpdateUserSettingsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.UpdateUserSettingsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateUserSettingsResponse {
    return new UpdateUserSettingsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateUserSettingsResponse {
    return new UpdateUserSettingsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateUserSettingsResponse {
    return new UpdateUserSettingsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateUserSettingsResponse | PlainMessage<UpdateUserSettingsResponse> | undefined, b: UpdateUserSettingsResponse | PlainMessage<UpdateUserSettingsResponse> | undefined): boolean {
    return proto3.util.equals(UpdateUserSettingsResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetUserSettingsRequest
 */
export class GetUserSettingsRequest extends Message<GetUserSettingsRequest> {
  constructor(data?: PartialMessage<GetUserSettingsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetUserSettingsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserSettingsRequest {
    return new GetUserSettingsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserSettingsRequest {
    return new GetUserSettingsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserSettingsRequest {
    return new GetUserSettingsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserSettingsRequest | PlainMessage<GetUserSettingsRequest> | undefined, b: GetUserSettingsRequest | PlainMessage<GetUserSettingsRequest> | undefined): boolean {
    return proto3.util.equals(GetUserSettingsRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetUserSettingsResponse
 */
export class GetUserSettingsResponse extends Message<GetUserSettingsResponse> {
  /**
   * @generated from field: neutral_diet.user.v1.UserSettings user_settings = 1;
   */
  userSettings?: UserSettings;

  constructor(data?: PartialMessage<GetUserSettingsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetUserSettingsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_settings", kind: "message", T: UserSettings },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserSettingsResponse {
    return new GetUserSettingsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserSettingsResponse {
    return new GetUserSettingsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserSettingsResponse {
    return new GetUserSettingsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserSettingsResponse | PlainMessage<GetUserSettingsResponse> | undefined, b: GetUserSettingsResponse | PlainMessage<GetUserSettingsResponse> | undefined): boolean {
    return proto3.util.equals(GetUserSettingsResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetUserInsightsRequest
 */
export class GetUserInsightsRequest extends Message<GetUserInsightsRequest> {
  constructor(data?: PartialMessage<GetUserInsightsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetUserInsightsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserInsightsRequest {
    return new GetUserInsightsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserInsightsRequest {
    return new GetUserInsightsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserInsightsRequest {
    return new GetUserInsightsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserInsightsRequest | PlainMessage<GetUserInsightsRequest> | undefined, b: GetUserInsightsRequest | PlainMessage<GetUserInsightsRequest> | undefined): boolean {
    return proto3.util.equals(GetUserInsightsRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.user.v1.GetUserInsightsResponse
 */
export class GetUserInsightsResponse extends Message<GetUserInsightsResponse> {
  /**
   * @generated from field: double overall_carbon_footprint = 1;
   */
  overallCarbonFootprint = 0;

  /**
   * @generated from field: int32 no_entries = 2;
   */
  noEntries = 0;

  /**
   * @generated from field: double daily_average_carbon_footprint_dietary_requirement = 3;
   */
  dailyAverageCarbonFootprintDietaryRequirement = 0;

  /**
   * @generated from field: double daily_average_carbon_footprint_overall = 4;
   */
  dailyAverageCarbonFootprintOverall = 0;

  /**
   * @generated from field: int32 streak_len = 5;
   */
  streakLen = 0;

  /**
   * @generated from field: bool is_streak_active = 6;
   */
  isStreakActive = false;

  constructor(data?: PartialMessage<GetUserInsightsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.GetUserInsightsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "overall_carbon_footprint", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 2, name: "no_entries", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 3, name: "daily_average_carbon_footprint_dietary_requirement", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 4, name: "daily_average_carbon_footprint_overall", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
    { no: 5, name: "streak_len", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 6, name: "is_streak_active", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetUserInsightsResponse {
    return new GetUserInsightsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetUserInsightsResponse {
    return new GetUserInsightsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetUserInsightsResponse {
    return new GetUserInsightsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: GetUserInsightsResponse | PlainMessage<GetUserInsightsResponse> | undefined, b: GetUserInsightsResponse | PlainMessage<GetUserInsightsResponse> | undefined): boolean {
    return proto3.util.equals(GetUserInsightsResponse, a, b);
  }
}

