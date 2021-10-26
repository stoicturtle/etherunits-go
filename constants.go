package etherunits

import (
	"math/big"

	"github.com/stoicturtle/ethereumunits-go/internal/utils"
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

var (
	wei    = big.NewInt(1)
	kwei   = utils.MulWei(3)
	mwei   = utils.MulWei(6)
	gwei   = utils.MulWei(9)
	szabo  = utils.MulWei(12)
	finney = utils.MulWei(15)
	ether  = utils.MulWei(18)
	kether = utils.MulEth(3)
	mether = utils.MulEth(6)
	gether = utils.MulEth(9)
	tether = utils.MulEth(12)
)

var unitsSlice = [_max]Unit{Wei, KWei, MWei, GWei, Szabo, Finney, Ether, KEther, MEther, GEther, TEther}

var unitValueMap = map[Unit]*big.Int{
	Wei:    wei,
	KWei:   kwei,
	MWei:   mwei,
	GWei:   gwei,
	Szabo:  szabo,
	Finney: finney,
	Ether:  ether,
	KEther: kether,
	MEther: mether,
	GEther: gether,
	TEther: tether,
}

var unitEthExpMap = map[Unit]int{
	Wei:    -18,
	KWei:   -15,
	MWei:   -12,
	GWei:   -9,
	Szabo:  -6,
	Finney: -3,
	Ether:  1,
	KEther: 3,
	MEther: 6,
	GEther: 9,
	TEther: 12,
}
