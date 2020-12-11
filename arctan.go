package bigfloat

import (
	"math/big"
)

// Arctan is the inverse tan
func Arctan(z *big.Float) *big.Float {
	prec := z.Prec()
	one := big.NewFloat(1).SetPrec(prec)
	minusone := big.NewFloat(-1).SetPrec(prec)
	zz := big.NewFloat(0).SetPrec(prec)
	zz.Set(z)
	sign := true
	if zz.Cmp(big.NewFloat(0).SetPrec(0)) < 0 {
		zz.Neg(zz)
		sign = false
	}
	if z.Cmp(one) >= 0 {
		pi := PI(prec)
		a, i := big.NewFloat(0).SetPrec(prec), big.NewFloat(1).SetPrec(prec)
		a.Quo(pi, big.NewFloat(2).SetPrec(prec))
		for {
			b := Pow(zz, i)
			b.Mul(b, i)
			b = b.Quo(one, b)
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

		return a
	} else if z.Cmp(minusone) > 0 {
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
		a = a.Add(z, a)
		return a
	}

	pi := PI(prec)
	a, i := big.NewFloat(0).SetPrec(prec), big.NewFloat(1).SetPrec(prec)
	a.Quo(pi, big.NewFloat(2).SetPrec(prec))
	a.Neg(a)
	for {
		b := Pow(zz, i)
		b.Mul(b, i)
		b = b.Quo(one, b)
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

	return a
}
