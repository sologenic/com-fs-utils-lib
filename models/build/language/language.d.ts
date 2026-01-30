import _m0 from "protobufjs/minimal";
export declare const protobufPackage = "language";
export declare enum Lang {
    LANG_NOT_USED = 0,
    ENGLISH = 1,
    SPANISH = 2,
    KOREAN = 3,
    UNRECOGNIZED = -1
}
export declare function langFromJSON(object: any): Lang;
export declare function langToJSON(object: Lang): string;
export interface Language {
    Language: Lang;
    UserConfigured: boolean;
}
export declare const Language: {
    encode(message: Language, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): Language;
    fromJSON(object: any): Language;
    toJSON(message: Language): unknown;
    create<I extends {
        Language?: Lang | undefined;
        UserConfigured?: boolean | undefined;
    } & {
        Language?: Lang | undefined;
        UserConfigured?: boolean | undefined;
    } & { [K in Exclude<keyof I, keyof Language>]: never; }>(base?: I | undefined): Language;
    fromPartial<I_1 extends {
        Language?: Lang | undefined;
        UserConfigured?: boolean | undefined;
    } & {
        Language?: Lang | undefined;
        UserConfigured?: boolean | undefined;
    } & { [K_1 in Exclude<keyof I_1, keyof Language>]: never; }>(object: I_1): Language;
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
