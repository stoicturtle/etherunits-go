package etherunits

import (
	"fmt"
	"math/big"
	"reflect"
	"strings"

	"github.com/ethereum/go-ethereum/common/math"
)

var (
	zero   = big.NewInt(0)
	zeroBF = big.NewFloat(0)
)

// ToWei converts any one of (u)int/8/16/32/64, float32/64, *big.Int, *big.Float, or string
// from its value in the denomination of fromUnit to its denominational value in Wei.
func ToWei(value interface{}, fromUnit Unit) (*big.Int, error) {
	val, err := makeBigFloat(value)
	if err != nil {
		return nil, err
	}

	valCmpZero := val.Cmp(zeroBF)
	if valCmpZero == 0 {
		return zero, nil
	}

	isNeg := valCmpZero == -1

	base := fromUnit.base()
	baseLength := fromUnit.baseLength()

	valStr := val.Text('f', -18)

	var (
		whole = "0"
		frac  = "0"
	)

	if strings.Contains(valStr, ".") {
		split := strings.Split(valStr, ".")
		if split[0] != "" {
			whole = split[0]
		}

		if split[1] != "" {
			frac = split[1]
		}
	} else {
		whole = valStr
	}

	for len(frac) < baseLength {
		frac += "0"
	}

	wholeBN, ok := new(big.Int).SetString(whole, 10)
	if !ok {
		return nil, fmt.Errorf("could not parse whole number string %[1]s into big.Int", whole)
	}

	fracBN, ok := new(big.Int).SetString(frac, 10)
	if !ok {
		return nil, fmt.Errorf("could not parse fractional number string %[1]s into big.Int", frac)
	}

	wei := base.Mul(base, wholeBN)
	wei = wei.Add(wei, fracBN)

	if isNeg {
		wei = wei.Neg(wei)
	}

	return wei, nil
}

func FromWei(value interface{}, toUnit Unit) (*big.Float, error) {
	val, err := makeBigInt(value)
	if err != nil {
		return nil, err
	}

	valCmpZero := val.Cmp(zero)
	if valCmpZero == 0 {
		return zeroBF, nil
	}

	isNeg := valCmpZero == -1

	base := toUnit.base()
	baseLength := toUnit.baseLength()

	frac := new(big.Int).Mod(val, big.NewInt(int64(baseLength))).String()

	for len(frac) < baseLength {
		frac = "0" + frac
	}

	weiStr := val.Div(val, base).String()
	if frac != "0" {
		weiStr += "." + frac
	}

	if isNeg {
		weiStr = "-" + weiStr
	}

	ether, ok := new(big.Float).SetString(weiStr)
	if !ok {
		return nil, fmt.Errorf("unable to parse value %[1]s into big.Float", weiStr)
	}

	return ether, nil
}

// ToEther converts any one of (u)int/8/16/32/64, float32/64, *big.Int, *big.Float, or string
// from its value in the denomination of fromUnit to its denominational value in Ether.
// func ToEther(value interface{}, fromUnit Unit) (*big.Float, error) {
//
// }

func makeBigInt(val interface{}) (*big.Int, error) {
	switch n := val.(type) {
	case *big.Int:
		return n, nil
	case string:
		v, ok := math.ParseBig256(n)
		if !ok {
			return nil, fmt.Errorf("unable to parse string '%[1]s' into big.Int", n)
		}

		return v, nil
	}

	valOf := reflect.ValueOf(val)
	valType := valOf.Type()

	switch {
	case valType.ConvertibleTo(typeOfUint64):
		if valType.ConvertibleTo(typeOfInt64) {
			n := valOf.Int()
			return big.NewInt(n), nil
		}

		n := valOf.Uint()
		return new(big.Int).SetUint64(n), nil
	case valType.ConvertibleTo(typeOfInt64):
		n := valOf.Int()
		return big.NewInt(n), nil
	}

	return nil, fmt.Errorf("unable to convert value of type %T to big.Int", val)
}

func makeBigFloat(val interface{}) (*big.Float, error) {
	f := new(big.Float)
	// f = f.SetPrec(big.MaxPrec)
	// f = f.SetMode(big.ToNearestAway)

	switch n := val.(type) {
	case float32:
		return f.SetFloat64(float64(n)), nil
	case float64:
		return f.SetFloat64(n), nil
	case *big.Float:
		// n = n.SetPrec(big.MaxPrec)
		// n = n.SetMode(big.ToNearestAway)
		return n, nil
	}

	n, err := makeBigInt(val)
	if err != nil {
		return nil, err
	}

	f = f.SetInt(n)

	return f, nil
}

var (
	typeOfInt64  = reflect.TypeOf(int64(1))
	typeOfUint64 = reflect.TypeOf(uint64(1))
)