package etherunits

import (
	"github.com/stoicturtle/etherunits-go/internal/unit"
)

// Unit represents a denominational value of currency on Ethereum(-based) blockchains.
type Unit = unit.Unit

const (
	// Wei represents the base unit of measurement for EVM units. 1 wei is equal to 10^(-18) ether.
	// All values on Ethereum(-based) blockchains are stored and operated on in values of wei.
	Wei = unit.Wei
	// KWei is equal to 10^(-15) ether, or 10^3 -- 1,000 -- wei.
	KWei = unit.KWei
	// MWei is equal to 10^(-12) ether, or 10^6 -- 1,000,000 (one million) -- wei.
	MWei = unit.MWei
	// GWei is equal to 10^(-9) ether, or 10^9 -- 1,000,000,000 (one billion) -- wei.
	// As a unit of measurement, it is often used for representing the gas costs of Ethereum transactions.
	GWei = unit.GWei
	// Szabo is equal to 10^(-6) ether, or 10^12 -- 1,000,000,000,000 (one trillion) -- wei.
	// Note that it is not a commonly used unit of measurement.
	Szabo = unit.Szabo
	// Finney is equal to 10^(-3) ether, or 10^15 -- 1,000,000,000,000,000 (one quadrillion) -- wei.
	// Note that it is not a commonly used unit of measurement.
	Finney = unit.Finney
	// Ether is likely the most commonly used unit of measurement for values of transactions, tokens, etc.,
	// on Ethereum(-based) blockchains, commonly in user-facing contexts.
	// 1 ether is equal to 10^18 -- 1,000,000,000,000,000,000 (one quintillion) -- wei
	Ether = unit.Ether
	// KEther is equal to 10^3 -- 1,000 -- ether.
	KEther = unit.KEther
	// MEther is equal to 10^6 -- 1,000,000 (one million) -- ether.
	MEther = unit.MEther
	// GEther is equal to 10^9 -- 1,000,000,000 (one billion) -- ether.
	GEther = unit.GEther
	// TEther is equal to 10^12 -- 1,000,000,000,000 (one trillion) -- ether.
	TEther = unit.TEther

	_max = unit.Max
)