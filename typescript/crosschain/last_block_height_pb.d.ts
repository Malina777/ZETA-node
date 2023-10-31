// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file crosschain/last_block_height.proto (package zetachain.zetacore.crosschain, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message zetachain.zetacore.crosschain.LastBlockHeight
 */
export declare class LastBlockHeight extends Message<LastBlockHeight> {
  /**
   * @generated from field: string creator = 1;
   */
  creator: string;

  /**
   * @generated from field: string index = 2;
   */
  index: string;

  /**
   * @generated from field: string chain = 3;
   */
  chain: string;

  /**
   * @generated from field: uint64 lastSendHeight = 4;
   */
  lastSendHeight: bigint;

  /**
   * @generated from field: uint64 lastReceiveHeight = 5;
   */
  lastReceiveHeight: bigint;

  constructor(data?: PartialMessage<LastBlockHeight>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.crosschain.LastBlockHeight";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): LastBlockHeight;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): LastBlockHeight;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): LastBlockHeight;

  static equals(a: LastBlockHeight | PlainMessage<LastBlockHeight> | undefined, b: LastBlockHeight | PlainMessage<LastBlockHeight> | undefined): boolean;
}

