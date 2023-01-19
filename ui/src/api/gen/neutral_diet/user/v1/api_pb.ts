// @generated by protoc-gen-es v0.5.0 with parameter "target=ts"
// @generated from file neutral_diet/user/v1/api.proto (package neutral_diet.user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { FoodLogItem } from "./food_item_log_pb.js";

/**
 * @generated from message neutral_diet.user.v1.AddFoodItemRequest
 */
export class AddFoodItemRequest extends Message<AddFoodItemRequest> {
  /**
   * @generated from field: neutral_diet.user.v1.FoodLogItem food_log_item = 2;
   */
  foodLogItem?: FoodLogItem;

  constructor(data?: PartialMessage<AddFoodItemRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.AddFoodItemRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "food_log_item", kind: "message", T: FoodLogItem },
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

  constructor(data?: PartialMessage<AddFoodItemResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.user.v1.AddFoodItemResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
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

