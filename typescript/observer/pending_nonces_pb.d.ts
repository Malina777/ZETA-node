// @generated by protoc-gen-es v1.3.0 with parameter "target=dts"
// @generated from file observer/pending_nonces.proto (package zetachain.zetacore.observer, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * store key is tss+chainid
 *
 * @generated from message zetachain.zetacore.observer.PendingNonces
 */
export declare class PendingNonces extends Message<PendingNonces> {
  /**
   * @generated from field: int64 nonce_low = 1;
   */
  nonceLow: bigint;

  /**
   * @generated from field: int64 nonce_high = 2;
   */
  nonceHigh: bigint;

  /**
   * @generated from field: int64 chain_id = 3;
   */
  chainId: bigint;

  /**
   * @generated from field: string tss = 4;
   */
  tss: string;

  constructor(data?: PartialMessage<PendingNonces>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "zetachain.zetacore.observer.PendingNonces";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): PendingNonces;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): PendingNonces;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): PendingNonces;

  static equals(a: PendingNonces | PlainMessage<PendingNonces> | undefined, b: PendingNonces | PlainMessage<PendingNonces> | undefined): boolean;
}

