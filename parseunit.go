package etherunits

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/stoicturtle/etherunits-go/internal/unit"
	"github.com/stoicturtle/etherunits-go/internal/utils"
)

// ParseUnit wraps ParseUnitFromString, ParseUnitFromBigInt, and ParseUnitFromBigFloat.
// If the passed parameter is an (u)int/8/16/32/64 or float32/64, it is converted
// to a respective big.Int or big.Float before being passed to its respective function.
// This function will return any errors returned by wrapped functions, or an error
// if the passed parameter is not a valid type to parse.
func ParseUnit(unit interface{}) (Unit, error) {
	if str, ok := unit.(string); ok {
		return ParseUnitFromString(str)
	}
	if bigint, ok := utils.BigIntishToBigInt(unit); ok {
		return ParseUnitFromBigInt(bigint)
	}
	if bigfloat, ok := utils.BigFloatishToBigFloat(unit); ok {
		return ParseUnitFromBigFloat(bigfloat)
	}

	return _max + 1, fmt.Errorf("ParseUnit(): invalid type %T for unit", unit)
}

// ParseUnitFromString returns the Unit corresponding to the passed unit name string.
// This function specifically downcases both the parameter and checked Unit names,
// so there's no need to worry about proper casing.
func ParseUnitFromString(u string) (Unit, error) {
	lower := strings.ToLower
	for _, u2 := range unit.Slice {
		if lower(u) == lower(u2.String()) {
			return u2, nil
		}
	}

	return _max + 1, fmt.Errorf("ParseUnitFromString(): unknown unit name %s", u)
}

// ParseUnitFromBigInt returns the Unit corresponding to the passed *big.Int value.
// If the passed *big.Int does not correspond to a known Unit value, an invalid Unit iota
// is returned along with an error.
func ParseUnitFromBigInt(u *big.Int) (Unit, error) {
	for _, u2 := range unit.Slice {
		if utils.BigIntEq(
			u,
			u2.ValueWei(),
		) {
			return u2, nil
		}
	}

	return _max + 1, fmt.Errorf("ParseUnitFromBigInt(): unit value %s does not correspond to any known units", u.String())
}

// ParseUnitFromBigFloat returns the Unit corresponding to the passed *big.Float value.
// If the passed *big.Float does not correspond to a known Unit value, an invalid Unit iota
// is returned along with an error.
func ParseUnitFromBigFloat(u *big.Float) (Unit, error) {
	for _, u2 := range unit.Slice {
		weiVal := new(big.Float).SetInt(u2.ValueWei())

		if utils.BigFloatEq(u, weiVal) {
			return u2, nil
		}
	}

	return _max + 1, fmt.Errorf("ParseUnitFromBigFloat(): unit value %s does not correspond to any known units", u.String())
}