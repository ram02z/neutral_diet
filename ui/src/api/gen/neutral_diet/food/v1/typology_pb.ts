// @generated by protoc-gen-es v0.2.1 with parameter "target=ts"
// @generated from file neutral_diet/food/v1/typology.proto (package neutral_diet.food.v1, syntax proto3)
/* eslint-disable */
/* @ts-nocheck */

import type {BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage} from "@bufbuild/protobuf";
import {Message, proto3} from "@bufbuild/protobuf";

/**
 * @generated from message neutral_diet.food.v1.Typology
 */
export class Typology extends Message<Typology> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: optional string description = 2;
   */
  description?: string;

  /**
   * @generated from field: optional int32 sub_typology_id = 3;
   */
  subTypologyId?: number;

  constructor(data?: PartialMessage<Typology>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime = proto3;
  static readonly typeName = "neutral_diet.food.v1.Typology";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "sub_typology_id", kind: "scalar", T: 5 /* ScalarType.INT32 */, opt: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Typology {
    return new Typology().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Typology {
    return new Typology().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Typology {
    return new Typology().fromJsonString(jsonString, options);
  }

  static equals(a: Typology | PlainMessage<Typology> | undefined, b: Typology | PlainMessage<Typology> | undefined): boolean {
    return proto3.util.equals(Typology, a, b);
  }
}

