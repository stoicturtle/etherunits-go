package etherunits_test

import (
	"math/big"
	"testing"

	"github.com/stoicturtle/etherunits-go"
)

var oneEtherWei *big.Int

func init() {
	var ok bool
	oneEtherWei, ok = new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		panic("wtf?")
	}
}

func TestParseUnit(t *testing.T) {
	type args struct {
		unit interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    etherunits.Unit
		wantErr bool
	}{
		{
			"wei from string",
			args{"wei"},
			etherunits.Wei,
			false,
		},
		{
			"ether from big.Int value",
			args{oneEtherWei},
			etherunits.Ether,
			false,
		},
		{
			"parsing unknown unit name fails",
			args{"SuperWei"},
			etherunits.Unit(11 + 1),
			true,
		},
		{
			"parsing unknown big.Int unit value fails",
			args{new(big.Int).Add(oneEtherWei, big.NewInt(1))},
			etherunits.Unit(11 + 1),
			true,
		},
		{
			"parsing unknown big.Float unit value fails",
			args{new(big.Float).SetInt(new(big.Int).Add(oneEtherWei, big.NewInt(1)))},
			etherunits.Unit(11 + 1),
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := etherunits.ParseUnit(tt.args.unit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUnit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseUnit() got = %v, want %v", got, tt.want)
			}
		})
	}
}
