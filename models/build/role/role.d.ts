export declare const protobufPackage = "role";
export declare enum Role {
    NOT_USED_ROLE = 0,
    NORMAL_USER = 1,
    VIEWER = 2,
    BROKER_ASSET_ADMINISTRATOR = 3,
    ORGANIZATION_ADMINISTRATOR = 4,
    BROKER_ASSET_VIEWER = 5,
    /** SOLOGENIC_ADMINISTRATOR - Highest level of administrator. Can only create new Organizations and assign an initial administrator in there */
    SOLOGENIC_ADMINISTRATOR = 6,
    /** TX_ADMINISTRATOR - Alias for SOLOGENIC_ADMINISTRATOR */
    TX_ADMINISTRATOR = 6,
    KYC_ADMINISTRATOR = 7,
    UNRECOGNIZED = -1
}
export declare function roleFromJSON(object: any): Role;
export declare function roleToJSON(object: Role): string;
