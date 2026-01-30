import _m0 from "protobufjs/minimal";
export declare const protobufPackage = "audit";
export interface Audit {
    ChangedBy: string;
    ChangedAt: Date | undefined;
    Reason?: string | undefined;
}
export declare const Audit: {
    encode(message: Audit, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Audit;
    fromJSON(object: any): Audit;
    toJSON(message: Audit): unknown;
    create<I extends {
        ChangedBy?: string | undefined;
        ChangedAt?: Date | undefined;
        Reason?: string | undefined;
    } & {
        ChangedBy?: string | undefined;
        ChangedAt?: Date | undefined;
        Reason?: string | undefined;
    } & { [K in Exclude<keyof I, keyof Audit>]: never; }>(base?: I | undefined): Audit;
    fromPartial<I_1 extends {
        ChangedBy?: string | undefined;
        ChangedAt?: Date | undefined;
        Reason?: string | undefined;
    } & {
        ChangedBy?: string | undefined;
        ChangedAt?: Date | undefined;
        Reason?: string | undefined;
    } & { [K_1 in Exclude<keyof I_1, keyof Audit>]: never; }>(object: I_1): Audit;
};
type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;
export type DeepPartial<T> = T extends Builtin ? T : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P : P & {
    [K in keyof P]: Exact<P[K], I[K]>;
} & {
    [K in Exclude<keyof I, KeysOfUnion<P>>]: never;
};
export {};
