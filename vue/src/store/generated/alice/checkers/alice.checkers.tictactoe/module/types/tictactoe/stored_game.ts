/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "alice.checkers.tictactoe";

export interface StoredGame {
  index: string;
  board: string;
  nextTurn: string;
  xPlayer: string;
  yPlayer: string;
}

const baseStoredGame: object = {
  index: "",
  board: "",
  nextTurn: "",
  xPlayer: "",
  yPlayer: "",
};

export const StoredGame = {
  encode(message: StoredGame, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.board !== "") {
      writer.uint32(18).string(message.board);
    }
    if (message.nextTurn !== "") {
      writer.uint32(26).string(message.nextTurn);
    }
    if (message.xPlayer !== "") {
      writer.uint32(34).string(message.xPlayer);
    }
    if (message.yPlayer !== "") {
      writer.uint32(42).string(message.yPlayer);
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
          message.nextTurn = reader.string();
          break;
        case 4:
          message.xPlayer = reader.string();
          break;
        case 5:
          message.yPlayer = reader.string();
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
    if (object.nextTurn !== undefined && object.nextTurn !== null) {
      message.nextTurn = String(object.nextTurn);
    } else {
      message.nextTurn = "";
    }
    if (object.xPlayer !== undefined && object.xPlayer !== null) {
      message.xPlayer = String(object.xPlayer);
    } else {
      message.xPlayer = "";
    }
    if (object.yPlayer !== undefined && object.yPlayer !== null) {
      message.yPlayer = String(object.yPlayer);
    } else {
      message.yPlayer = "";
    }
    return message;
  },

  toJSON(message: StoredGame): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.board !== undefined && (obj.board = message.board);
    message.nextTurn !== undefined && (obj.nextTurn = message.nextTurn);
    message.xPlayer !== undefined && (obj.xPlayer = message.xPlayer);
    message.yPlayer !== undefined && (obj.yPlayer = message.yPlayer);
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
    if (object.nextTurn !== undefined && object.nextTurn !== null) {
      message.nextTurn = object.nextTurn;
    } else {
      message.nextTurn = "";
    }
    if (object.xPlayer !== undefined && object.xPlayer !== null) {
      message.xPlayer = object.xPlayer;
    } else {
      message.xPlayer = "";
    }
    if (object.yPlayer !== undefined && object.yPlayer !== null) {
      message.yPlayer = object.yPlayer;
    } else {
      message.yPlayer = "";
    }
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
