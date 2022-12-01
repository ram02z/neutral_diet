// @generated by protoc-gen-es v0.2.1 with parameter "target=ts"
// @generated from file neutral_diet/food/v1/food_item.proto (package neutral_diet.food.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

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

