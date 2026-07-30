# Commission Settings Model

This is a broker commission settings model that is shared across several models. It provides a standardized way to configure commission calculations for broker operations.

## Message: `CommissionSettings`

Represents commission settings at the user level, which can override organization-level settings.

### Fields

- **`Commission`** (decimal.Decimal, required): The commission amount charged for an order. Must be between 0 and 10000 with at most 2 fraction digits.
- **`CommissionType`** (CommissionType, required): Specifies how the commission value is calculated. Must be a defined enum value other than `NOT_USED_COMMISSION_TYPE`.

### Usage

This model is typically embedded in user-level configurations to allow per-user commission overrides. When `CommissionSettings` is provided, both fields are required and validated together:
- `Commission` provides the base value
- `CommissionType` determines how that value is applied
- Requests with null/missing `Commission` or `CommissionType` are rejected and must not be persisted

## Enum: `CommissionType`

Defines the method by which commission is calculated.

### Values

- **`NOT_USED_COMMISSION_TYPE`** (0): Default/unused value
- **`NOTIONAL`** (1): Charge commission on a per-order basis (default). The commission amount is applied directly to the order value
- **`QTY`** (2): Charge commission on a per-quantity/contract basis, pro-rated. The commission is calculated based on the number of units
- **`BPS`** (3): Commission expressed in basis points (percentage). The value is converted to a notional amount for commission calculation (max two decimal places). One basis point = 0.01%

### Example Scenarios

- **NOTIONAL**: Fixed $10 commission per order
- **QTY**: $0.50 commission per share/contract
- **BPS**: 25 basis points (0.25%) of the order value

### Notes

- Both `Commission` and `CommissionType` are required whenever `CommissionSettings` is set
- `NOT_USED_COMMISSION_TYPE` is not a valid value for persisted commission settings
- For `BPS` calculations, ensure the commission value represents basis points (e.g., 25 for 0.25%)
- The `decimal.Decimal` type ensures precise commission calculations without floating-point errors