import _m0 from "protobufjs/minimal";
export declare const protobufPackage = "orderproperties";
/** Side is order side. */
export declare enum Side {
    /** SIDE_NOT_APPLICABLE - SIDE_NOT_APPLICABLE reserves the default value, to protect against unexpected settings. */
    SIDE_NOT_APPLICABLE = 0,
    /** SIDE_BUY - SIDE_BUY means that the order is to buy base_denom quantity with the price. */
    SIDE_BUY = 1,
    /** SIDE_SELL - SIDE_SELL means that the order is to sell base_denom quantity with the price. */
    SIDE_SELL = 2,
    UNRECOGNIZED = -1
}
export declare function sideFromJSON(object: any): Side;
export declare function sideToJSON(object: Side): string;
export declare enum TimeInForce {
    NOT_USED_TIME_IN_FORCE = 0,
    DAY = 1,
    /** GOOD_TIL_CANCELED - GTC */
    GOOD_TIL_CANCELED = 2,
    /** AT_THE_OPENING - OPG */
    AT_THE_OPENING = 3,
    /** AT_THE_CLOSE - CLS */
    AT_THE_CLOSE = 4,
    /** IMMEDIATE_OR_CANCEL - IOC */
    IMMEDIATE_OR_CANCEL = 5,
    /** FILL_OR_KILL - FOK */
    FILL_OR_KILL = 6,
    /** GOOD_TIL_TIME - GTT */
    GOOD_TIL_TIME = 7,
    UNRECOGNIZED = -1
}
export declare function timeInForceFromJSON(object: any): TimeInForce;
export declare function timeInForceToJSON(object: TimeInForce): string;
export declare enum OrderType {
    /** NOT_APPLICABLE_ORDER_TYPE - Changed not used to NOT_APPLICABLE_ORDER_TYPE to avoid naming conflicts with PURCHASE in the action enum */
    NOT_APPLICABLE_ORDER_TYPE = 0,
    /** ORDER_TYPE_PURCHASE - Avoid naming conflicts with PURCHASE in the action enum */
    ORDER_TYPE_PURCHASE = 1,
    ORDER_TYPE_SELL = 2,
    UNRECOGNIZED = -1
}
export declare function orderTypeFromJSON(object: any): OrderType;
export declare function orderTypeToJSON(object: OrderType): string;
export declare enum ProcessState {
    NOT_USED_PROCESS_STATE = 0,
    PROCESS_STATE_OPEN = 1,
    PROCESS_STATE_LOCKED = 2,
    PROCESS_STATE_FAILED = 3,
    PROCESS_STATE_PROCESSED = 4,
    UNRECOGNIZED = -1
}
export declare function processStateFromJSON(object: any): ProcessState;
export declare function processStateToJSON(object: ProcessState): string;
export declare enum ClearingBroker {
    NOT_USED_CLEARING_BROKER = 0,
    ALPACA = 1,
    RQD = 2,
    UNRECOGNIZED = -1
}
export declare function clearingBrokerFromJSON(object: any): ClearingBroker;
export declare function clearingBrokerToJSON(object: ClearingBroker): string;
export interface ProcessInfo {
    ProcessState: ProcessState;
    ProcessedAt: Date | undefined;
}
export declare const ProcessInfo: {
    encode(message: ProcessInfo, writer?: _m0.Writer): _m0.Writer;
    decode(input: _m0.Reader | Uint8Array, length?: number): ProcessInfo;
    fromJSON(object: any): ProcessInfo;
    toJSON(message: ProcessInfo): unknown;
    create<I extends {
        ProcessState?: ProcessState | undefined;
        ProcessedAt?: Date | undefined;
    } & {
        ProcessState?: ProcessState | undefined;
        ProcessedAt?: Date | undefined;
    } & { [K in Exclude<keyof I, keyof ProcessInfo>]: never; }>(base?: I | undefined): ProcessInfo;
    fromPartial<I_1 extends {
        ProcessState?: ProcessState | undefined;
        ProcessedAt?: Date | undefined;
    } & {
        ProcessState?: ProcessState | undefined;
        ProcessedAt?: Date | undefined;
    } & { [K_1 in Exclude<keyof I_1, keyof ProcessInfo>]: never; }>(object: I_1): ProcessInfo;
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
