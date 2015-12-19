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
	a2.Copy(a).SetPrec(prec)
	b2.Copy(b).SetPrec(prec)

	t, t2 := new(big.Float), new(big.Float)
	ttest := new(big.Float)
	for ttest.Cmp(a2) != 0 { // iterate until a2 reach a fixed value
		ttest.Copy(a2)
		t.Add(a2, b2).Mul(t, half)
		b2 = Sqrt(t2.Mul(a2, b2))
		a2.Copy(t)
	}

	return a2
}

// pi returns pi to prec bits of precision
func pi(prec uint) *big.Float {

	precExt := prec + 64

	half := new(big.Float).SetPrec(precExt).SetFloat64(0.5)
	one := new(big.Float).SetPrec(precExt).SetInt64(1)
	two := new(big.Float).SetPrec(precExt).SetInt64(2)
	four := new(big.Float).SetPrec(precExt).SetInt64(4)
	temp := new(big.Float)

	a := new(big.Float).SetPrec(precExt).SetInt64(1)      // a_0 = 1
	b := new(big.Float).Quo(one, Sqrt(two))               // b_0 = 1 / sqrt(2)
	t := new(big.Float).SetPrec(precExt).SetFloat64(0.25) // t_0 = 1 / 4
	p := new(big.Float).SetPrec(precExt).SetInt64(1)      // p_0 = 1

	steps := math.Log2(float64(precExt))

	a2 := new(big.Float)
	for i := 0; i < int(steps); i++ {
		a2.Add(a, b).Mul(a2, half) // a_{n+1} = (a_{n} + b_{n}) / 2
		b = Sqrt(temp.Mul(a, b))   // b_{n+1} = sqrt(a_{n} * b_{n})

		temp.Sub(a, a2).Mul(temp, temp) // temp = (a_{n} - a_{n+1})²
		t.Sub(t, temp.Mul(p, temp))     // t_{n+1} = t_{n} - p_{n} * temp

		p.Mul(two, p) // p_{n+1} = 2 * p_{n}

		a.Copy(a2)
	}

	temp.Add(a, b)
	temp.Mul(temp, temp)
	res := new(big.Float).Quo(temp, t.Mul(four, t)) // pi = (a_{n+1} + b_{n+1})² / (4 * t_{n+1})

	return res.SetPrec(prec)
}
