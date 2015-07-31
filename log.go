package floatutils

import (
	"math"
	"math/big"
)

func Log(z *big.Float) *big.Float {

	// panic on negative z
	if z.Sign() == -1 {
		panic("Natural logarithm of negative number")
	}

	// Log(0) = -Inf
	if z.Sign() == 0 {
		return big.NewFloat(math.Inf(-1))
	}

	// Log(+Inf) = +Inf
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1))
	}

	prec := z.Prec() + 64
	var neg bool
	if z.Cmp(new(big.Float).SetInt64(1)) < 0 {
		neg = true
	}

	x := new(big.Float).SetPrec(prec)

	if neg {
		x.Quo(new(big.Float).SetInt64(1), z)
	} else {
		x.Set(z)
	}

	// scale x until x >= 2**(prec/2)
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	lim := new(big.Float)
	lim = Pow(two, int(prec/2))

	k := 0
	for x.Cmp(lim) < 0 {
		x.Mul(x, x)
		k++
	}

	// now we can use log_big
	res := log_big(x)
	if neg {
		res.Mul(res, new(big.Float).SetInt64(-1))
	}

	return res.Quo(res, Pow(two, k)).SetPrec(z.Prec()) // scale the result back
}

// log_big computes the natural log of z using the fact that
//     log(z) = π / (2 * AGM(1, 4/z))
// if
//     z >= 2**(prec/2),
// where prec is the desired precision (in bits)
func log_big(z *big.Float) *big.Float {

	prec := z.Prec()

	one := new(big.Float).SetPrec(prec).SetInt64(1)
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	four := new(big.Float).SetPrec(prec).SetInt64(4)

	pi := pi(prec)

	t := new(big.Float)
	agm := Agm(one, t.Quo(four, z))

	return t.Quo(pi, t.Mul(two, agm))
}
