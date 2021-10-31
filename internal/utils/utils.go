package utils

import (
	"fmt"
	"math"
	"math/big"
)

type ComparableBigNum interface {
	fmt.Stringer
	Cmp(*big.Int) int
}

var (
	wei   = big.NewInt(1)
	ether = MulWei(18)
)

func newBigInt() *big.Int {
	return new(big.Int)
}

func MulBigInt(x, y *big.Int) *big.Int {
	return newBigInt().Mul(x, y)
}

func bigFloatToBigInt(f *big.Float) *big.Int {
	b := new(big.Int)
	_, _ = f.Int(b)

	return b
}

func MulWei(pow10Exp int) *big.Int {
	f := big.NewFloat(math.Pow10(pow10Exp))
	val := bigFloatToBigInt(f)

	return MulBigInt(wei, val)
}

func MulEth(pow10Exp int) *big.Int {
	f := big.NewFloat(math.Pow10(pow10Exp))
	val := bigFloatToBigInt(f)

	return MulBigInt(ether, val)
}

func BigIntEq(x, y *big.Int) bool {
	return x.Cmp(y) == 0
}

func BigFloatEq(x, y *big.Float) bool {
	return x.Cmp(y) == 0
}