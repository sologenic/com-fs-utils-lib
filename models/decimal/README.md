# Decimal

Represents a fixed-point decimal number.
Use `value` and `exp` with `decimal.New()` from `github.com/shopspring/decimal` to create a decimal.Decimal object.
This ensures precise handling without converting to floats or doubles.
The number is represented as: Number = value * 10 ^ exp.
When converting from decimal.Decimal to this internal Decimal message, use the `decimalToInternalDecimal` function in this repo.
