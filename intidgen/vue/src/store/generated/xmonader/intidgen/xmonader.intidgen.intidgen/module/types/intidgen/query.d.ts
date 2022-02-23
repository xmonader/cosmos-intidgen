import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../intidgen/params";
export declare const protobufPackage = "xmonader.intidgen.intidgen";
/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}
/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
    /** params holds all the parameters of this module. */
    params: Params | undefined;
}
export interface QueryLastidRequest {
}
export interface QueryLastidResponse {
    lastid: string;
}
export declare const QueryParamsRequest: {
    encode(_: QueryParamsRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest;
    fromJSON(_: any): QueryParamsRequest;
    toJSON(_: QueryParamsRequest): unknown;
    fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest;
};
export declare const QueryParamsResponse: {
    encode(message: QueryParamsResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse;
    fromJSON(object: any): QueryParamsResponse;
    toJSON(message: QueryParamsResponse): unknown;
    fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse;
};
export declare const QueryLastidRequest: {
    encode(_: QueryLastidRequest, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryLastidRequest;
    fromJSON(_: any): QueryLastidRequest;
    toJSON(_: QueryLastidRequest): unknown;
    fromPartial(_: DeepPartial<QueryLastidRequest>): QueryLastidRequest;
};
export declare const QueryLastidResponse: {
    encode(message: QueryLastidResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): QueryLastidResponse;
    fromJSON(object: any): QueryLastidResponse;
    toJSON(message: QueryLastidResponse): unknown;
    fromPartial(object: DeepPartial<QueryLastidResponse>): QueryLastidResponse;
};
/** Query defines the gRPC querier service. */
export interface Query {
    /** Parameters queries the parameters of the module. */
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    /** Queries a list of Lastid items. */
    Lastid(request: QueryLastidRequest): Promise<QueryLastidResponse>;
}
export declare class QueryClientImpl implements Query {
    private readonly rpc;
    constructor(rpc: Rpc);
    Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
    Lastid(request: QueryLastidRequest): Promise<QueryLastidResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};
