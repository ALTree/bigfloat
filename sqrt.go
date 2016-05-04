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

	// setup newton

	// f(t) = t² - z
	f := func(t *big.Float) *big.Float {
		x := new(big.Float)
		x.Mul(t, t)
		return x.Sub(x, z)
	}

	// 1/f'(t) = 1/(2t)
	dfInv := func(t *big.Float) *big.Float {
		x := new(big.Float).SetPrec(t.Prec())
		one, two := big.NewFloat(1), big.NewFloat(2)
		return x.Quo(one, x.Mul(two, t))
	}

	// initial guess
	zf, _ := z.Float64()
	guess := new(big.Float)
	if zfs := math.Sqrt(zf); zfs != 0 && zfs != math.Inf(+1) {
		guess.SetFloat64(zfs)
	} else {
		// how many correct digits the "halven exponent"
		// trick gives us? Guess 2...
		one := big.NewFloat(1).SetPrec(2)
		guess.SetMantExp(one, z.MantExp(nil)/2)
	}

	// call newton
	x := newton(f, dfInv, guess, z.Prec())
	return x
}

func Sqrt2(z *big.Float) *big.Float {

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

	half := new(big.Float).SetFloat64(0.5)
	three := new(big.Float).SetInt64(3)

	// Compute sqrt(z) via 1/sqrt(z) to avoid divisions.
	// Applying Newton to (1/x²) - z = 0 gives
	//     x_{n+1} = 0.5x_{n}(3 - zx²)
	// which uses only 3 multiplications, and converge
	// quadratically.

	// x will hold the value of 1/sqrt(z)
	x := new(big.Float).SetPrec(prec)

	// get initial estimate using IEEE-754 math
	zf, _ := z.Float64()
	if zfs := math.Sqrt(zf); zfs != 0 && 1/zfs != 0 {
		x.SetFloat64(1 / zfs)
	} else {
		return sqrtBig(z)
	}

	// we need at least log_2(prec) iterations
	steps := int(math.Log2(float64(prec)))

	t := new(big.Float)
	t2 := new(big.Float)

	for i := 0; i < steps; i++ {
		t.Mul(x, x)     // t = x²
		t.Mul(t, z)     // t = zx²
		t.Sub(three, t) // t = 3 - zx²
		t.Mul(t, half)  // t = 0.5(3 - zx²)
		t2.Copy(x)      // otherwise x won't be reused
		x.Mul(t2, t)    // x = 0.5x(3 - zx²)
	}

	// sqrt(z) = z * (1/sqrt(z))
	x.Mul(x, z)

	return x.SetPrec(z.Prec())
}

func sqrtBig(z *big.Float) *big.Float {

	prec := z.Prec() + 64

	one := new(big.Float).SetPrec(prec).SetInt64(1)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	x := new(big.Float).SetPrec(prec).SetMantExp(one, z.MantExp(nil)/2)

	t := new(big.Float)

	// Classic Newton iteration:
	//     x_{n+1} = 1/2 * ( x_n + (S / x_n) )

	steps := int(math.Log2(float64(prec))) + 1
	for i := 0; i < steps; i++ {
		t.Quo(z, x)    // t = S / x_n
		t.Add(t, x)    // t = x_n + (S / x_n)
		x.Mul(t, half) // x = t / 2
	}

	return x.SetPrec(z.Prec())
}
