package utils

import (
	"fmt"
	"math"
	"math/big"
	"reflect"
	"strconv"
)

var (
	wei   = big.NewInt(1)
	ether = MulWei(18)
)

func bigFloatToBigInt(f *big.Float) *big.Int {
	b := new(big.Int)
	_, _ = f.Int(b)

	return b
}

func MulWei(pow10Exp int) *big.Int {
	f := big.NewFloat(math.Pow10(pow10Exp))
	val := bigFloatToBigInt(f)

	return new(big.Int).Mul(wei, val)
}

func MulEth(pow10Exp int) *big.Int {
	f := big.NewFloat(math.Pow10(pow10Exp))
	val := bigFloatToBigInt(f)

	return new(big.Int).Mul(ether, val)
}

func BigIntEq(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

func BigFloatEq(x, y *big.Float) bool {
	return x.Cmp(y) == 0
}

func isIntish(n interface{}) bool {
	switch n.(type) {
	case int, int8, int16, int32, int64:
		return true
	case *int, *int8, *int16, *int32, *int64:
		return true
	default:
		return false
	}
}

func makeFormatStr(n interface{}) string {
	var n64Str string

	reflectVal := reflect.ValueOf(n)
	if reflectVal.Kind() == reflect.Ptr {
		n64Str = fmt.Sprintf("%v", reflect.Indirect(reflectVal))
	} else {
		n64Str = fmt.Sprintf("%v", n)
	}

	return n64Str
}

func intishToBigInt(n interface{}) *big.Int {
	x, err := strconv.ParseInt(makeFormatStr(n), 10, 64)
	if err != nil {
		panic(err)
	}

	return big.NewInt(x)
}

func isUintish(n interface{}) bool {
	switch n.(type) {
	case uint, uint8, uint16, uint32, uint64:
		return true
	case *uint, *uint8, *uint16, *uint32, *uint64:
		return true
	default:
		return false
	}
}

func uintishToBigInt(n interface{}) *big.Int {
	x, err := strconv.ParseUint(makeFormatStr(n), 10, 64)
	if err != nil {
		panic(err)
	}

	return new(big.Int).SetUint64(x)
}

func BigIntishToBigInt(n interface{}) (*big.Int, bool) {
	val, ok := n.(*big.Int)
	if ok {
		return val, true
	}

	if isIntish(n) {
		return intishToBigInt(n), true
	}

	if isUintish(n) {
		return uintishToBigInt(n), true
	}

	str, ok := n.(string)
	if ok {
		val, err := strToBigInt(str)
		return val, err == nil
	}

	return nil, false
}

func isFloatish(n interface{}) bool {
	switch n.(type) {
	case float32, float64:
		return true
	case *float32, *float64:
		return true
	default:
		return false
	}
}

func floatishToBigFloat(n interface{}) *big.Float {
	var x float64

	switch val := n.(type) {
	case float32:
		x = float64(val)
	case *float32:
		x = float64(*val)
	case float64:
		x = val
	case *float64:
		x = *val
	}

	return big.NewFloat(x)
}

func BigFloatishToBigFloat(n interface{}) (*big.Float, bool) {
	if val, ok := n.(*big.Float); ok {
		return val, true
	}

	if isFloatish(n) {
		return floatishToBigFloat(n), true
	}

	if str, ok := n.(string); ok {
		val, err := strToBigFloat(str)
		return val, err == nil
	}

	return nil, false
}

func strToBigInt(str string) (*big.Int, error) {
	v, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return nil, fmt.Errorf("cannot convert string %s to *big.Int", str)
	}

	return v, nil
}

func strToBigFloat(str string) (*big.Float, error) {
	v, ok := new(big.Float).SetString(str)
	if !ok {
		return nil, fmt.Errorf("cannot convert string %s to *big.Float", str)
	}

	return v, nil
}
