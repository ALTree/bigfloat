package floats

import "math/big"

// Pow returns a big.Float representation of z**n. Precision is the same as the one
// of the first argument. The function panics when n is negative.
func Pow(z *big.Float, n int) *big.Float {

	if n < 0 {
		panic("Pow: negative power")
	}

	// Pow(z, 0) = 1.0
	if n == 0 {
		return new(big.Float).SetPrec(z.Prec()).SetFloat64(1.0)
	}

	// Pow(z, 1) = z
	// Pow(+Inf, n) = +Inf
	if n == 1 || z.IsInf() {
		x := new(big.Float)
		return x.Copy(z)
	}

	// Pow(-Inf, n) gives error
	if z.Signbit() && z.IsInf() {
		panic("Pow: -Inf base")
	}

	// get mantissa and exponent of z
	mant := new(big.Float)
	exp := z.MantExp(mant)

	// result's exponent
	exp = exp * n

	// result's mantissa
	x := new(big.Float).SetPrec(z.Prec()).SetFloat64(1.0)

	// Classic right-to-left binary exponentiation
	for n > 0 {
		if n%2 == 1 {
			x.Mul(x, mant)
		}
		n >>= 1
		mant.Mul(mant, mant)
	}

	return new(big.Float).SetPrec(z.Prec()).SetMantExp(x, exp)
}
