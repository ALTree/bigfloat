package floatutils

import (
	"math"
	"math/big"
)

// agm returns the arithmetic-geometric mean of a and b, to
// max(a.Prec, b.Prec) bits of precision.
func agm(a, b *big.Float) *big.Float {

	var prec uint
	if a.Prec() > b.Prec() {
		prec = a.Prec()
	} else {
		prec = b.Prec()
	}

	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	a2 := new(big.Float)
	b2 := new(big.Float)
	a2.Copy(a)
	b2.Copy(b)

	// we need at least 2 * Log_2(prec) steps
	steps := int(math.Log2(float64(prec)))*2 + 1
	for i := 0; i < steps; i++ {
		t, t2 := new(big.Float), new(big.Float)
		a2, b2 = t.Add(a2, b2).Mul(t, half), Sqrt(t2.Mul(a2, b2))
	}

	return a2
}
