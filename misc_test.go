package floatutils_test

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/ALTree/floatutils"
)

var want string = "3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798"

func TestPi(t *testing.T) {
	for _, test := range []struct {
		prec           uint
		decimal_digits int
	}{
		{50, 15},
		{100, 30},
		{200, 60},
		{300, 90},
	} {
		pi := floatutils.Pi(test.prec)

		if pi.Text('f', test.decimal_digits) != want[:test.decimal_digits+2] {
			t.Errorf("pi(%d)\ngot  %s\nwant %s",
				test.prec,
				pi.Text('f', test.decimal_digits),
				want[:test.decimal_digits+2])
		}
	}

}

// returns false if |a - b| > 2**(-lim)
func compareFloats(a, b *big.Float, lim uint) bool {
	limit := new(big.Float).SetPrec(lim)
	limit.Parse("2e-"+strconv.Itoa(int(lim)), 10)

	sub := new(big.Float).SetPrec(lim)
	sub.Sub(a, b)

	if sub.Abs(sub).Cmp(limit) > 0 {
		return false
	}

	return true
}
