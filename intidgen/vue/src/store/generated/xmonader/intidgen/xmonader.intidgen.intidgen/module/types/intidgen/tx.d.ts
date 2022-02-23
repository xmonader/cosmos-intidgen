import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "xmonader.intidgen.intidgen";
export interface MsgGenid {
    creator: string;
}
export interface MsgGenidResponse {
    id: string;
}
export declare const MsgGenid: {
    encode(message: MsgGenid, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgGenid;
    fromJSON(object: any): MsgGenid;
    toJSON(message: MsgGenid): unknown;
    fromPartial(object: DeepPartial<MsgGenid>): MsgGenid;
};
export declare const MsgGenidResponse: {
    encode(message: MsgGenidResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgGenidResponse;
    fromJSON(object: any): MsgGenidResponse;
    toJSON(message: MsgGenidResponse): unknown;
    fromPartial(object: DeepPartial<MsgGenidResponse>): MsgGenidResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Genid(request: MsgGenid): Promise<MsgGenidResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Genid(request: MsgGenid): Promise<MsgGenidResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
