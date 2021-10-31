package etherunits

import (
	"math/big"
)

// MaxUint256 is the maximum value of Solidity's uint256 type,
// and is equal to 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff.
var MaxUint256 = new(big.Int).Sub(
	new(big.Int).Exp(bigInt10, maxUintExp, nil),
	big.NewInt(1),
)

var (
	bigInt10   = big.NewInt(10)
	maxUintExp = big.NewInt(256)
)