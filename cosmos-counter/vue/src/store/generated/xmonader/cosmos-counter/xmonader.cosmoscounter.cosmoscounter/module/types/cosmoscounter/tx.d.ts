import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "xmonader.cosmoscounter.cosmoscounter";
export interface MsgIncrCounter {
    creator: string;
}
export interface MsgIncrCounterResponse {
    count: number;
}
export declare const MsgIncrCounter: {
    encode(message: MsgIncrCounter, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgIncrCounter;
    fromJSON(object: any): MsgIncrCounter;
    toJSON(message: MsgIncrCounter): unknown;
    fromPartial(object: DeepPartial<MsgIncrCounter>): MsgIncrCounter;
};
export declare const MsgIncrCounterResponse: {
    encode(message: MsgIncrCounterResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgIncrCounterResponse;
    fromJSON(object: any): MsgIncrCounterResponse;
    toJSON(message: MsgIncrCounterResponse): unknown;
    fromPartial(object: DeepPartial<MsgIncrCounterResponse>): MsgIncrCounterResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    IncrCounter(request: MsgIncrCounter): Promise<MsgIncrCounterResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    IncrCounter(request: MsgIncrCounter): Promise<MsgIncrCounterResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
