/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "alice.checkers.oddnum";

export interface StoredGame {
  index: string;
  board: string;
  turn: string;
  black: string;
  red: string;
  wager: number;
}

const baseStoredGame: object = {
  index: "",
  board: "",
  turn: "",
  black: "",
  red: "",
  wager: 0,
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.board !== "") {
      writer.uint32(18).string(message.board);
    }
    if (message.turn !== "") {
      writer.uint32(26).string(message.turn);
    }
    if (message.black !== "") {
      writer.uint32(34).string(message.black);
    }
    if (message.red !== "") {
      writer.uint32(42).string(message.red);
    }
    if (message.wager !== 0) {
      writer.uint32(48).uint64(message.wager);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredGame {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredGame } as StoredGame;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.board = reader.string();
          break;
        case 3:
          message.turn = reader.string();
          break;
        case 4:
          message.black = reader.string();
          break;
        case 5:
          message.red = reader.string();
          break;
        case 6:
          message.wager = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = String(object.board);
    } else {
      message.board = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = String(object.turn);
    } else {
      message.turn = "";
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = String(object.black);
    } else {
      message.black = "";
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = String(object.red);
    } else {
      message.red = "";
    }
    if (object.wager !== undefined && object.wager !== null) {
      message.wager = Number(object.wager);
    } else {
      message.wager = 0;
    }
    return message;
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.board !== undefined && (obj.board = message.board);
    message.turn !== undefined && (obj.turn = message.turn);
    message.black !== undefined && (obj.black = message.black);
    message.red !== undefined && (obj.red = message.red);
    message.wager !== undefined && (obj.wager = message.wager);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredGame>): StoredGame {
    const message = { ...baseStoredGame } as StoredGame;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.board !== undefined && object.board !== null) {
      message.board = object.board;
    } else {
      message.board = "";
    }
    if (object.turn !== undefined && object.turn !== null) {
      message.turn = object.turn;
    } else {
      message.turn = "";
    }
    if (object.black !== undefined && object.black !== null) {
      message.black = object.black;
    } else {
      message.black = "";
    }
    if (object.red !== undefined && object.red !== null) {
      message.red = object.red;
    } else {
      message.red = "";
    }
    if (object.wager !== undefined && object.wager !== null) {
      message.wager = object.wager;
    } else {
      message.wager = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

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

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
