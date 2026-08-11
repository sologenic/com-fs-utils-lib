import _m0 from "protobufjs/minimal";
import { Decimal } from "../sologenic/com-fs-utils-lib/models/decimal/decimal";
export declare const protobufPackage = "commission";
/**
 * CommissionType defines the specific method used to interpret and calculate
 * the commission charged to the end user for a given order.
 *
 * SUBMISSION GUIDELINES:
 * When submitting an order, the actual commission value must be provided alongside
 * this enum as a string representing a numeric value (e.g., "1.50") to prevent
 * floating-point precision loss.
 */
export declare enum CommissionType {
    NOT_USED_COMMISSION_TYPE = 0,
    /**
     * NOTIONAL - Charge commission on a fixed per-order basis.
     * The submitted string value is applied as a flat fee for the entire order,
     * regardless of the quantity of shares or contracts executed.
     *
     * Example Submission:
     * - Submitted Value: "2.50"
     * - Result: A flat $2.50 commission is collected for the order.
     */
    NOTIONAL = 1,
    /**
     * QTY - Charge commission on a per-quantity or per-contract basis, pro-rated.
     * The total commission is calculated by multiplying the executed order quantity
     * by the submitted string value.
     *
     * Example Submission:
     * - Submitted Value: "0.05"
     * - Order Quantity: 100 shares
     * - Result: $5.00 total commission (100 * 0.05).
     */
    QTY = 2,
    /**
     * BPS - Charge commission in basis points (BPS).
     * The commission is calculated as a percentage of the total notional value
     * of the order. 1 basis point equals 0.01% (or 0.0001).
     * The system automatically converts the order to a notional amount to
     * calculate the final fee. The submitted BPS value can have up to two decimal places.
     *
     * Example Submission:
     * - Submitted Value: "12.50" (representing 12.50 bps or 0.125%)
     * - Order Notional Value: $10,000
     * - Result: $12.50 total commission ($10,000 * 0.125%).
     */
    BPS = 3,
    UNRECOGNIZED = -1
}
export declare function commissionTypeFromJSON(object: any): CommissionType;
export declare function commissionTypeToJSON(object: CommissionType): string;
/** Broker API specific commission fields for user level (overrrides organization level) */
export interface CommissionSettings {
    Commission: Decimal | undefined;
    CommissionType: CommissionType;
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
