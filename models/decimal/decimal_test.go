package decimal

import (
	"math"
	"testing"

	sdecimal "github.com/shopspring/decimal"
)

func TestFromFloat64(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected float64
	}{
		{"Zero", 0.0, 0.0},
		{"Positive integer", 123.0, 123.0},
		{"Negative integer", -456.0, -456.0},
		{"Positive decimal", 123.456, 123.456},
		{"Negative decimal", -123.456, -123.456},
		{"Small decimal", 0.001, 0.001},
		{"Large number", 999999.999, 999999.999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FromFloat64(tt.input)
			if result == nil {
				t.Fatal("FromFloat64 returned nil")
			}

			// Convert back to float64 to compare
			actual := result.Float64()
			if math.Abs(actual-tt.expected) > 1e-10 {
				t.Errorf("FromFloat64(%f) = %f, expected %f", tt.input, actual, tt.expected)
			}
		})
	}
}

func TestDecimal_Add(t *testing.T) {
	tests := []struct {
		name     string
		d1       float64
		d2       float64
		expected float64
	}{
		{"Zero + Zero", 0.0, 0.0, 0.0},
		{"Positive + Positive", 123.456, 78.9, 202.356},
		{"Negative + Negative", -123.456, -78.9, -202.356},
		{"Positive + Negative", 123.456, -78.9, 44.556},
		{"Negative + Positive", -123.456, 78.9, -44.556},
		{"Large numbers", 999999.999, 1.001, 1000001.0},
		{"Small decimals", 0.001, 0.002, 0.003},
		{"Different scales", 1.2345, 67.89, 69.1245},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := FromFloat64(tt.d1)
			d2 := FromFloat64(tt.d2)

			// Store original values for verification
			origD1 := d1.Float64()
			origD2 := d2.Float64()

			// Perform addition
			result := d1.AddDecimal(d2)

			// Check result
			actual := result.Float64()
			if math.Abs(actual-tt.expected) > 1e-10 {
				t.Errorf("Add(%f, %f) = %f, expected %f", origD1, origD2, actual, tt.expected)
			}

			// Verify that d1 was not modified
			if d1.Float64() != origD1 {
				t.Errorf("Add operation modified the first operand: %f became %f", origD1, d1.Float64())
			}

			// Verify that d2 was not modified
			if d2.Float64() != origD2 {
				t.Errorf("Add operation modified the second operand: %f became %f", origD2, d2.Float64())
			}
		})
	}
}

func TestDecimal_Add_Precision(t *testing.T) {
	// Test precision with known decimal arithmetic issues in floating point
	d1 := FromFloat64(0.1)
	d2 := FromFloat64(0.2)

	result := d1.AddDecimal(d2)
	actual := result.Float64()

	// Should be exactly 0.3, not 0.30000000000000004
	expected := 0.3
	if math.Abs(actual-expected) > 1e-15 {
		t.Errorf("Precision test failed: 0.1 + 0.2 = %f, expected %f", actual, expected)
	}
}

func TestDecimal_SubDecimal(t *testing.T) {
	tests := []struct {
		name     string
		d1       float64
		d2       float64
		expected float64
	}{
		{"Zero - Zero", 0.0, 0.0, 0.0},
		{"Positive - Positive (larger - smaller)", 123.456, 78.9, 44.556},
		{"Positive - Positive (smaller - larger)", 78.9, 123.456, -44.556},
		{"Negative - Negative", -123.456, -78.9, -44.556},
		{"Positive - Negative", 123.456, -78.9, 202.356},
		{"Negative - Positive", -123.456, 78.9, -202.356},
		{"Large numbers", 1000001.0, 999999.999, 1.001},
		{"Small decimals", 0.003, 0.001, 0.002},
		{"Different scales", 69.1245, 67.89, 1.2345},
		{"Same number", 123.456, 123.456, 0.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := FromFloat64(tt.d1)
			d2 := FromFloat64(tt.d2)

			// Store original values for verification
			origD1 := d1.Float64()
			origD2 := d2.Float64()

			// Perform subtraction
			result := d1.SubDecimal(d2)

			// Check result
			actual := result.Float64()
			if math.Abs(actual-tt.expected) > 1e-10 {
				t.Errorf("SubDecimal(%f, %f) = %f, expected %f", origD1, origD2, actual, tt.expected)
			}

			// Verify that d1 was not modified
			if d1.Float64() != origD1 {
				t.Errorf("SubDecimal operation modified the first operand: %f became %f", origD1, d1.Float64())
			}

			// Verify that d2 was not modified
			if d2.Float64() != origD2 {
				t.Errorf("SubDecimal operation modified the second operand: %f became %f", origD2, d2.Float64())
			}
		})
	}
}

