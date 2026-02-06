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
    /** TX_ACCOUNT_MANAGER - Manages partners and partner assignments */
    TX_ACCOUNT_MANAGER = 8,
    /** TX_USER_MANAGER - Manages user/referral aspects */
    TX_USER_MANAGER = 9,
    /** PARTNER_BROKER_ADMIN - Can manage partner users for broker partner, and view partner dashboard */
    PARTNER_BROKER_ADMIN = 10,
    /** PARTNER_BROKER_ACCOUNT_MANAGER - Day-to-day broker ops, create orgs via wizard, approvals if allowed */
    PARTNER_BROKER_ACCOUNT_MANAGER = 11,
    /** PARTNER_SALES_ADMIN - Partner scope hierarchy for sales channel */
    PARTNER_SALES_ADMIN = 12,
    PARTNER_SALES_ACCOUNT_MANAGER = 13,
    UNRECOGNIZED = -1
}
export declare function roleFromJSON(object: any): Role;
export declare function roleToJSON(object: Role): string;
