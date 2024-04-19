// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file crosschain/rate_limiter_flags.proto (package zetachain.zetacore.crosschain, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message zetachain.zetacore.crosschain.RateLimiterFlags
 */
export declare class RateLimiterFlags extends Message<RateLimiterFlags> {
  /**
   * @generated from field: bool enabled = 1;
   */
  enabled: boolean;

  /**
   * window in blocks
   *
   * @generated from field: int64 window = 2;
   */
  window: bigint;

  /**
   * rate in azeta per block
   *
   * @generated from field: string rate = 3;
   */
  rate: string;

  /**
   * conversion in azeta per token
   *
   * @generated from field: repeated zetachain.zetacore.crosschain.Conversion conversions = 4;
   */
  conversions: Conversion[];

  constructor(data?: PartialMessage<RateLimiterFlags>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.RateLimiterFlags";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RateLimiterFlags;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RateLimiterFlags;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RateLimiterFlags;

  static equals(a: RateLimiterFlags | PlainMessage<RateLimiterFlags> | undefined, b: RateLimiterFlags | PlainMessage<RateLimiterFlags> | undefined): boolean;
}

/**
 * @generated from message zetachain.zetacore.crosschain.Conversion
 */
export declare class Conversion extends Message<Conversion> {
  /**
   * @generated from field: string zrc20 = 1;
   */
  zrc20: string;

  /**
   * @generated from field: string rate = 2;
   */
  rate: string;

  constructor(data?: PartialMessage<Conversion>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.Conversion";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Conversion;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Conversion;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Conversion;

  static equals(a: Conversion | PlainMessage<Conversion> | undefined, b: Conversion | PlainMessage<Conversion> | undefined): boolean;
}

