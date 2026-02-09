# Commission Settings Model

This is a broker commission settings model that is shared across several models. It provides a standardized way to configure commission calculations for broker operations.

## Message: `CommissionSettings`

Represents commission settings at the user level, which can override organization-level settings.

### Fields

- **`Commission`** (decimal.Decimal, optional): The commission amount charged for an order
- **`CommissionType`** (CommissionType, optional): Specifies how the commission value is calculated

### Usage

This model is typically embedded in user-level configurations to allow per-user commission overrides. When both fields are set, they work together:
- `Commission` provides the base value
- `CommissionType` determines how that value is applied

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

- When `CommissionType` is not specified, `NOTIONAL` is typically used as the default
- For `BPS` calculations, ensure the commission value represents basis points (e.g., 25 for 0.25%)
- The `decimal.Decimal` type ensures precise commission calculations without floating-point errors