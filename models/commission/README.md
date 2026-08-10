# Commission Settings Model

This is a broker commission settings model that is shared across several models. It provides a standardized way to configure commission calculations for broker operations.

## Message: `CommissionSettings`

Represents commission settings at the user level, which can override organization-level settings.

### Fields

- **`Commission`** (decimal.Decimal, required): The commission amount charged for an order. Must be between 0 and 10000 with at most 2 fraction digits. Stored as `Value * 10^Exp` (for example `"12.50"` is `{Value: 1250, Exp: -2}`).
- **`CommissionType`** (CommissionType, required): Specifies how the commission value is calculated. Must be a defined enum value other than `NOT_USED_COMMISSION_TYPE`.

### Usage

This model is typically embedded in user-level configurations to allow per-user commission overrides. When `CommissionSettings` is provided, both fields are required and validated together:
- `Commission` provides the base value
- `CommissionType` determines how that value is applied
- Requests with null/missing `Commission` or `CommissionType` are rejected and must not be persisted

When submitting an order, provide the commission value as a numeric string (e.g. `"1.50"`) alongside `CommissionType` to avoid floating-point precision loss.

## Enum: `CommissionType`

Defines the method by which commission is calculated.

### Values

- **`NOT_USED_COMMISSION_TYPE`** (0): Default/unused value
- **`NOTIONAL`** (1): Flat fee per order, regardless of quantity
- **`QTY`** (2): Fee per quantity/contract, pro-rated (`quantity * commission`)
- **`BPS`** (3): Fee in basis points of order notional (1 bps = 0.01%). Up to two decimal places.

### Example Scenarios

- **NOTIONAL**: Submitted `"2.50"` → flat $2.50 commission for the order
- **QTY**: Submitted `"0.05"` with quantity 100 → $5.00 total commission
- **BPS**: Submitted `"12.50"` (12.50 bps / 0.125%) on a $10,000 notional → $12.50 total commission

### Notes

- Both `Commission` and `CommissionType` are required whenever `CommissionSettings` is set
- `NOT_USED_COMMISSION_TYPE` is not a valid value for persisted commission settings
- For `BPS`, the commission value is basis points (e.g. `25` for 0.25%, or `12.50` for 0.125%)
- The `decimal.Decimal` type ensures precise commission calculations without floating-point errors
