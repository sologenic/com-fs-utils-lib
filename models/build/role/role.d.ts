export declare const protobufPackage = "role";
export declare enum Role {
    NOT_USED_ROLE = 0,
    /** NORMAL_USER - End-user role */
    NORMAL_USER = 1,
    /** BROKER_ASSET_ADMINISTRATOR - Broker & asset management roles for an organization */
    BROKER_ASSET_ADMINISTRATOR = 3,
    ORGANIZATION_ADMINISTRATOR = 4,
    /** TX_SYSTEM_MANAGER - Highest level of access including access to secrets and important settings */
    TX_SYSTEM_MANAGER = 21,
    /** TX_ADMINISTRATOR - Highest level of access excluding access to secrets and important settings */
    TX_ADMINISTRATOR = 6,
    /**
     * TX_ACCOUNT_MANAGER - Manages everything related to partners
     * TODO: Rename to TX_PARTNERS_MANAGER
     */
    TX_ACCOUNT_MANAGER = 8,
    /** TX_MARKETING_MANAGER - Manages marketing aspects */
    TX_MARKETING_MANAGER = 20,
    /**
     * KYC_ADMINISTRATOR - Manages user information, KYC stuff
     * TODO: Rename to TX_KYC_MANAGER
     */
    KYC_ADMINISTRATOR = 7,
    /** PARTNER_BROKER_ADMIN - TODO: Rename (get rid of "admin" and replace with "manager") */
    PARTNER_BROKER_ADMIN = 10,
    /**
     * PARTNER_BROKER_ACCOUNT_MANAGER - Day-to-day operations
     * TODO: Rename and think if it can be merged with PARTNER_BROKER_ADMIN
     */
    PARTNER_BROKER_ACCOUNT_MANAGER = 11,
    /** PARTNER_SALES_ADMIN - TODO: Rename (get rid of "admin" and replace with "manager") */
    PARTNER_SALES_ADMIN = 12,
    /**
     * PARTNER_SALES_ACCOUNT_MANAGER - Day-to-day operations
     * TODO: Rename and think if it can be merged with PARTNER_SALES_ADMIN
     */
    PARTNER_SALES_ACCOUNT_MANAGER = 13,
    UNRECOGNIZED = -1
}
export declare function roleFromJSON(object: any): Role;
export declare function roleToJSON(object: Role): string;