func TestDecimal_SubDecimal_Precision(t *testing.T) {
	// Test precision with known decimal arithmetic issues in floating point
	d1 := FromFloat64(0.3)
	d2 := FromFloat64(0.1)

	result := d1.SubDecimal(d2)
	actual := result.Float64()

	// Should be exactly 0.2, not 0.19999999999999998
	expected := 0.2
	if math.Abs(actual-expected) > 1e-15 {
		t.Errorf("Precision test failed: 0.3 - 0.1 = %f, expected %f", actual, expected)
	}
}

func TestDecimal_MulDecimal(t *testing.T) {
	tests := []struct {
		name     string
		d1       float64
		d2       float64
		expected float64
	}{
		{"Zero * Zero", 0.0, 0.0, 0.0},
		{"Zero * Positive", 0.0, 123.456, 0.0},
		{"Positive * Zero", 123.456, 0.0, 0.0},
		{"One * One", 1.0, 1.0, 1.0},
		{"Positive * Positive", 12.5, 8.0, 100.0},
		{"Negative * Negative", -12.5, -8.0, 100.0},
		{"Positive * Negative", 12.5, -8.0, -100.0},
		{"Negative * Positive", -12.5, 8.0, -100.0},
		{"Decimals", 1.5, 2.5, 3.75},
		{"Small numbers", 0.001, 0.002, 0.000002},
		{"Large numbers", 1000.0, 999.999, 999999.0},
		{"Mixed scales", 12.345, 6.7, 82.7115},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := FromFloat64(tt.d1)
			d2 := FromFloat64(tt.d2)

			// Store original values for verification
			origD1 := d1.Float64()
			origD2 := d2.Float64()

			// Perform multiplication
			result := d1.MulDecimal(d2)

			// Check result
			actual := result.Float64()
			if math.Abs(actual-tt.expected) > 1e-12 {
				t.Errorf("MulDecimal(%f, %f) = %f, expected %f", origD1, origD2, actual, tt.expected)
			}

			// Verify that d1 was not modified
			if d1.Float64() != origD1 {
				t.Errorf("MulDecimal operation modified the first operand: %f became %f", origD1, d1.Float64())
			}

			// Verify that d2 was not modified
			if d2.Float64() != origD2 {
				t.Errorf("MulDecimal operation modified the second operand: %f became %f", origD2, d2.Float64())
			}
		})
	}
}

func TestDecimal_Mul(t *testing.T) {
	tests := []struct {
		name     string
		d1       float64
		d2       float64
		expected float64
	}{
		{"Zero * Zero", 0.0, 0.0, 0.0},
		{"Zero * Positive", 0.0, 123.456, 0.0},
		{"Positive * Zero", 123.456, 0.0, 0.0},
		{"One * One", 1.0, 1.0, 1.0},
		{"Positive * Positive", 12.5, 8.0, 100.0},
		{"Negative * Negative", -12.5, -8.0, 100.0},
		{"Positive * Negative", 12.5, -8.0, -100.0},
		{"Negative * Positive", -12.5, 8.0, -100.0},
		{"Decimals", 1.5, 2.5, 3.75},
		{"Small numbers", 0.001, 0.002, 0.000002},
		{"Large numbers", 1000.0, 999.999, 999999.0},
		{"Mixed scales", 12.345, 6.7, 82.7115},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d1 := FromFloat64(tt.d1)

			// Store original value for verification
			origD1 := d1.Float64()

			// Perform multiplication (this doesn't modify d1)
			actual := d1.Mul(tt.d2)

			// Check result
			if math.Abs(actual-tt.expected) > 1e-12 {
				t.Errorf("Mul(%f, %f) = %f, expected %f", origD1, tt.d2, actual, tt.expected)
			}

			// Verify that d1 was not modified
			if d1.Float64() != origD1 {
				t.Errorf("Mul operation modified the original decimal: %f became %f", origD1, d1.Float64())
			}
		})
	}
}

