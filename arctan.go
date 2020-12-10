package bigfloat

import (
	"math/big"
)

// Arctan is the inverse tan
func Arctan(z *big.Float) *big.Float {
	prec := z.Prec()
	one := big.NewFloat(1).SetPrec(prec)
	two := big.NewFloat(2).SetPrec(prec)
	zz := big.NewFloat(0).SetPrec(prec)
	zz.Set(z)
	if zz.Sign() < 0 {
		zz.Sub(zz, one)
	} else {
		zz.Add(zz, one)
	}
	zz.Quo(zz, two)
	x, _ := zz.Int(nil)
	xx := big.NewFloat(0).SetPrec(prec)
	xx.SetInt(x)
	xx.Mul(xx, two)
	zz.Set(z)
	zz.Sub(zz, xx)
	negative := zz.Sign() < 0
	if negative {
		zz = zz.Abs(zz)
	}
	sign := true
	a, i := big.NewFloat(0).SetPrec(prec), big.NewFloat(3).SetPrec(prec)
	for {
		b := Pow(zz, i)
		b = b.Quo(b, i)
		cp := big.NewFloat(0).SetPrec(prec).Set(a)
		if sign {
			a.Sub(a, b)
		} else {
			a.Add(a, b)
		}
		sign = !sign
		if cp.Cmp(a) == 0 {
			break
		}
		for j := 0; j < 2; j++ {
			i.Add(i, one)
		}
	}
	a = a.Add(zz, a)
	if negative {
		a = a.Neg(a)
	}
	return a
}
