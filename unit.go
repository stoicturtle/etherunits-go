package etherunits

import (
	"fmt"
	"math"
	"math/big"
)

//go:generate stringer -type=Unit

// Unit represents a denominational value of currency on Ethereum(-based) blockchains.
type Unit uint8

func (u Unit) ValueWei() *big.Int {
	value, ok := unitValueMap[u]
	if !ok {
		panic(fmt.Errorf("Unit.ValueWei(): no Wei value for Unit %s", u.String()))
	}

	return value
}

func (u Unit) ValueEther() *big.Float {
	if u == Ether {
		return big.NewFloat(1)
	}

	val := math.Pow10(u.etherExponent())

	return big.NewFloat(val)
}

func (u Unit) etherExponent() int {
	return unitEthExpMap[u]
}

func (u Unit) valid() bool {
	for _, unit := range unitsSlice {
		if u == unit {
			return true
		}
	}

	return false
}

const (
	// Wei represents the base unit of measurement for EVM units. 1 wei is equal to 10^(-18) ether.
	// All values on Ethereum(-based) blockchains are stored and operated on in values of wei.
	Wei Unit = iota
	// KWei is equal to 10^(-15) ether, or 10^3 -- 1,000 -- wei.
	KWei
	// MWei is equal to 10^(-12) ether, or 10^6 -- 1,000,000 (one million) -- wei.
	MWei
	// GWei is equal to 10^(-9) ether, or 10^9 -- 1,000,000,000 (one billion) -- wei.
	// As a unit of measurement, it is often used for representing the gas costs of Ethereum transactions.
	GWei
	// Szabo is equal to 10^(-6) ether, or 10^12 -- 1,000,000,000,000 (one trillion) -- wei.
	// Note that it is not a commonly used unit of measurement.
	Szabo
	// Finney is equal to 10^(-3) ether, or 10^15 -- 1,000,000,000,000,000 (one quadrillion) -- wei.
	// Note that it is not a commonly used unit of measurement.
	Finney
	// Ether is likely the most commonly used unit of measurement for values of transactions, tokens, etc.,
	// on Ethereum(-based) blockchains, commonly in user-facing contexts.
	// 1 ether is equal to 10^18 -- 1,000,000,000,000,000,000 (one quintillion) -- wei
	Ether
	// KEther is equal to 10^3 -- 1,000 -- ether.
	KEther
	// MEther is equal to 10^6 -- 1,000,000 (one million) -- ether.
	MEther
	// GEther is equal to 10^9 -- 1,000,000,000 (one billion) -- ether.
	GEther
	// TEther is equal to 10^12 -- 1,000,000,000,000 (one trillion) -- ether.
	TEther

	_max
)