func TestDecimal_Float64(t *testing.T) {
	tests := []struct {
		name  string
		input float64
	}{
		{"Zero", 0.0},
		{"Positive integer", 123.0},
		{"Negative integer", -456.0},
		{"Positive decimal", 123.456},
		{"Negative decimal", -123.456},
		{"Small decimal", 0.001},
		{"Large number", 999999.999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := FromFloat64(tt.input)
			result := d.Float64()

			if math.Abs(result-tt.input) > 1e-10 {
				t.Errorf("Float64() roundtrip failed: %f -> %f", tt.input, result)
			}
		})
	}
}

// Test edge cases and error conditions
func TestDecimal_EdgeCases(t *testing.T) {
	t.Run("Chain operations", func(t *testing.T) {
		d := FromFloat64(10.0)
		d = d.AddDecimal(FromFloat64(5.0)) // d = 15.0
		d = d.MulDecimal(FromFloat64(2.0)) // d = 30.0

		expected := 30.0
		actual := d.Float64()

		if math.Abs(actual-expected) > 1e-10 {
			t.Errorf("Chain operations failed: expected %f, got %f", expected, actual)
		}
	})

	t.Run("Chain operations with subtraction", func(t *testing.T) {
		d := FromFloat64(100.0)
		d = d.SubDecimal(FromFloat64(30.0)) // d = 70.0
		d = d.AddDecimal(FromFloat64(15.0)) // d = 85.0
		d = d.MulDecimal(FromFloat64(2.0))  // d = 170.0

		expected := 170.0
		actual := d.Float64()

		if math.Abs(actual-expected) > 1e-10 {
			t.Errorf("Chain operations with subtraction failed: expected %f, got %f", expected, actual)
		}
	})

	t.Run("Very large numbers", func(t *testing.T) {
		d1 := FromFloat64(1e15)
		d2 := FromFloat64(1e15)

		result := d1.AddDecimal(d2)
		expected := 2e15
		actual := result.Float64()

		if math.Abs(actual-expected) > 1e5 { // Allow for some floating point error with large numbers
			t.Errorf("Large number addition failed: expected %f, got %f", expected, actual)
		}
	})

	t.Run("Very small numbers", func(t *testing.T) {
		d1 := FromFloat64(1e-15)
		d2 := FromFloat64(1e-15)

		result := d1.AddDecimal(d2)
		expected := 2e-15
		actual := result.Float64()

		if math.Abs(actual-expected) > 1e-20 {
			t.Errorf("Small number addition failed: expected %f, got %f", expected, actual)
		}
	})
}

// Test compatibility with shopspring/decimal directly
func TestShopspringCompatibility(t *testing.T) {
	t.Run("Addition results match shopspring/decimal", func(t *testing.T) {
		// Test that our wrapper produces the same results as direct shopspring usage
		f1, f2 := 123.456, 78.9

		// Our implementation
		d1 := FromFloat64(f1)
		d2 := FromFloat64(f2)
		result := d1.AddDecimal(d2)
		ourResult := result.Float64()

		// Direct shopspring
		sd1 := sdecimal.NewFromFloat(f1)
		sd2 := sdecimal.NewFromFloat(f2)
		shopspringResult, _ := sd1.Add(sd2).Float64()

		if math.Abs(ourResult-shopspringResult) > 1e-15 {
			t.Errorf("Addition results don't match shopspring: our=%f, shopspring=%f", ourResult, shopspringResult)
		}
	})

	t.Run("Subtraction results match shopspring/decimal", func(t *testing.T) {
		// Test that our wrapper produces the same results as direct shopspring usage
		f1, f2 := 123.456, 78.9

		// Our implementation
		d1 := FromFloat64(f1)
		d2 := FromFloat64(f2)
		result := d1.SubDecimal(d2)
		ourResult := result.Float64()

		// Direct shopspring
		sd1 := sdecimal.NewFromFloat(f1)
		sd2 := sdecimal.NewFromFloat(f2)
		shopspringResult, _ := sd1.Sub(sd2).Float64()

		if math.Abs(ourResult-shopspringResult) > 1e-15 {
			t.Errorf("Subtraction results don't match shopspring: our=%f, shopspring=%f", ourResult, shopspringResult)
		}
	})
}
