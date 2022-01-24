package etherunits

import (
	"fmt"
	"math/big"
)

// Unit represents a denominational value of currency on Ethereum(-based) blockchains.
type Unit string

func UnitFromDecimals(decimals uint8) Unit {
	var unit Unit

	for u := range unitValueMap {
		if 18-u.baseLength() == int(decimals) {
			unit = u
		}
	}

	if unit == "" {
		panic(fmt.Errorf("unable to find a unit from decimal value %[1]d", decimals))
	}

	return unit
}

func (u Unit) base() *big.Int {
	val, ok := unitValueMap[u]
	if !ok {
		panic(fmt.Errorf("unknown unit %[1]s", u))
	}

	base, ok := new(big.Int).SetString(val, 10)
	if !ok {
		panic(fmt.Errorf("unable to parse %[1]s to big.Int", val))
	}

	return base
}

func (u Unit) baseLength() int {
	val, ok := unitValueMap[u]
	if !ok {
		panic(fmt.Errorf("unknown unit %[1]s", u))
	}

	return len(val) - 1
}

const (
	// Wei represents the base unit of measurement for EVM units. 1 wei is equal to 10^(-18) ether.
	// All values on Ethereum(-based) blockchains are stored and operated on in values of wei.
	Wei Unit = "wei"
	// KWei is equal to 10^(-15) ether, or 10^3 -- 1,000 -- wei.
	KWei Unit = "kwei"
	// MWei is equal to 10^(-12) ether, or 10^6 -- 1,000,000 (one million) -- wei.
	MWei Unit = "mwei"
	// GWei is equal to 10^(-9) ether, or 10^9 -- 1,000,000,000 (one billion) -- wei.
	// As a unit of measurement, it is often used for representing the gas costs of Ethereum transactions.
	GWei Unit = "gwei"
	// Szabo is equal to 10^(-6) ether, or 10^12 -- 1,000,000,000,000 (one trillion) -- wei.
	// Note that it is not a commonly used unit of measurement.
	Szabo Unit = "szabo"
	// Finney is equal to 10^(-3) ether, or 10^15 -- 1,000,000,000,000,000 (one quadrillion) -- wei.
	// Note that it is not a commonly used unit of measurement.
	Finney Unit = "finney"
	// Ether is likely the most commonly used unit of measurement for values of transactions, tokens, etc.,
	// on Ethereum(-based) blockchains, commonly in user-facing contexts.
	// 1 ether is equal to 10^18 -- 1,000,000,000,000,000,000 (one quintillion) -- wei
	Ether Unit = "ether"
	// KEther is equal to 10^3 -- 1,000 -- ether.
	KEther Unit = "kether"
	// MEther is equal to 10^6 -- 1,000,000 (one million) -- ether.
	MEther Unit = "mether"
	// GEther is equal to 10^9 -- 1,000,000,000 (one billion) -- ether.
	GEther Unit = "gether"
	// TEther is equal to 10^12 -- 1,000,000,000,000 (one trillion) -- ether.
	TEther Unit = "tether"
)

var unitValueMap = map[Unit]string{
	Wei:    "1",
	KWei:   "1000",
	MWei:   "1000000",
	GWei:   "1000000000",
	Szabo:  "1000000000000",
	Finney: "1000000000000000",
	Ether:  "1000000000000000000",
	KEther: "1000000000000000000000",
	MEther: "1000000000000000000000000",
	GEther: "1000000000000000000000000000",
	TEther: "1000000000000000000000000000000",
}