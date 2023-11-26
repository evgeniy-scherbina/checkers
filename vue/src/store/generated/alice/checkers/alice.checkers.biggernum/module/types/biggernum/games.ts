/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "alice.checkers.biggernum";

export interface Games {
  index: string;
  player1: string;
  player2: string;
}

const baseGames: object = { index: "", player1: "", player2: "" };

export const Games = {
  encode(message: Games, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.player1 !== "") {
      writer.uint32(18).string(message.player1);
    }
    if (message.player2 !== "") {
      writer.uint32(26).string(message.player2);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Games {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGames } as Games;
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Games {
    const message = { ...baseGames } as Games;
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
    return message;
  },

  toJSON(message: Games): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.player1 !== undefined && (obj.player1 = message.player1);
    message.player2 !== undefined && (obj.player2 = message.player2);
    return obj;
  },

  fromPartial(object: DeepPartial<Games>): Games {
    const message = { ...baseGames } as Games;
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
