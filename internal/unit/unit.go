package unit

import (
	"fmt"
	"math"
	"math/big"
)

//go:generate stringer -type=Unit

// Unit represents a denominational value of currency on Ethereum(-based) blockchains.
type Unit uint8

func (u Unit) ValueWei() *big.Int {
	value, ok := ValueMap[u]
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

func (u Unit) valid() bool {
	for _, unit := range Slice {
		if u == unit {
			return true
		}
	}

	return false
}

func (u Unit) etherExponent() int {
	return ethExponentMap[u]
}

const (
	Wei Unit = iota
	KWei
	MWei
	GWei
	Szabo
	Finney
	Ether
	KEther
	MEther
	GEther
	TEther
	Max
)

func ValidUnit(u Unit) bool {
	return u.valid()
}