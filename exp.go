package floats

import (
	"math"
	"math/big"
)

// Exp returns a big.Float representation of exp(z). Precision is
// the same as the one of the argument.
func Exp(z *big.Float) *big.Float {

	// exp(0) == 1
	if z.Sign() == 0 {
		return big.NewFloat(1).SetPrec(z.Prec())
	}

	var prec uint = 32
	var guard uint = 64

	x := new(big.Float).SetPrec(prec + guard)

	// get initial estimate using IEEE-754 math
	zf, _ := z.Float64()
	if zfs := math.Exp(zf); zfs != 0 && zfs != math.Inf(+1) {
		x.SetFloat64(zfs)
	} else {
		panic("mat.Exp(zf) == +Inf or Zero (not implemented)")
	}

	t := new(big.Float).SetPrec(prec + guard)

	// Solve log(x) - z = 0 for x to find exp(z)
	// using Newton.
	for prec < 2*z.Prec() {
		t = Log(x)  // t = log(x_n)
		t.Sub(t, z) // t = log(x_n) - z
		t.Mul(t, x) // t = x_n * (log(x_n) - z)
		x.Sub(x, t) // x_{n+1} = x_n - x_n * (log(x_n) - z)
		prec *= 2
		x.SetPrec(prec + guard)
	}

	return x.SetPrec(z.Prec())

}
