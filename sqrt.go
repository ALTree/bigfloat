package floatutils

import (
	"math"
	"math/big"
)

func Sqrt(z *big.Float) *big.Float {

	// panic on negative z < 0
	if z.Sign() == -1 {
		panic("square root of negative number")
	}

	// Sqrt(±0) = ±0
	if z.Sign() == 0 {
		return big.NewFloat(float64(z.Sign()))
	}

	// Sqrt(+Inf) = +Inf
	if z.IsInf() {
		return big.NewFloat(math.Inf(+1))
	}

	// We need extra-precision to make Sqrt results 100%
	// compatible with what math.Sqrt returns on regular float64s.
	// Not sure how much extra bits we really need, but
	// forcing a whole new Word in the nat representation seems reasonable.
	// 32 seems to work too, but leave 64 for now.
	// Performance is about 40% worse when prec ~ 10,
	// otherwise there's no difference.
	prec := z.Prec() + 64 // TODO: find a better value

	one := new(big.Float).SetPrec(prec).SetInt64(1)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// initial estimate
	x := new(big.Float).SetPrec(prec).SetMantExp(one, z.MantExp(nil)/2)

	t := new(big.Float).SetPrec(prec)

	// Newton:
	//     x_{n+1} = 1/2 * ( x_n + (S / x_n) )

	// we need at least log_2(prec) iterations
	steps := math.Log2(float64(prec))

	for i := 0; i < int(steps); i++ {
		t.Quo(z, x)    // t = S / x_n
		t.Add(t, x)    // t = x_n + (S / x_n)
		x.Mul(t, half) // x = t / 2
	}

	return x.SetPrec(z.Prec())
}
