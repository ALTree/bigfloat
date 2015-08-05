package floatutils_test

import (
	"math"
	"math/big"
	"math/rand"
	"strconv"
	"testing"

	"github.com/ALTree/floatutils"
)

func TestSqrt(t *testing.T) {
	for _, test := range []struct {
		x    string
		prec uint
		want string
	}{
		// 80 decimal digits are enough to give us 250 binary digits when parsed by the Parse function
		{"0.5", 250, "0.7071067811865475244008443621048490392848359376884740365883398689953662392310535"},
		{"2.0", 250, "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070"},
		{"3.0", 250, "1.7320508075688772935274463415058723669428052538103806280558069794519330169088000"},
		{"4.0", 250, "2.0000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"1000.0", 250, "31.6227766016837933199889354443271853371955513932521682685750485279259443863923822"},

		{"5e64", 250, "2.2360679774997896964091736687312762354406183596115257242708972454105209256378048e32"},
		{"7e128", 250, "2.6457513110645905905016157536392604257102591830824501803683344592010688232302836e64"},

		{"5e-256", 250, "2.2360679774997896964091736687312762354406183596115257242708972454105209256378048e-128"},
		{"7e-512", 250, "2.6457513110645905905016157536392604257102591830824501803683344592010688232302836e-256"},
	} {
		want := new(big.Float).SetPrec(test.prec)
		want.Parse(test.want, 10)

		x := new(big.Float).SetPrec(test.prec)
		x.Parse(test.x, 10)

		z := floatutils.Sqrt(x)

		// test if precision is correctly set
		if z.Prec() != test.prec {
			t.Errorf("Sqrt(%v): got %d prec, want %d prec", x, z.Prec(), test.prec)
		}

		// test returned value
		if !compareFloats(want, z, test.prec, t) {
			t.Errorf("Sqrt(%v): error is too big.\nwant = %.100e\ngot  = %.100e\n", x, z, want)
		}
	}
}

func TestSqrt32Small(t *testing.T) {
	for i := 0; i < 3e5; i++ {
		r := rand.Float32() * 1e1
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Sqrt(x).Float32()
		want := math.Sqrt(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt32Big(t *testing.T) {
	for i := 0; i < 3e5; i++ {
		r := rand.Float32()*1e6 + 1e3
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Sqrt(x).Float32()
		want := math.Sqrt(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt64Small(t *testing.T) {
	for i := 0; i < 3e5; i++ {
		r := rand.Float64() * 1e1
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(r)
		if z != want || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt64Big(t *testing.T) {
	for i := 0; i < 3e5; i++ {
		r := rand.Float64()*1e6 + 1e3
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(r)
		if z != want || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrtSpecialValues(t *testing.T) {
	for i, f := range []float64{
		+0.0,
		-0.0,
		math.Inf(+1),
	} {
		x := big.NewFloat(f).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(f)
		if z != want || acc != big.Exact {
			t.Errorf("%d) Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", i, f, z, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

func benchmarkSqrt(num float64, prec uint, b *testing.B) {
	b.ReportAllocs()
	x := new(big.Float).SetPrec(prec).SetFloat64(num)
	var f *big.Float
	for n := 0; n < b.N; n++ {
		f = floatutils.Sqrt(x)
	}

	result = f
}

func BenchmarkSqrt2Prec10(b *testing.B)     { benchmarkSqrt(2, 1e1, b) }
func BenchmarkSqrt2Prec100(b *testing.B)    { benchmarkSqrt(2, 1e2, b) }
func BenchmarkSqrt2Prec1000(b *testing.B)   { benchmarkSqrt(2, 1e3, b) }
func BenchmarkSqrt2Prec10000(b *testing.B)  { benchmarkSqrt(2, 1e4, b) }
func BenchmarkSqrt2Prec100000(b *testing.B) { benchmarkSqrt(2, 1e5, b) }

// ---------- Benchmarks ----------

// global benchmark dummy variable
var result *big.Float

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
