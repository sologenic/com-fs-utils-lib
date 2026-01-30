package decimal

import (
	"github.com/shopspring/decimal"
)

func DecimalToInternalDecimal(d *decimal.Decimal) *Decimal {
	if d == nil {
		return nil
	}
	return &Decimal{
		Value: d.CoefficientInt64(),
		Exp:   d.Exponent(),
	}
}

func InternalDecimalToDecimal(d *Decimal) *decimal.Decimal {
	if d == nil {
		return nil
	}
    dec := decimal.New(d.Value, d.Exp)
    return &dec
}
