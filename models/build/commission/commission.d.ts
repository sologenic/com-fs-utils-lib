import _m0 from "protobufjs/minimal";
import { Decimal } from "../sologenic/com-fs-utils-lib/go/decimal/decimal";
export declare const protobufPackage = "commission";
export declare enum CommissionType {
    NOT_USED_COMMISSION_TYPE = 0,
    /** NOTIONAL - Charge commission on a per order basis (default) */
    NOTIONAL = 1,
    /** QTY - Charge commission on a per qty/contract basis, pro rated */
    QTY = 2,
    /** BPS - Commission expressed in basis points (percent), converted to notional amount for purposes of calculating commission(max two decimal places) */
    BPS = 3,
    UNRECOGNIZED = -1
}
export declare function commissionTypeFromJSON(object: any): CommissionType;
export declare function commissionTypeToJSON(object: CommissionType): string;
/** Broker API specific commission fields for user level (overrrides organization level) */
export interface CommissionSettings {
    /** Commission charged for the order */
    Commission?: Decimal | undefined;
    /** How commission field value is calculated */
    CommissionType?: CommissionType | undefined;
}
export declare const CommissionSettings: {
    encode(message: CommissionSettings, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): CommissionSettings;
    fromJSON(object: any): CommissionSettings;
    toJSON(message: CommissionSettings): unknown;
    create<I extends {
        Commission?: {
            Value?: number | undefined;
            Exp?: number | undefined;
        } | undefined;
        CommissionType?: CommissionType | undefined;
    } & {
        Commission?: ({
            Value?: number | undefined;
            Exp?: number | undefined;
        } & {
            Value?: number | undefined;
            Exp?: number | undefined;
        } & { [K in Exclude<keyof I["Commission"], keyof Decimal>]: never; }) | undefined;
        CommissionType?: CommissionType | undefined;
    } & { [K_1 in Exclude<keyof I, keyof CommissionSettings>]: never; }>(base?: I | undefined): CommissionSettings;
    fromPartial<I_1 extends {
        Commission?: {
            Value?: number | undefined;
            Exp?: number | undefined;
        } | undefined;
        CommissionType?: CommissionType | undefined;
    } & {
        Commission?: ({
            Value?: number | undefined;
            Exp?: number | undefined;
        } & {
            Value?: number | undefined;
            Exp?: number | undefined;
        } & { [K_2 in Exclude<keyof I_1["Commission"], keyof Decimal>]: never; }) | undefined;
        CommissionType?: CommissionType | undefined;
    } & { [K_3 in Exclude<keyof I_1, keyof CommissionSettings>]: never; }>(object: I_1): CommissionSettings;
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
