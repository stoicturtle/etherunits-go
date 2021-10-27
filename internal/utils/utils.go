package utils

import (
	"math"
	"math/big"
)

var (
	wei   = big.NewInt(1)
	ether = MulWei(18)
)

func bigFloatToBigInt(f *big.Float) *big.Int {
	b := new(big.Int)
	_, _ = f.Int(b)

	return b
}

func MulWei(pow10Exp int) *big.Int {
	f := big.NewFloat(math.Pow10(pow10Exp))
	val := bigFloatToBigInt(f)

	return new(big.Int).Mul(wei, val)
}

func MulEth(pow10Exp int) *big.Int {
	f := big.NewFloat(math.Pow10(pow10Exp))
	val := bigFloatToBigInt(f)

	return new(big.Int).Mul(ether, val)
}

func BigIntEq(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

func BigFloatEq(x, y *big.Float) bool {
	return x.Cmp(y) == 0
}
