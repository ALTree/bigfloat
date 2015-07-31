package floatutils

import (
	"math"
	"math/big"
)

func Agm(a, b *big.Float) *big.Float {

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
		// fmt.Printf("a = %.15f\nb = %.15f\n\n", a2, b2)
		t, t2 := new(big.Float), new(big.Float)
		a2, b2 = t.Add(a2, b2).Mul(t, half), Sqrt(t2.Mul(a2, b2))
	}

	return a2
}
