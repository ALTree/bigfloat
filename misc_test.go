package floatutils

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

func TestAgm(t *testing.T) {
	for _, test := range []struct {
		a, b string
		prec uint
		want string
	}{
		// 80 decimal digits are enough to give us 250 binary digits when parsed from the Parse function
		{"1", "2", 250, "1.45679103104690686918643238326508197497386394322130559079417238326792645458025090"},
		{"0.1", "10", 250, "2.62166887202249236694777079630390572380399050903895551188037667381460119951168227"},
		{"1", "1e-100", 250, "0.00678105574557545088242855030146059654964512756957057536190947144205645979940838"},
		{"1", "1e-1000", 250, "0.00068177770584125429009142279977150550091076560383392436961785384435695772011031"},
		{"1", "1e-10000", 250, "0.00006821471075727537090702967297546363165634224190033526366868574357734785728139"},
		{"1", "1e-100000", 250, "6.82184069763569414455971548285692956456215254228241030877479419726412629945746452e-6"},
	} {
		want := new(big.Float).SetPrec(test.prec)
		want.Parse(test.want, 10)

		a := new(big.Float).SetPrec(test.prec)
		a.Parse(test.a, 10)

		b := new(big.Float).SetPrec(test.prec)
		b.Parse(test.b, 10)

		z := agm(a, b)

		if z.Prec() != test.prec {
			t.Errorf("agm(%v, %v): got %d prec, want %d prec", a, b, z.Prec(), test.prec)
		}

		// test returned value
		if !compareFloats(want, z, test.prec, t) {
			t.Errorf("agm(%v, %v): error is too big.\nwant = %.100e\ngot  = %.100e\n", a, b, z, want)
		}

	}
}

func TestPi(t *testing.T) {
	piStr := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899"
	want := new(big.Float).SetPrec(250)
	want.Parse(piStr, 10)

	z := pi(250)

	if z.Prec() != 250 {
		t.Errorf("pi(250): got %d prec, want %d prec", z.Prec(), 250)
	}

	// test returned value
	if !compareFloats(want, z, 250, t) {
		t.Errorf("pi(250): error is too big.\nwant = %.100e\ngot  = %.100e\n", z, want)
	}

}

// see sqrt_test.go
func compareFloats(a, b *big.Float, lim uint, t *testing.T) bool {

	limit := new(big.Float).SetPrec(lim)

	dl := int(float64(lim)*math.Log10(2)) - 1 // lim in decimal
	limit.Parse("1e-"+strconv.Itoa(dl), 10)

	sub := new(big.Float).SetPrec(lim)
	sub.Sub(a, b)

	// scale limit
	limit.SetMantExp(limit, a.MantExp(nil))

	if sub.Abs(sub).Cmp(limit) > 0 {
		t.Errorf("limit = %.100f\n", limit)
		t.Errorf("sub   = %.100f\n", sub)
		return false
	}

	return true
}
