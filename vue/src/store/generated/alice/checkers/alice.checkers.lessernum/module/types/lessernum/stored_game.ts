/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "alice.checkers.lessernum";

export interface StoredGame {
  index: string;
  player1: string;
  player2: string;
  playerToMove: number;
  move1: number;
  move2: number;
  winner: string;
  wager: number;
}

const baseStoredGame: object = {
  index: "",
  player1: "",
  player2: "",
  playerToMove: 0,
  move1: 0,
  move2: 0,
  winner: "",
  wager: 0,
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.player1 !== "") {
      writer.uint32(18).string(message.player1);
    }
    if (message.player2 !== "") {
      writer.uint32(26).string(message.player2);
    }
    if (message.playerToMove !== 0) {
      writer.uint32(32).uint64(message.playerToMove);
    }
    if (message.move1 !== 0) {
      writer.uint32(40).uint64(message.move1);
    }
    if (message.move2 !== 0) {
      writer.uint32(48).uint64(message.move2);
    }
    if (message.winner !== "") {
      writer.uint32(58).string(message.winner);
    }
    if (message.wager !== 0) {
      writer.uint32(64).uint64(message.wager);
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
          message.player1 = reader.string();
          break;
        case 3:
          message.player2 = reader.string();
          break;
        case 4:
          message.playerToMove = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.move1 = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.move2 = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.winner = reader.string();
          break;
        case 8:
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
    if (object.player1 !== undefined && object.player1 !== null) {
      message.player1 = String(object.player1);
    } else {
      message.player1 = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = String(object.player2);
    } else {
      message.player2 = "";
    }
    if (object.playerToMove !== undefined && object.playerToMove !== null) {
      message.playerToMove = Number(object.playerToMove);
    } else {
      message.playerToMove = 0;
    }
    if (object.move1 !== undefined && object.move1 !== null) {
      message.move1 = Number(object.move1);
    } else {
      message.move1 = 0;
    }
    if (object.move2 !== undefined && object.move2 !== null) {
      message.move2 = Number(object.move2);
    } else {
      message.move2 = 0;
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = String(object.winner);
    } else {
      message.winner = "";
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
    message.player1 !== undefined && (obj.player1 = message.player1);
    message.player2 !== undefined && (obj.player2 = message.player2);
    message.playerToMove !== undefined &&
      (obj.playerToMove = message.playerToMove);
    message.move1 !== undefined && (obj.move1 = message.move1);
    message.move2 !== undefined && (obj.move2 = message.move2);
    message.winner !== undefined && (obj.winner = message.winner);
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
    if (object.player1 !== undefined && object.player1 !== null) {
      message.player1 = object.player1;
    } else {
      message.player1 = "";
    }
    if (object.player2 !== undefined && object.player2 !== null) {
      message.player2 = object.player2;
    } else {
      message.player2 = "";
    }
    if (object.playerToMove !== undefined && object.playerToMove !== null) {
      message.playerToMove = object.playerToMove;
    } else {
      message.playerToMove = 0;
    }
    if (object.move1 !== undefined && object.move1 !== null) {
      message.move1 = object.move1;
    } else {
      message.move1 = 0;
    }
    if (object.move2 !== undefined && object.move2 !== null) {
      message.move2 = object.move2;
    } else {
      message.move2 = 0;
    }
    if (object.winner !== undefined && object.winner !== null) {
      message.winner = object.winner;
    } else {
      message.winner = "";
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
