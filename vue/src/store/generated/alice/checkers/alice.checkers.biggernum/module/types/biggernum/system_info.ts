/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "alice.checkers.biggernum";

export interface SystemInfo {}

const baseSystemInfo: object = {};

export const SystemInfo = {
  encode(_: SystemInfo, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): SystemInfo {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseSystemInfo } as SystemInfo;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    return message;
  },

  toJSON(_: SystemInfo): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<SystemInfo>): SystemInfo {
    const message = { ...baseSystemInfo } as SystemInfo;
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
