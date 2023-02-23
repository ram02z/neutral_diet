// @generated by protoc-gen-es v0.5.0 with parameter "target=ts"
// @generated from file neutral_diet/food/v1/food_item.proto (package neutral_diet.food.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Source } from "./source_pb.js";

/**
 * @generated from message neutral_diet.food.v1.FoodItem
 */
export class FoodItem extends Message<FoodItem> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: int32 typology_id = 3;
   */
  typologyId = 0;

  /**
   * @generated from field: neutral_diet.food.v1.FoodItem.CfType cf_type = 4;
   */
  cfType = FoodItem_CfType.UNSPECIFIED;

  constructor(data?: PartialMessage<FoodItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.FoodItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "typology_id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 4, name: "cf_type", kind: "enum", T: proto3.getEnumType(FoodItem_CfType) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FoodItem {
    return new FoodItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FoodItem {
    return new FoodItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FoodItem {
    return new FoodItem().fromJsonString(jsonString, options);
  }

  static equals(a: FoodItem | PlainMessage<FoodItem> | undefined, b: FoodItem | PlainMessage<FoodItem> | undefined): boolean {
    return proto3.util.equals(FoodItem, a, b);
  }
}

/**
 * @generated from enum neutral_diet.food.v1.FoodItem.CfType
 */
export enum FoodItem_CfType {
  /**
   * @generated from enum value: CF_TYPE_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: CF_TYPE_TYPOLOGY = 1;
   */
  TYPOLOGY = 1,

  /**
   * @generated from enum value: CF_TYPE_SUB_TYPOLOGY = 2;
   */
  SUB_TYPOLOGY = 2,

  /**
   * @generated from enum value: CF_TYPE_ITEM = 3;
   */
  ITEM = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(FoodItem_CfType)
proto3.util.setEnumType(FoodItem_CfType, "neutral_diet.food.v1.FoodItem.CfType", [
  { no: 0, name: "CF_TYPE_UNSPECIFIED" },
  { no: 1, name: "CF_TYPE_TYPOLOGY" },
  { no: 2, name: "CF_TYPE_SUB_TYPOLOGY" },
  { no: 3, name: "CF_TYPE_ITEM" },
]);

/**
 * @generated from message neutral_diet.food.v1.FoodItemInfo
 */
export class FoodItemInfo extends Message<FoodItemInfo> {
  /**
   * @generated from field: string typology_name = 1;
   */
  typologyName = "";

  /**
   * @generated from field: string sub_typology_name = 2;
   */
  subTypologyName = "";

  /**
   * @generated from field: int64 no_sources = 3;
   */
  noSources = protoInt64.zero;

  /**
   * @generated from field: repeated neutral_diet.food.v1.Source sources = 4;
   */
  sources: Source[] = [];

  constructor(data?: PartialMessage<FoodItemInfo>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.FoodItemInfo";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "typology_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "sub_typology_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "no_sources", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "sources", kind: "message", T: Source, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FoodItemInfo {
    return new FoodItemInfo().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FoodItemInfo {
    return new FoodItemInfo().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FoodItemInfo {
    return new FoodItemInfo().fromJsonString(jsonString, options);
  }

  static equals(a: FoodItemInfo | PlainMessage<FoodItemInfo> | undefined, b: FoodItemInfo | PlainMessage<FoodItemInfo> | undefined): boolean {
    return proto3.util.equals(FoodItemInfo, a, b);
  }
}

/**
 * @generated from message neutral_diet.food.v1.AggregateFoodItem
 */
export class AggregateFoodItem extends Message<AggregateFoodItem> {
  /**
   * @generated from field: int32 id = 1;
   */
  id = 0;

  /**
   * @generated from field: string food_name = 2;
   */
  foodName = "";

  /**
   * @generated from field: double median_carbon_footprint = 3;
   */
  medianCarbonFootprint = 0;

  constructor(data?: PartialMessage<AggregateFoodItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.AggregateFoodItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "food_name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "median_carbon_footprint", kind: "scalar", T: 1 /* ScalarType.DOUBLE */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AggregateFoodItem {
    return new AggregateFoodItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AggregateFoodItem {
    return new AggregateFoodItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AggregateFoodItem {
    return new AggregateFoodItem().fromJsonString(jsonString, options);
  }

  static equals(a: AggregateFoodItem | PlainMessage<AggregateFoodItem> | undefined, b: AggregateFoodItem | PlainMessage<AggregateFoodItem> | undefined): boolean {
    return proto3.util.equals(AggregateFoodItem, a, b);
  }
}

