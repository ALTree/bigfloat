package floats

import (
	"math"
	"math/big"
)

// Exp returns a big.Float representation of exp(z). Precision is
// the same as the one of the argument. The function returns +Inf
// when z = +Inf, and 0 when z = -Inf.
func Exp(z *big.Float) *big.Float {

	// exp(0) == 1
	if z.Sign() == 0 {
		return big.NewFloat(1).SetPrec(z.Prec())
	}

	// Exp(+Inf) = +Inf
	if z.IsInf() && z.Sign() > 0 {
		return big.NewFloat(math.Inf(+1)).SetPrec(z.Prec())
	}

	// Exp(-Inf) = 0
	if z.IsInf() && z.Sign() < 0 {
		return big.NewFloat(0).SetPrec(z.Prec())
	}

	x := new(big.Float)

	// try to get initial estimate using IEEE-754 math
	zf, _ := z.Float64()
	if zfs := math.Exp(zf); zfs == math.Inf(+1) || zfs == 0 {
		// too big for IEEE-754 math, perform
		// argument reduction using e^{2x} = (e^x)²
		halfZ := new(big.Float).Set(z).SetPrec(z.Prec() + 64)
		halfZ.Quo(z, big.NewFloat(2))
		halfExp := Exp(halfZ)
		return x.Mul(halfExp, halfExp).SetPrec(z.Prec())
	} else {
		// we got a nice IEEE-754 estimate
		x.SetFloat64(zfs)
	}

	var prec uint = 32

	t := new(big.Float).SetPrec(prec + 64) // guard digits

	// Solve log(x) - z = 0 for x to find exp(z),
	// using Newton's method.
	for prec < 2*z.Prec() {
		t = Log(x)  // t = log(x_n)
		t.Sub(t, z) // t = log(x_n) - z
		t.Mul(t, x) // t = x_n * (log(x_n) - z)
		x.Sub(x, t) // x_{n+1} = x_n - x_n * (log(x_n) - z)
		prec *= 2
		x.SetPrec(prec + 64)
	}

	return x.SetPrec(z.Prec())

}
