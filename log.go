package floatutils

import (
	"math"
	"math/big"
)

// Log returns a big.Float representation of the natural logarithm of z. Precision is
// the same as the one of the argument. The function panics if z is negative, returns -Inf
// when z = 0, and +Inf when z = +Inf
func Log(z *big.Float) *big.Float {

	// panic on negative z
	if z.Sign() == -1 {
		panic("Log: argument is negative")
	}

	// Log(0) = -Inf
	if z.Sign() == 0 {
		return big.NewFloat(math.Inf(-1))
	}

	one := new(big.Float).SetInt64(1)

	// Log(1) = 0
	if z.Cmp(one) == 0 {
		return new(big.Float).SetPrec(z.Prec()).SetInt64(0)
	}

	// Log(+Inf) = +Inf
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1))
	}

	// 64 bits of guard digits like in sqrt.go
	prec := z.Prec() + 64

	x := new(big.Float).SetPrec(prec)

	// if 0 < x < 1 we compute log(x) as -log(1/x)
	var neg bool
	if z.Cmp(one) < 0 {
		x.Quo(one, z)
		neg = true
	} else {
		x.Set(z)
	}

	// We scale up x until x >= 2**(prec/2), and then we'll be
	// allowed to use the AGM formula for Log(x).
	// Double x until the condition is met, and keep track of
	// the number of doubling we did (needed to scale back later).
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	lim := Pow(two, int(prec/2))

	k := 0
	for x.Cmp(lim) < 0 {
		x.Mul(x, x)
		k++
	}

	res := logBig(x)

	// change sign if the z was < 1
	if neg {
		res.Neg(res)
	}

	// scale the result back dividing by 2**k
	res.Quo(res, Pow(two, k))

	return res.SetPrec(z.Prec())
}

// log_big computes the natural log of z using the fact that
//     log(z) = π / (2 * AGM(1, 4/z))
// if
//     z >= 2**(prec/2),
// where prec is the desired precision (in bits)
func logBig(z *big.Float) *big.Float {

	prec := z.Prec()

	one := new(big.Float).SetPrec(prec).SetInt64(1)
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	four := new(big.Float).SetPrec(prec).SetInt64(4)

	pi := pi(prec)

	t := new(big.Float)
	agm := agm(one, t.Quo(four, z))

	return t.Quo(pi, t.Mul(two, agm))
}
