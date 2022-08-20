/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "rio.rio";

export interface MsgCreateCert {
  creator: string;
  title: string;
}

export interface MsgCreateCertResponse {}

const baseMsgCreateCert: object = { creator: "", title: "" };

export const MsgCreateCert = {
  encode(message: MsgCreateCert, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.title !== "") {
      writer.uint32(18).string(message.title);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCert {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCert } as MsgCreateCert;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.title = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateCert {
    const message = { ...baseMsgCreateCert } as MsgCreateCert;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = String(object.title);
    } else {
      message.title = "";
    }
    return message;
  },

  toJSON(message: MsgCreateCert): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.title !== undefined && (obj.title = message.title);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateCert>): MsgCreateCert {
    const message = { ...baseMsgCreateCert } as MsgCreateCert;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.title !== undefined && object.title !== null) {
      message.title = object.title;
    } else {
      message.title = "";
    }
    return message;
  },
};

const baseMsgCreateCertResponse: object = {};

export const MsgCreateCertResponse = {
  encode(_: MsgCreateCertResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateCertResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateCertResponse } as MsgCreateCertResponse;
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

  fromJSON(_: any): MsgCreateCertResponse {
    const message = { ...baseMsgCreateCertResponse } as MsgCreateCertResponse;
    return message;
  },

  toJSON(_: MsgCreateCertResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgCreateCertResponse>): MsgCreateCertResponse {
    const message = { ...baseMsgCreateCertResponse } as MsgCreateCertResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CreateCert(request: MsgCreateCert): Promise<MsgCreateCertResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  CreateCert(request: MsgCreateCert): Promise<MsgCreateCertResponse> {
    const data = MsgCreateCert.encode(request).finish();
    const promise = this.rpc.request("rio.rio.Msg", "CreateCert", data);
    return promise.then((data) =>
      MsgCreateCertResponse.decode(new Reader(data))
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
