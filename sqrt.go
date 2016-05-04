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

	mant := new(big.Float)
	exp := z.MantExp(mant)

	switch exp % 2 {
	case 1:
		mant.Mul(big.NewFloat(2), mant)
	case -1:
		mant.Mul(big.NewFloat(0.5), mant)
	}

	var x *big.Float
	if z.Prec() <= 128 {
		x = sqrtDirect(mant)
	} else {
		x = sqrtInverse(mant)
	}

	return x.SetMantExp(x, exp/2).SetPrec(z.Prec())

}

// compute sqrt(z) using newton to solve
// x² - z = 0 for x
func sqrtDirect(z *big.Float) *big.Float {
	// f(t) = t² - z
	f := func(t *big.Float) *big.Float {
		x := new(big.Float)
		x.Mul(t, t)
		return x.Sub(x, z)
	}

	// 1/f'(t) = 1/(2t)
	dfInv := func(t *big.Float) *big.Float {
		x := new(big.Float)
		one, two := big.NewFloat(1), big.NewFloat(2)
		return x.Quo(one, x.Mul(two, t))
	}

	// initial guess
	zf, _ := z.Float64()
	guess := big.NewFloat(math.Sqrt(zf))

	return newton(f, dfInv, guess, z.Prec())
}

// compute sqrt(z) using newton to solve
// 1/x² - z = 0 for x and then inverting.
// Avoids Quo() calls.
func sqrtInverse(z *big.Float) *big.Float {
	// f(t)/f'(t) = -0.5t(1 - zt²)
	f := func(t *big.Float) *big.Float {
		t1 := new(big.Float)
		x := new(big.Float)
		one := big.NewFloat(1)
		half := big.NewFloat(-0.5)

		t1.Mul(t, t)        // t1 = t²
		t1.Mul(t1, z)       // t1 = zt²
		t1.Sub(one, t1)     // t1 = 1 - zt²
		t1.Mul(t1, half)    // t1 = 0.5(1 - zt²)
		return x.Mul(t, t1) // x = 0.5t(1 - zt²)
	}

	// initial guess
	zf, _ := z.Float64()
	guess := big.NewFloat(1 / math.Sqrt(zf))

	// There's another operations after newton,
	// so we need to force it to return at least
	// a few guard digits. Use 32.
	x := newton2(f, guess, z.Prec()+32)
	return x.Mul(z, x)
}
