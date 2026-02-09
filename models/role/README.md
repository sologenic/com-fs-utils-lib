# Role

The role model defines user roles and permissions within the authentication and authorization system. Roles are stored in the database and used to determine user permissions and access levels throughout the application.

## Enum: `Role`

Represents all available user roles in the system, organized in a hierarchical structure.

### Standard User Roles

- **`NOT_USED_ROLE`** (0): Default/unused value
- **`NORMAL_USER`** (1): Standard user with basic permissions
- **`VIEWER`** (2): Read-only access role

### Broker & Asset Management Roles

- **`BROKER_ASSET_ADMINISTRATOR`** (3): Administrator with full control over broker assets
- **`BROKER_ASSET_VIEWER`** (5): Read-only access to broker assets

### Organization & Administrative Roles

- **`ORGANIZATION_ADMINISTRATOR`** (4): Administrator for an organization with full organizational control
- **`SOLOGENIC_ADMINISTRATOR`** (6): Highest level administrator. Can create new organizations and assign initial administrators
- **`TX_ADMINISTRATOR`** (6): Alias for `SOLOGENIC_ADMINISTRATOR` (uses `allow_alias = true`)
- **`KYC_ADMINISTRATOR`** (7): Administrator responsible for Know Your Customer (KYC) operations

### Transaction & User Management Roles

- **`TX_ACCOUNT_MANAGER`** (8): Manages partners and partner assignments
- **`TX_USER_MANAGER`** (9): Manages user/referral aspects

### Partner Roles - Broker Scope

- **`PARTNER_BROKER_ADMIN`** (10): Can manage partner users for broker partner and view partner dashboard
- **`PARTNER_BROKER_ACCOUNT_MANAGER`** (11): Day-to-day broker operations, create orgs via wizard, approvals if allowed

### Partner Roles - Sales Channel Scope

- **`PARTNER_SALES_ADMIN`** (12): Sales channel administrator
- **`PARTNER_SALES_ACCOUNT_MANAGER`** (13): Sales channel account manager

## Role Hierarchy

The roles follow a hierarchical structure:

1. **Standard Users**: `NORMAL_USER`, `VIEWER`
2. **Broker Roles**: `BROKER_ASSET_ADMINISTRATOR`, `BROKER_ASSET_VIEWER`
3. **Organization Roles**: `ORGANIZATION_ADMINISTRATOR`
4. **System Roles**: `SOLOGENIC_ADMINISTRATOR` / `TX_ADMINISTRATOR` (highest level)
5. **Specialized Roles**: `KYC_ADMINISTRATOR`, `TX_ACCOUNT_MANAGER`, `TX_USER_MANAGER`
6. **Partner Roles**: Broker and Sales channel scopes

## Important Notes

### Alias Support

The enum uses `allow_alias = true` to support:
- `SOLOGENIC_ADMINISTRATOR` and `TX_ADMINISTRATOR` both map to value `6`
- This allows different naming conventions while maintaining the same permission level

### Scope Hierarchy

Partner roles are organized by scope:
- **Broker scope**: Roles 10-11 handle broker partner operations
- **Sales channel scope**: Roles 12-13 handle sales channel operations

## Use Cases

- User authentication and authorization
- Permission checking
- Access control lists (ACLs)
- Role-based access control (RBAC)
- Database role storage
- API endpoint authorization

## Notes

- Roles are stored in the database and referenced by their numeric values
- The `SOLOGENIC_ADMINISTRATOR` role has the highest privileges and can create organizations
- Partner roles are scoped to specific operational areas (broker vs sales)
- Always use `NOT_USED_ROLE` (0) as the default to catch uninitialized role values
