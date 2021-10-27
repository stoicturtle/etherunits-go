package etherunits_test

import (
	"fmt"
	"math"
	"math/big"
	"testing"

	"github.com/stoicturtle/etherunits-go"
)

const someWeiValue = "5332160000000000000"

var (
	someWeiValueBigint   = big.NewInt(5332160000000000000)
	someWeiValueBigfloat = new(big.Float).SetInt(someWeiValueBigint)
)

type TestArgs struct {
	value    interface{}
	fromUnit etherunits.Unit
}

type TestCase struct {
	name    string
	args    TestArgs
	want    fmt.Stringer
	wantErr bool
}

func TestToEther(t *testing.T) {
	t.Parallel()

	getWant := func(wantUnit etherunits.Unit, val float64) *big.Float {
		return new(big.Float).Mul(wantUnit.ValueEther(), big.NewFloat(val))
	}

	var (
		testFloatPtr float64 = 17.2
		testIntPtr   int64   = 17
	)

	maxuint256Float := new(big.Float).SetInt(etherunits.MaxUint256)
	maxuint256Want := new(big.Float).Mul(
		big.NewFloat(math.Pow10(-18)),
		maxuint256Float,
	)

	tests := []TestCase{
		{
			"from wei str",
			TestArgs{someWeiValue, etherunits.Wei},
			getWant(etherunits.Ether, 5.33216),
			false,
		},
		{
			"from wei bigint",
			TestArgs{someWeiValueBigint, etherunits.Wei},
			getWant(etherunits.Ether, 5.33216),
			false,
		},
		{
			"from wei bigfloat",
			TestArgs{someWeiValueBigfloat, etherunits.Wei},
			getWant(etherunits.Ether, 5.33216),
			false,
		},
		{
			"from gwei",
			TestArgs{44.5, etherunits.GWei},
			getWant(etherunits.GWei, 44.5),
			false,
		},
		{
			"from tether whole number",
			TestArgs{12, etherunits.TEther},
			getWant(etherunits.TEther, 12),
			false,
		},
		{
			"from tether decimal number",
			TestArgs{13.5221, etherunits.TEther},
			getWant(etherunits.TEther, 13.5221),
			false,
		},
		{
			"from mether float pointer",
			TestArgs{&testFloatPtr, etherunits.MEther},
			getWant(etherunits.MEther, testFloatPtr),
			false,
		},
		{
			"from mether int pointer",
			TestArgs{&testIntPtr, etherunits.MEther},
			getWant(etherunits.MEther, float64(testIntPtr)),
			false,
		},
		{
			"test MaxUint256",
			TestArgs{etherunits.MaxUint256, etherunits.Wei},
			maxuint256Want,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := etherunits.ToEther(tt.args.value, tt.args.fromUnit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToEther() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want.String() {
				t.Errorf("ToEther() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToWei(t *testing.T) {
	t.Parallel()

	getWant := func(wantUnit etherunits.Unit, val float64) *big.Int {
		f := big.NewFloat(val)
		n := new(big.Float).SetInt(wantUnit.ValueWei())

		res := new(big.Int)
		_, _ = new(big.Float).Mul(f, n).Int(res)

		return res
	}

	tests := []TestCase{
		{
			"from ether str",
			TestArgs{"5.33216", etherunits.Ether},
			someWeiValueBigint,
			false,
		},
		{
			"from ether bigfloat",
			TestArgs{big.NewFloat(5.33216), etherunits.Ether},
			someWeiValueBigint,
			false,
		},
		{
			"from gwei whole number",
			TestArgs{44, etherunits.GWei},
			getWant(etherunits.GWei, 44),
			false,
		},
		{
			"from gwei decimal number",
			TestArgs{45.21, etherunits.GWei},
			getWant(etherunits.GWei, 45.21),
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got, err := etherunits.ToWei(tt.args.value, tt.args.fromUnit)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToWei() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got.String() != tt.want.String() {
				t.Errorf("ToWei() got = %v, want %v", got, tt.want)
			}
		})
	}
}
