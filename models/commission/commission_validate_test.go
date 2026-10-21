package commission

import (
	"testing"

	"buf.build/go/protovalidate"
	"google.golang.org/genproto/googleapis/type/decimal"
)

func TestCommissionSettingsProtovalidate(t *testing.T) {
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("protovalidate.New: %v", err)
	}

	tests := []struct {
		name    string
		msg     *CommissionSettings
		wantErr bool
	}{
		{
			name: "valid notional 2.50",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "2.50"},
				CommissionType: CommissionType_NOTIONAL,
			},
		},
		{
			name: "valid qty 0.05",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "0.05"},
				CommissionType: CommissionType_QTY,
			},
		},
		{
			name: "valid bps 12.50",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "12.50"},
				CommissionType: CommissionType_BPS,
			},
		},
		{
			name: "valid max 10000 integer bps",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "10000"},
				CommissionType: CommissionType_BPS,
			},
		},
		{
			name: "valid max 10000.00 notional",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "10000.00"},
				CommissionType: CommissionType_NOTIONAL,
			},
		},
		{
			name: "valid max qty 100",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "100"},
				CommissionType: CommissionType_QTY,
			},
		},
		{
			name: "valid 100.01 notional",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "100.01"},
				CommissionType: CommissionType_NOTIONAL,
			},
		},
		{
			name: "valid zero bps",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "0"},
				CommissionType: CommissionType_BPS,
			},
		},
		{
			name: "missing commission",
			msg: &CommissionSettings{
				CommissionType: CommissionType_BPS,
			},
			wantErr: true,
		},
		{
			name: "unused commission type",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "25"},
				CommissionType: CommissionType_NOT_USED_COMMISSION_TYPE,
			},
			wantErr: true,
		},
		{
			name: "negative notional",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "-1"},
				CommissionType: CommissionType_NOTIONAL,
			},
			wantErr: true,
		},
		{
			name: "negative qty",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "-0.05"},
				CommissionType: CommissionType_QTY,
			},
			wantErr: true,
		},
		{
			name: "negative bps",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "-1"},
				CommissionType: CommissionType_BPS,
			},
			wantErr: true,
		},
		{
			name: "bps more than 2 fraction digits",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "0.001"},
				CommissionType: CommissionType_BPS,
			},
			wantErr: true,
		},
		{
			name: "notional more than 2 fraction digits",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "0.001"},
				CommissionType: CommissionType_NOTIONAL,
			},
			wantErr: true,
		},
		{
			name: "qty above 100",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "100.01"},
				CommissionType: CommissionType_QTY,
			},
			wantErr: true,
		},
		{
			name: "bps above 10000",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "20000"},
				CommissionType: CommissionType_BPS,
			},
			wantErr: true,
		},
		{
			name: "zero value but invalid fraction digits",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "0.000"},
				CommissionType: CommissionType_NOTIONAL,
			},
			wantErr: true,
		},
		{
			name: "exponent form rejected",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: "1e2"},
				CommissionType: CommissionType_QTY,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() err=%v, wantErr=%v", err, tt.wantErr)
			}
		})
	}
}
