package etherunits

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/stoicturtle/etherunits-go/internal/unit"
	"github.com/stoicturtle/etherunits-go/internal/utils"
)

// ToWei converts any one of (*)(u)int/8/16/32/64, (*)float32/64, *big.Int, *big.Float, or string
// from its value in the denomination of fromUnit to its denominational value in Wei.
func ToWei(value interface{}, fromUnit Unit) (*big.Int, error) {
	if !unit.ValidUnit(fromUnit) {
		return nil, fmt.Errorf("ToWei(): invalid unit %s", fromUnit.String())
	}

	// it's easier to just convert whatever value into a denomination of Ether
	// and then just convert that value to Wei than it is to duplicate the same damn
	// math anyway.
	ethVal, err := ToEther(value, fromUnit)
	if err != nil {
		err = fmt.Errorf(strings.Replace(err.Error(), "ToEther", "ToWei", 1))
		return nil, err
	}

	toEther := Wei.ValueEther()
	weiVal, _ := ethVal.
		SetPrec(toEther.Prec()).
		Quo(ethVal, toEther).
		Int(nil)

	return weiVal, nil
}

// ToEther converts any one of (*)(u)int/8/16/32/64, (*)float32/64, *big.Int, *big.Float, or string
// from its value in the denomination of fromUnit to its denominational value in Ether.
func ToEther(value interface{}, fromUnit Unit) (*big.Float, error) {
	if !unit.ValidUnit(fromUnit) {
		return nil, fmt.Errorf("ToEther(): invalid unit %s", fromUnit.String())
	}

	multiplier := fromUnit.ValueEther()

	var f *big.Float
	if tryBigInt, isBigInt := utils.BigIntishToBigInt(value); isBigInt {
		f = new(big.Float).SetInt(tryBigInt)
	} else if tryBigFloat, isBigFloat := utils.BigFloatishToBigFloat(value); isBigFloat {
		f = tryBigFloat
	} else {
		return nil, fmt.Errorf("ToEther(): unable to parse %v into a *big.Int or *big.Float", value)
	}

	f.SetPrec(multiplier.Prec())
	if fromUnit == Ether {
		return f, nil
	}

	return f.Mul(f, multiplier), nil
}