package floatutils_test

import (
	"fmt"
	"math"
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

// returns true if |a - b| < 2**(-lim)
func compareFloats(a, b *big.Float, lim uint, t *testing.T) bool {

	limit := new(big.Float).SetPrec(lim)

	decimal_lim := int(float64(lim)*math.Log10(2)) - 1
	limit.Parse("1e-"+strconv.Itoa(decimal_lim), 10)

	sub := new(big.Float).SetPrec(lim)
	sub.Sub(a, b)

	fmt.Printf("limit = %.150f\n", limit)

	// scale limit
	limit.SetMantExp(limit, a.MantExp(nil))

	fmt.Printf("limit = %.150f\n", limit)

	if sub.Abs(sub).Cmp(limit) > 0 {
		t.Errorf("limit = %.100f\n", limit)
		t.Errorf("sub   = %.100f\n", sub)
		return false
	}

	return true
}

// ---------- Benchmarks ----------

// global benchmark dummy variable
var result *big.Float
