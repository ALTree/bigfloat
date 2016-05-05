package floats

import (
	"math"
	"math/big"
)

// agm returns the arithmetic-geometric mean of a and b, to
// max(a.Prec, b.Prec) bits of precision. a and b must have
// the same precision.
func agm(a, b *big.Float) *big.Float {

	if a.Prec() != b.Prec() {
		panic("agm: different precisions")
	}

	var prec uint = a.Prec() + 64

	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	a2 := new(big.Float)
	b2 := new(big.Float)
	a2.Copy(a).SetPrec(prec)
	b2.Copy(b).SetPrec(prec)

	checkA := new(big.Float).Copy(a2).SetPrec(prec - 64)
	checkB := new(big.Float).Copy(b2).SetPrec(prec - 64)
	t1 := new(big.Float)

	// iterate until a2 == b2 (with a2 and b2 reduced to the
	// desired precision, i.e. ignoring the guard digits)
	for checkA.Cmp(checkB) != 0 {
		t1.Copy(a2)
		a2.Add(a2, b2).Mul(a2, half)
		b2 = Sqrt(b2.Mul(t1, b2))

		checkA.Copy(a2).SetPrec(prec - 64)
		checkB.Copy(b2).SetPrec(prec - 64)
	}

	return checkA
}

var piCache *big.Float
var piCachePrec uint
var enablePiCache bool = true

func init() {
	if !enablePiCache {
		return
	}

	piCache, _, _ = new(big.Float).SetPrec(1024).Parse("3."+
		"14159265358979323846264338327950288419716939937510"+
		"58209749445923078164062862089986280348253421170679"+
		"82148086513282306647093844609550582231725359408128"+
		"48111745028410270193852110555964462294895493038196"+
		"44288109756659334461284756482337867831652712019091"+
		"45648566923460348610454326648213393607260249141273"+
		"72458700660631558817488152092096282925409171536444", 10)

	piCachePrec = 1024
}

// pi returns pi to prec bits of precision
func pi(prec uint) *big.Float {

	if prec <= piCachePrec && enablePiCache {
		return new(big.Float).Copy(piCache).SetPrec(prec)
	}

	precExt := prec + 64

	half := big.NewFloat(0.5)
	one := big.NewFloat(1)
	two := big.NewFloat(2)
	four := big.NewFloat(4)

	temp := new(big.Float)
	a := new(big.Float).SetPrec(precExt).SetInt64(1)         // a_0 = 1
	b := new(big.Float).Quo(one, Sqrt(two.SetPrec(precExt))) // b_0 = 1/sqrt(2)
	t := new(big.Float).SetPrec(precExt).SetFloat64(0.25)    // t_0 = 1/4
	p := new(big.Float).SetPrec(precExt).SetInt64(1)         // p_0 = 1

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
	res := new(big.Float).Quo(temp, t.Mul(four, t)) // pi = (a_{n+1} + b_{n+1})² / (4t_{n+1})
	res.SetPrec(prec)

	if enablePiCache {
		piCache.Copy(res)
		piCachePrec = prec
	}

	return res
}

// returns an approximate (to precision dPrec) solution to
//    f(t) = 0
// using the Newton Method.
// fOverDf needs to be a fuction returning f(t)/f'(t).
// t must not be changed by fOverDf.
// guess is the initial guess (and it's not preserved).
func newton(fOverDf func(z *big.Float) *big.Float, guess *big.Float, dPrec uint) *big.Float {

	prec, guard := guess.Prec(), uint(64)
	guess.SetPrec(prec + guard)

	for prec < 2*dPrec {
		guess.Sub(guess, fOverDf(guess))
		prec *= 2
		guess.SetPrec(prec + guard)
	}

	return guess.SetPrec(dPrec)
}
