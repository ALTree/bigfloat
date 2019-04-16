package bigfloat

import "math/big"

// Pow returns a big.Float representation of z**w. Precision is the same as the one
// of the first argument. The function panics when z is negative.
func Pow(z *big.Float, w *big.Float) *big.Float {

	// Negative bases with non-integer exponents are undefined
	if z.Sign() < 0 {
		if w.IsInt() {
			// Flip sign of base and proceed
			z = new(big.Float).Copy(z).Neg(z)
			// If exponent is odd, then return Neg of Pow
			if !new(big.Float).Quo(w, big.NewFloat(2)).IsInt() {
				return new(big.Float).SetPrec(z.Prec()).Neg(Pow(z, w))
			}
		} else {
			panic(ErrNaN{"Pow: negative base with non-integer exponent"})
		}
	}

	// Pow(z, 0) = 1.0
	if w.Sign() == 0 {
		return big.NewFloat(1).SetPrec(z.Prec())
	}

	// Pow(z, 1) = z
	if w.Cmp(big.NewFloat(1)) == 0 {
		return new(big.Float).Copy(z)
	}

	// Pow(+Inf, n) = +Inf for n > 0
	// Pow(+Inf, n) = +0 for n < 0
	if z.IsInf() {
		if w.Sign() > 0 {
			return new(big.Float).Copy(z)
		}
		return big.NewFloat(0).SetPrec(z.Prec())
	}

	// Pow(z, -w) = 1 / Pow(z, w)
	if w.Sign() < 0 {
		x := new(big.Float)
		zExt := new(big.Float).Copy(z).SetPrec(z.Prec() + 64)
		wNeg := new(big.Float).Neg(w)
		return x.Quo(big.NewFloat(1), Pow(zExt, wNeg)).SetPrec(z.Prec())
	}

	// // w integer fast path (disabled because introduces rounding
	// // errors)
	// if false && w.IsInt() {
	// 	wi, _ := w.Int64()
	// 	return powInt(z, int(wi))
	// }

	// compute w**z as exp(z log(w))
	x := new(big.Float).SetPrec(z.Prec() + 64)
	logZ := Log(new(big.Float).Copy(z).SetPrec(z.Prec() + 64))
	x.Mul(w, logZ)
	x = Exp(x)
	return x.SetPrec(z.Prec())

}

// // fast path for z**w when w is an integer
// func powInt(z *big.Float, w int) *big.Float {

// 	// get mantissa and exponent of z
// 	mant := new(big.Float)
// 	exp := z.MantExp(mant)

// 	// result's exponent
// 	exp = exp * w

// 	// result's mantissa
// 	x := big.NewFloat(1).SetPrec(z.Prec())

// 	// Classic right-to-left binary exponentiation
// 	for w > 0 {
// 		if w%2 == 1 {
// 			x.Mul(x, mant)
// 		}
// 		w >>= 1
// 		mant.Mul(mant, mant)
// 	}

// 	return new(big.Float).SetMantExp(x, exp)
// }
