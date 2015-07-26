package floatutils

import (
	"math"
	"math/big"
	"strconv"
	"testing"
)

// returns true if |a - b| < limit, where
// limit = 0.00 ... 001 having lim precision,
// scaled by the magnitude of a
func compareFloats(a, b *big.Float, lim uint, t *testing.T) bool {

	limit := new(big.Float).SetPrec(lim)

	decimal_lim := int(float64(lim)*math.Log10(2)) - 1
	limit.Parse("1e-"+strconv.Itoa(decimal_lim), 10)

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

// ---------- Benchmarks ----------

// global benchmark dummy variable
var result *big.Float
