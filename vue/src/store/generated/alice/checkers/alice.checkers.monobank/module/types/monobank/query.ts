/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../monobank/params";
import { Balance } from "../monobank/balance";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";

export const protobufPackage = "alice.checkers.monobank";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetBalanceRequest {
  address: string;
}

export interface QueryGetBalanceResponse {
  balance: Balance | undefined;
}

export interface QueryAllBalanceRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllBalanceResponse {
  balance: Balance[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetBalanceRequest: object = { address: "" };

export const QueryGetBalanceRequest = {
  encode(
    message: QueryGetBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetBalanceRequest } as QueryGetBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBalanceRequest {
    const message = { ...baseQueryGetBalanceRequest } as QueryGetBalanceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    return message;
  },

  toJSON(message: QueryGetBalanceRequest): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBalanceRequest>
  ): QueryGetBalanceRequest {
    const message = { ...baseQueryGetBalanceRequest } as QueryGetBalanceRequest;
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    return message;
  },
};

const baseQueryGetBalanceResponse: object = {};

export const QueryGetBalanceResponse = {
  encode(
    message: QueryGetBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.balance !== undefined) {
      Balance.encode(message.balance, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryGetBalanceResponse,
    } as QueryGetBalanceResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.balance = Balance.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetBalanceResponse {
    const message = {
      ...baseQueryGetBalanceResponse,
    } as QueryGetBalanceResponse;
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = Balance.fromJSON(object.balance);
    } else {
      message.balance = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetBalanceResponse): unknown {
    const obj: any = {};
    message.balance !== undefined &&
      (obj.balance = message.balance
        ? Balance.toJSON(message.balance)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetBalanceResponse>
  ): QueryGetBalanceResponse {
    const message = {
      ...baseQueryGetBalanceResponse,
    } as QueryGetBalanceResponse;
    if (object.balance !== undefined && object.balance !== null) {
      message.balance = Balance.fromPartial(object.balance);
    } else {
      message.balance = undefined;
    }
    return message;
  },
};

const baseQueryAllBalanceRequest: object = {};

export const QueryAllBalanceRequest = {
  encode(
    message: QueryAllBalanceRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBalanceRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllBalanceRequest } as QueryAllBalanceRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllBalanceRequest {
    const message = { ...baseQueryAllBalanceRequest } as QueryAllBalanceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBalanceRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBalanceRequest>
  ): QueryAllBalanceRequest {
    const message = { ...baseQueryAllBalanceRequest } as QueryAllBalanceRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllBalanceResponse: object = {};

export const QueryAllBalanceResponse = {
  encode(
    message: QueryAllBalanceResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.balance) {
      Balance.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllBalanceResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseQueryAllBalanceResponse,
    } as QueryAllBalanceResponse;
    message.balance = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.balance.push(Balance.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllBalanceResponse {
    const message = {
      ...baseQueryAllBalanceResponse,
    } as QueryAllBalanceResponse;
    message.balance = [];
    if (object.balance !== undefined && object.balance !== null) {
      for (const e of object.balance) {
        message.balance.push(Balance.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllBalanceResponse): unknown {
    const obj: any = {};
    if (message.balance) {
      obj.balance = message.balance.map((e) =>
        e ? Balance.toJSON(e) : undefined
      );
    } else {
      obj.balance = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllBalanceResponse>
  ): QueryAllBalanceResponse {
    const message = {
      ...baseQueryAllBalanceResponse,
    } as QueryAllBalanceResponse;
    message.balance = [];
    if (object.balance !== undefined && object.balance !== null) {
      for (const e of object.balance) {
        message.balance.push(Balance.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Balance by index. */
  Balance(request: QueryGetBalanceRequest): Promise<QueryGetBalanceResponse>;
  /** Queries a list of Balance items. */
  BalanceAll(request: QueryAllBalanceRequest): Promise<QueryAllBalanceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(
      "alice.checkers.monobank.Query",
      "Params",
      data
    );
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Balance(request: QueryGetBalanceRequest): Promise<QueryGetBalanceResponse> {
    const data = QueryGetBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "alice.checkers.monobank.Query",
      "Balance",
      data
    );
    return promise.then((data) =>
      QueryGetBalanceResponse.decode(new Reader(data))
    );
  }

  BalanceAll(
    request: QueryAllBalanceRequest
  ): Promise<QueryAllBalanceResponse> {
    const data = QueryAllBalanceRequest.encode(request).finish();
    const promise = this.rpc.request(
      "alice.checkers.monobank.Query",
      "BalanceAll",
      data
    );
    return promise.then((data) =>
      QueryAllBalanceResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

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
