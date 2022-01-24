package etherunits_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/math"

	"github.com/stoicturtle/etherunits-go"
)

func TestFromWei(t *testing.T) {
	type args struct {
		value  interface{}
		toUnit etherunits.Unit
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Float
		wantErr bool
	}{
		{
			"test to ether",
			args{math.MustParseBig256("450000000000000000000"), etherunits.Ether},
			big.NewFloat(450),
			false,
		},
		{
			"test to gwei",
			args{math.MustParseBig256("451000000000000000000"), etherunits.GWei},
			big.NewFloat(451000000000),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := etherunits.FromWei(tt.args.value, tt.args.toUnit)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromWei() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.Cmp(got) != 0 {
				t.Errorf("FromWei() got = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}

func TestToWei(t *testing.T) {
	type args struct {
		value    interface{}
		fromUnit etherunits.Unit
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			"test from ether",
			args{450.554, etherunits.Ether},
			math.MustParseBig256("450554000000000000000"),
			false,
		},
		{
			"test from gwei",
			args{big.NewInt(451000000000), etherunits.GWei},
			math.MustParseBig256("451000000000000000000"),
			false,
		},
		{
			"test from tether",
			args{0.000000000450554, etherunits.TEther},
			math.MustParseBig256("450554000000000000000"),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := etherunits.ToWei(tt.args.value, tt.args.fromUnit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToWei() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want.Cmp(got) != 0 {
				t.Errorf("ToWei() got = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}