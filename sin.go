package bigfloat

import (
	"math/big"
)

// Sin compute the sin of the value
func Sin(z *big.Float) *big.Float {
	prec := z.Prec()
	one := big.NewFloat(1).SetPrec(prec)
	zz := big.NewFloat(0).SetPrec(prec)
	zz.Set(z)
	sign := true
	if zz.Cmp(big.NewFloat(0).SetPrec(0)) < 0 {
		zz.Neg(zz)
		sign = false
	}
	a, d, i := big.NewFloat(0).SetPrec(prec), big.NewFloat(6).SetPrec(prec),
		big.NewFloat(3).SetPrec(prec)
	for {
		b := Pow(zz, i)
		b = b.Quo(b, d)
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
			d.Mul(d, i)
		}
	}
	return a.Add(z, a)
}
