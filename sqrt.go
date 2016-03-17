// Package floats provides the implementation of a few additional operations for the
// standard library big.Float type.
package floats

import (
	"math"
	"math/big"
)

// Sqrt returns a big.Float representation of the square root of z. Precision is
// the same as the one of the argument. The function panics if z is negative, returns ±0
// when z = ±0, and +Inf when z = +Inf.
func Sqrt(z *big.Float) *big.Float {

	// panic on negative z
	if z.Sign() == -1 {
		panic("Sqrt: argument is negative")
	}

	// Sqrt(±0) = ±0
	if z.Sign() == 0 {
		return big.NewFloat(float64(z.Sign()))
	}

	// Sqrt(+Inf) = +Inf
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1))
	}

	prec := z.Prec() + 64 // guard digits

	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)
	one := new(big.Float).SetPrec(prec).SetInt64(1)
	three := new(big.Float).SetPrec(prec).SetInt64(3)

	// Compute sqrt(z) via 1/sqrt(z) to avoid divisions.
	// Applying Newton to (1/x²) - z = 0 gives
	//     x_{n+1} = 0.5x_{n}(3 - zx²)
	// which uses only 3 multiplications, and converge
	// quadratically.

	// x will hold the value of 1/sqrt(z)
	x := new(big.Float).SetPrec(prec)

	// get initial estimate using IEEE-754 math
	zf, _ := z.Float64()
	if zfs := math.Sqrt(zf); zfs != 0 {
		x.SetFloat64(1 / zfs)
	} else {
		x.SetMantExp(one, z.MantExp(nil)/2)
		x.Quo(one, x)
	}

	// we need at least log_2(prec) iterations
	steps := int(math.Log2(float64(prec)))

	t := new(big.Float)

	for i := 0; i < steps; i++ {
		t.Mul(x, x)     // t = x²
		t.Mul(t, z)     // t = zx²
		t.Sub(three, t) // t = 3 - zx²
		t.Mul(t, half)
		x.Mul(x, t)
	}

	// sqrt(z) = z * (1/sqrt(z))
	x.Mul(x, z)

	return x.SetPrec(z.Prec())
}
