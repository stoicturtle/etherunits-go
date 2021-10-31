package utils

import (
	"fmt"
	"math/big"
)

func isFloatish(n interface{}) bool {
	switch n.(type) {
	case float32, float64:
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
	case float64:
		x = val
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

func strToBigFloat(str string) (*big.Float, error) {
	v, ok := new(big.Float).SetString(str)
	if !ok {
		return nil, fmt.Errorf("cannot convert string %s to *big.Float", str)
	}

	return v, nil
}