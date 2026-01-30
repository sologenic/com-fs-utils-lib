package decimal

import (
	sdecimal "github.com/shopspring/decimal"
)

func FromFloat64(f float64) *Decimal {
	d := sdecimal.NewFromFloat(f)
	return &Decimal{
		Value: d.CoefficientInt64(),
		Exp:   d.Exponent(),
	}
}

// Adds another decimal to this decimal and returns a new decimal
func (d *Decimal) AddDecimal(d2 *Decimal) *Decimal {
	dec1 := sdecimal.New(d.Value, d.Exp)
	dec2 := sdecimal.New(d2.Value, d2.Exp)
	result := dec1.Add(dec2)
	return &Decimal{
		Value: result.CoefficientInt64(),
		Exp:   result.Exponent(),
	}
}

// Subtracts another decimal from this decimal and returns a new decimal
func (d *Decimal) SubDecimal(d2 *Decimal) *Decimal {
	dec1 := sdecimal.New(d.Value, d.Exp)
	dec2 := sdecimal.New(d2.Value, d2.Exp)
	result := dec1.Sub(dec2)
	return &Decimal{
		Value: result.CoefficientInt64(),
		Exp:   result.Exponent(),
	}
}

// Multiplies this decimal by another decimal and returns a new decimal
func (d *Decimal) MulDecimal(d2 *Decimal) *Decimal {
	dec1 := sdecimal.New(d.Value, d.Exp)
	dec2 := sdecimal.New(d2.Value, d2.Exp)
	result := dec1.Mul(dec2)
	return &Decimal{
		Value: result.CoefficientInt64(),
		Exp:   result.Exponent(),
	}
}

func (d *Decimal) Mul(d2 float64) float64 {
	f, _ := sdecimal.New(d.Value, d.Exp).Mul(sdecimal.NewFromFloat(d2)).Float64()
	return f
}

func (d *Decimal) Float64() float64 {
	f, _ := sdecimal.New(d.Value, d.Exp).Float64()
	return f
}

func (d *Decimal) IsZero() bool {
	return d.Value == 0
}

func (d *Decimal) DivDecimal(d2 *Decimal) *Decimal {
	dec1 := sdecimal.New(d.Value, d.Exp)
	dec2 := sdecimal.New(d2.Value, d2.Exp)
	result := dec1.Div(dec2)
	return &Decimal{
		Value: result.CoefficientInt64(),
		Exp:   result.Exponent(),
	}
}

func (d *Decimal) Round(precision int32) *Decimal {
	dec := sdecimal.New(d.Value, d.Exp)
	result := dec.Round(precision)
	return &Decimal{
		Value: result.CoefficientInt64(),
		Exp:   result.Exponent(),
	}
}
