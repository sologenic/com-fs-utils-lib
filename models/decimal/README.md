# Decimal

Represents a fixed-point decimal number for precise financial and mathematical calculations without floating-point precision errors.

## Message: `Decimal`

A fixed-point decimal representation that avoids the precision issues of floating-point types.

### Fields

- **`Value`** (int64, required): The integer value component
- **`Exp`** (int32, required): The exponent component (power of 10)

### Representation

The decimal number is calculated as: **Number = Value × 10^Exp**

### Examples

- `Value: 12345, Exp: -2` = 123.45
- `Value: 100, Exp: 0` = 100
- `Value: 1, Exp: -4` = 0.0001

## Usage in Go

### Creating from shopspring/decimal

```go
import "github.com/shopspring/decimal"

// Create a decimal.Decimal
d := decimal.NewFromFloat(123.45)

// Convert to internal Decimal message
// Use the decimalToInternalDecimal function from this package
internalDecimal := decimalToInternalDecimal(d)
```

### Creating from internal Decimal message

```go
import "github.com/shopspring/decimal"

// Convert from internal Decimal message
internalDecimal := &decimal.Decimal{Value: 12345, Exp: -2}
d := decimal.New(internalDecimal.Value, internalDecimal.Exp)
```

## Usage in TypeScript

The generated TypeScript code provides type-safe access to the `Value` and `Exp` fields.

## Important Notes

- **Never convert to float/double**: Always use the `Value` and `Exp` fields directly
- **Precision**: This format maintains exact precision for financial calculations
- **Conversion**: Use the provided conversion functions rather than manual conversion
- **Import requirement**: When importing this `Decimal` message in another proto file, ensure the build script includes the corresponding `protoc` command to generate Go and gRPC code

## Use Cases

- Financial calculations (prices, commissions, balances)
- Precise mathematical operations
- Currency amounts
- Percentage calculations
- Any scenario requiring exact decimal precision
