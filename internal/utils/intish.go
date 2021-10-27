package utils

import (
	"fmt"
	"math/big"
	"reflect"
	"strconv"
)

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

func strToBigInt(str string) (*big.Int, error) {
	v, ok := new(big.Int).SetString(str, 10)
	if !ok {
		return nil, fmt.Errorf("cannot convert string %s to *big.Int", str)
	}

	return v, nil
}
