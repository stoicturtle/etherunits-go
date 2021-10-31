package utils

import (
	"fmt"
	"math/big"
	"strconv"
)

func integerishToBigInt(n interface{}) (*big.Int, bool) {
	switch val := n.(type) {
	case int, int8, int16, int32, int64:
		x, err := strconv.ParseInt(fmt.Sprintf("%v", val), 10, 64)
		if err != nil {
			panic(err)
		}
		return big.NewInt(x), true
	case uint, uint8, uint16, uint32, uint64:
		x, err := strconv.ParseUint(fmt.Sprintf("%v", val), 10, 64)
		if err != nil {
			panic(err)
		}
		return new(big.Int).SetUint64(x), true
	default:
		return nil, false
	}
}

func BigIntishToBigInt(n interface{}) (*big.Int, bool) {
	switch val := n.(type) {
	case *big.Int:
		return val, true
	case string:
		return new(big.Int).SetString(val, 10)
	default:
		return integerishToBigInt(val)
	}
}

func BigFloatishToBigFloat(n interface{}) (*big.Float, bool) {
	switch val := n.(type) {
	case *big.Float:
		return val, true
	case float32:
		return BigFloatishToBigFloat(float64(val))
	case float64:
		return big.NewFloat(val), true
	case string:
		v, ok := new(big.Float).SetString(val)
		if !ok {
			fmt.Println(fmt.Errorf("cannot convert string %s to *big.Float", val))
			return nil, false
		}

		return v, ok
	default:
		return nil, false
	}
}