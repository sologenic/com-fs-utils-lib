package commission

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/sologenic/com-fs-utils-lib/models/decimal"
)

func TestCommissionSettingsProtovalidate(t *testing.T) {
	v, err := protovalidate.New()
	if err != nil {
		t.Fatalf("protovalidate.New: %v", err)
	}

	bps := CommissionType_BPS
	notional := CommissionType_NOTIONAL
	qty := CommissionType_QTY
	unused := CommissionType_NOT_USED_COMMISSION_TYPE

	tests := []struct {
		name    string
		msg     *CommissionSettings
		wantErr bool
	}{
		{
			name: "valid notional 2.50",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 250, Exp: -2},
				CommissionType: &notional,
			},
		},
		{
			name: "valid qty 0.05",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 5, Exp: -2},
				CommissionType: &qty,
			},
		},
		{
			name: "valid bps 12.50",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 1250, Exp: -2},
				CommissionType: &bps,
			},
		},
		{
			name: "valid max 10000 integer",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 10000, Exp: 0},
				CommissionType: &bps,
			},
		},
		{
			name: "valid max 10000.00",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 1000000, Exp: -2},
				CommissionType: &bps,
			},
		},
		{
			name: "valid max from float encoding Value=1 Exp=4",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 1, Exp: 4},
				CommissionType: &bps,
			},
		},
		{
			name: "valid 100.01",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 10001, Exp: -2},
				CommissionType: &bps,
			},
		},
		{
			name: "valid zero",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 0, Exp: 0},
				CommissionType: &bps,
			},
		},
		{
			name: "missing commission",
			msg: &CommissionSettings{
				CommissionType: &bps,
			},
			wantErr: true,
		},
		{
			name: "missing commission type",
			msg: &CommissionSettings{
				Commission: &decimal.Decimal{Value: 25, Exp: 0},
			},
			wantErr: true,
		},
		{
			name: "unused commission type",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 25, Exp: 0},
				CommissionType: &unused,
			},
			wantErr: true,
		},
		{
			name: "negative commission",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: -1, Exp: 0},
				CommissionType: &bps,
			},
			wantErr: true,
		},
		{
			name: "more than 2 fraction digits",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 1, Exp: -3},
				CommissionType: &bps,
			},
			wantErr: true,
		},
		{
			name: "above 10000 with fraction digits",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 1000001, Exp: -2},
				CommissionType: &bps,
			},
			wantErr: true,
		},
		{
			name: "above 10000 via positive exp",
			msg: &CommissionSettings{
				Commission:     &decimal.Decimal{Value: 2, Exp: 4},
				CommissionType: &bps,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.msg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("Validate() err=%v wantErr=%v", err, tt.wantErr)
			}
		})
	}
}
