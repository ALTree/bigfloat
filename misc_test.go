package floatutils

import (
	"math/big"
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
	pi_str := "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899"
	want := new(big.Float).SetPrec(250)
	want.Parse(pi_str, 10)

	z := pi(250)

	if z.Prec() != 250 {
		t.Errorf("pi(250): got %d prec, want %d prec", z.Prec(), 250)
	}

	// test returned value
	if !compareFloats(want, z, 250, t) {
		t.Errorf("pi(250): error is too big.\nwant = %.100e\ngot  = %.100e\n", z, want)
	}

}
