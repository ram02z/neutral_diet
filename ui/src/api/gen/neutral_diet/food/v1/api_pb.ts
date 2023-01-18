// @generated by protoc-gen-es v0.5.0 with parameter "target=ts"
// @generated from file neutral_diet/food/v1/api.proto (package neutral_diet.food.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { LifeCycle } from "./life_cycle_pb.js";
import { AggregateFoodItem, FoodItem } from "./food_item_pb.js";
import { Typology } from "./typology_pb.js";
import { SubTypology } from "./sub_typology_pb.js";
import { Source } from "./source_pb.js";
import { Region } from "./region_pb.js";

/**
 * @generated from message neutral_diet.food.v1.CreateLifeCycleRequest
 */
export class CreateLifeCycleRequest extends Message<CreateLifeCycleRequest> {
  /**
   * @generated from field: neutral_diet.food.v1.LifeCycle life_cycle = 1;
   */
  lifeCycle?: LifeCycle;

  constructor(data?: PartialMessage<CreateLifeCycleRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateLifeCycleRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "life_cycle", kind: "message", T: LifeCycle },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateLifeCycleRequest {
    return new CreateLifeCycleRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateLifeCycleRequest {
    return new CreateLifeCycleRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateLifeCycleRequest {
    return new CreateLifeCycleRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateLifeCycleRequest | PlainMessage<CreateLifeCycleRequest> | undefined, b: CreateLifeCycleRequest | PlainMessage<CreateLifeCycleRequest> | undefined): boolean {
    return proto3.util.equals(CreateLifeCycleRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateLifeCycleResponse
 */
export class CreateLifeCycleResponse extends Message<CreateLifeCycleResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<CreateLifeCycleResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateLifeCycleResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateLifeCycleResponse {
    return new CreateLifeCycleResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateLifeCycleResponse {
    return new CreateLifeCycleResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateLifeCycleResponse {
    return new CreateLifeCycleResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateLifeCycleResponse | PlainMessage<CreateLifeCycleResponse> | undefined, b: CreateLifeCycleResponse | PlainMessage<CreateLifeCycleResponse> | undefined): boolean {
    return proto3.util.equals(CreateLifeCycleResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateFoodItemRequest
 */
export class CreateFoodItemRequest extends Message<CreateFoodItemRequest> {
  /**
   * @generated from field: neutral_diet.food.v1.FoodItem food_item = 1;
   */
  foodItem?: FoodItem;

  constructor(data?: PartialMessage<CreateFoodItemRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateFoodItemRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "food_item", kind: "message", T: FoodItem },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateFoodItemRequest {
    return new CreateFoodItemRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateFoodItemRequest {
    return new CreateFoodItemRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateFoodItemRequest {
    return new CreateFoodItemRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateFoodItemRequest | PlainMessage<CreateFoodItemRequest> | undefined, b: CreateFoodItemRequest | PlainMessage<CreateFoodItemRequest> | undefined): boolean {
    return proto3.util.equals(CreateFoodItemRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateFoodItemResponse
 */
export class CreateFoodItemResponse extends Message<CreateFoodItemResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<CreateFoodItemResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateFoodItemResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateFoodItemResponse {
    return new CreateFoodItemResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateFoodItemResponse {
    return new CreateFoodItemResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateFoodItemResponse {
    return new CreateFoodItemResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateFoodItemResponse | PlainMessage<CreateFoodItemResponse> | undefined, b: CreateFoodItemResponse | PlainMessage<CreateFoodItemResponse> | undefined): boolean {
    return proto3.util.equals(CreateFoodItemResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateTypologyRequest
 */
export class CreateTypologyRequest extends Message<CreateTypologyRequest> {
  /**
   * @generated from field: neutral_diet.food.v1.Typology typology = 1;
   */
  typology?: Typology;

  constructor(data?: PartialMessage<CreateTypologyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateTypologyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "typology", kind: "message", T: Typology },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTypologyRequest {
    return new CreateTypologyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTypologyRequest {
    return new CreateTypologyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTypologyRequest {
    return new CreateTypologyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTypologyRequest | PlainMessage<CreateTypologyRequest> | undefined, b: CreateTypologyRequest | PlainMessage<CreateTypologyRequest> | undefined): boolean {
    return proto3.util.equals(CreateTypologyRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateTypologyResponse
 */
export class CreateTypologyResponse extends Message<CreateTypologyResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<CreateTypologyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateTypologyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateTypologyResponse {
    return new CreateTypologyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateTypologyResponse {
    return new CreateTypologyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateTypologyResponse {
    return new CreateTypologyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateTypologyResponse | PlainMessage<CreateTypologyResponse> | undefined, b: CreateTypologyResponse | PlainMessage<CreateTypologyResponse> | undefined): boolean {
    return proto3.util.equals(CreateTypologyResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateSubTypologyRequest
 */
export class CreateSubTypologyRequest extends Message<CreateSubTypologyRequest> {
  /**
   * @generated from field: neutral_diet.food.v1.SubTypology sub_typology = 1;
   */
  subTypology?: SubTypology;

  constructor(data?: PartialMessage<CreateSubTypologyRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateSubTypologyRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "sub_typology", kind: "message", T: SubTypology },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSubTypologyRequest {
    return new CreateSubTypologyRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSubTypologyRequest {
    return new CreateSubTypologyRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSubTypologyRequest {
    return new CreateSubTypologyRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSubTypologyRequest | PlainMessage<CreateSubTypologyRequest> | undefined, b: CreateSubTypologyRequest | PlainMessage<CreateSubTypologyRequest> | undefined): boolean {
    return proto3.util.equals(CreateSubTypologyRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateSubTypologyResponse
 */
export class CreateSubTypologyResponse extends Message<CreateSubTypologyResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<CreateSubTypologyResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateSubTypologyResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSubTypologyResponse {
    return new CreateSubTypologyResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSubTypologyResponse {
    return new CreateSubTypologyResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSubTypologyResponse {
    return new CreateSubTypologyResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSubTypologyResponse | PlainMessage<CreateSubTypologyResponse> | undefined, b: CreateSubTypologyResponse | PlainMessage<CreateSubTypologyResponse> | undefined): boolean {
    return proto3.util.equals(CreateSubTypologyResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateSourceRequest
 */
export class CreateSourceRequest extends Message<CreateSourceRequest> {
  /**
   * @generated from field: neutral_diet.food.v1.Source source = 1;
   */
  source?: Source;

  constructor(data?: PartialMessage<CreateSourceRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateSourceRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "source", kind: "message", T: Source },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSourceRequest {
    return new CreateSourceRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSourceRequest {
    return new CreateSourceRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSourceRequest {
    return new CreateSourceRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSourceRequest | PlainMessage<CreateSourceRequest> | undefined, b: CreateSourceRequest | PlainMessage<CreateSourceRequest> | undefined): boolean {
    return proto3.util.equals(CreateSourceRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateSourceResponse
 */
export class CreateSourceResponse extends Message<CreateSourceResponse> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  constructor(data?: PartialMessage<CreateSourceResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateSourceResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateSourceResponse {
    return new CreateSourceResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateSourceResponse {
    return new CreateSourceResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateSourceResponse {
    return new CreateSourceResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateSourceResponse | PlainMessage<CreateSourceResponse> | undefined, b: CreateSourceResponse | PlainMessage<CreateSourceResponse> | undefined): boolean {
    return proto3.util.equals(CreateSourceResponse, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateRegionRequest
 */
export class CreateRegionRequest extends Message<CreateRegionRequest> {
  /**
   * @generated from field: neutral_diet.food.v1.Region region = 1;
   */
  region?: Region;

  constructor(data?: PartialMessage<CreateRegionRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateRegionRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "region", kind: "message", T: Region },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRegionRequest {
    return new CreateRegionRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRegionRequest {
    return new CreateRegionRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRegionRequest {
    return new CreateRegionRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRegionRequest | PlainMessage<CreateRegionRequest> | undefined, b: CreateRegionRequest | PlainMessage<CreateRegionRequest> | undefined): boolean {
    return proto3.util.equals(CreateRegionRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.CreateRegionResponse
 */
export class CreateRegionResponse extends Message<CreateRegionResponse> {
  constructor(data?: PartialMessage<CreateRegionResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.CreateRegionResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateRegionResponse {
    return new CreateRegionResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateRegionResponse {
    return new CreateRegionResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateRegionResponse {
    return new CreateRegionResponse().fromJsonString(jsonString, options);
  }

  static equals(a: CreateRegionResponse | PlainMessage<CreateRegionResponse> | undefined, b: CreateRegionResponse | PlainMessage<CreateRegionResponse> | undefined): boolean {
    return proto3.util.equals(CreateRegionResponse, a, b);
  }
}

/**
 * TODO: allow filtering by region 
 *
 * @generated from message neutral_diet.food.v1.ListAggregateFoodItemsRequest
 */
export class ListAggregateFoodItemsRequest extends Message<ListAggregateFoodItemsRequest> {
  constructor(data?: PartialMessage<ListAggregateFoodItemsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.ListAggregateFoodItemsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAggregateFoodItemsRequest {
    return new ListAggregateFoodItemsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAggregateFoodItemsRequest {
    return new ListAggregateFoodItemsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAggregateFoodItemsRequest {
    return new ListAggregateFoodItemsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAggregateFoodItemsRequest | PlainMessage<ListAggregateFoodItemsRequest> | undefined, b: ListAggregateFoodItemsRequest | PlainMessage<ListAggregateFoodItemsRequest> | undefined): boolean {
    return proto3.util.equals(ListAggregateFoodItemsRequest, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.ListAggregateFoodItemsResponse
 */
export class ListAggregateFoodItemsResponse extends Message<ListAggregateFoodItemsResponse> {
  /**
   * @generated from field: repeated neutral_diet.food.v1.AggregateFoodItem food_items = 1;
   */
  foodItems: AggregateFoodItem[] = [];

  constructor(data?: PartialMessage<ListAggregateFoodItemsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.ListAggregateFoodItemsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "food_items", kind: "message", T: AggregateFoodItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAggregateFoodItemsResponse {
    return new ListAggregateFoodItemsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAggregateFoodItemsResponse {
    return new ListAggregateFoodItemsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAggregateFoodItemsResponse {
    return new ListAggregateFoodItemsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAggregateFoodItemsResponse | PlainMessage<ListAggregateFoodItemsResponse> | undefined, b: ListAggregateFoodItemsResponse | PlainMessage<ListAggregateFoodItemsResponse> | undefined): boolean {
    return proto3.util.equals(ListAggregateFoodItemsResponse, a, b);
  }
}

