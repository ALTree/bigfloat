package floatutils_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floatutils"
)

func TestFloatPow(t *testing.T) {
	for _, test := range []struct {
		x    string
		n    int
		prec uint
		want string
	}{
		// 80 decimal digits are enough to give us 250 binary digits when parsed from the Parse function
		{"1.0", 8, 250, "1.00000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"2.0", 8, 250, "256.00000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"3.0", 8, 250, "6561.00000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"2.5", 16, 250, "2.32830643653869628906250000000000000000000000000000000000000000000000000000000000e6"},

		{"7e-5", 6, 250, "1.17649000000000000000000000000000000000000000000000000000000000000000000000000000e-25"},
		{"3.1415926535", 8, 250, "9488.531013900958727114670854001336330645937640279072830667961177817021989796878906250"},

		{"-2.0", 8, 250, "256.00000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"-2.0", 9, 250, "-512.00000000000000000000000000000000000000000000000000000000000000000000000000000000"},
		{"-2.5", 16, 250, "2.32830643653869628906250000000000000000000000000000000000000000000000000000000000e6"},
	} {
		want := new(big.Float).SetPrec(test.prec)
		want.Parse(test.want, 10)

		x := new(big.Float).SetPrec(test.prec)
		x.Parse(test.x, 10)

		z := floatutils.Pow(x, test.n)

		if z.Prec() != test.prec {
			t.Errorf("pow(%v, %d): got %d prec, want %d prec", x, test.n, z.Prec(), test.prec)
		}

		// test returned value
		if !compareFloats(want, z, test.prec, t) {
			t.Errorf("pow(%e, %d): error is too big.\nwant = %.100e\ngot  = %.100e\n", x, test.n, z, want)
		}

	}
}

func testPow64(n int, t *testing.T) {
	for i := 0; i < 2e5; i++ {
		r := rand.Float64() * 1e5
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Pow(x, n).Float64()
		want := math.Pow(r, float64(n))
		if z != want {
			t.Errorf("sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestPow64Exp2(t *testing.T)   { testPow64(2, t) }
func TestPow64Exp16(t *testing.T)  { testPow64(16, t) }
func TestPow64Exp27(t *testing.T)  { testPow64(27, t) }
func TestPow64Exp63(t *testing.T)  { testPow64(63, t) }
func TestPow64Exp100(t *testing.T) { testPow64(100, t) }

func TestPowSpecialValues(t *testing.T) {
	for i, test := range []struct {
		f float64
		n int
	}{
		{0.0, 2},
		{1.0, 2},
		{math.Inf(+1), 2},
		{math.Inf(+1), 3},

		{-0.0, 2},
		{-1.0, 2},

		{-0.0, 3},
		{-1.0, 3},
	} {
		x := big.NewFloat(test.f).SetPrec(53)
		z, acc := floatutils.Pow(x, test.n).Float64()
		want := math.Pow(test.f, float64(test.n))
		if z != want {
			t.Errorf("%d) sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", i, test.f, z, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

func benchmarkPow(prec uint, exp int, b *testing.B) {
	x := new(big.Float).SetPrec(prec).SetFloat64(2.5)
	var f *big.Float
	for n := 0; n < b.N; n++ {
		f = floatutils.Pow(x, exp)
	}

	result = f
}

func BenchmarkPow2Prec10(b *testing.B)     { benchmarkPow(1e1, 2, b) }
func BenchmarkPow2Prec100(b *testing.B)    { benchmarkPow(1e2, 2, b) }
func BenchmarkPow2Prec1000(b *testing.B)   { benchmarkPow(1e3, 2, b) }
func BenchmarkPow2Prec10000(b *testing.B)  { benchmarkPow(1e4, 2, b) }
func BenchmarkPow2Prec100000(b *testing.B) { benchmarkPow(1e5, 2, b) }

func BenchmarkPow31Prec10(b *testing.B)     { benchmarkPow(1e1, 31, b) }
func BenchmarkPow31Prec100(b *testing.B)    { benchmarkPow(1e2, 31, b) }
func BenchmarkPow31Prec1000(b *testing.B)   { benchmarkPow(1e3, 31, b) }
func BenchmarkPow31Prec10000(b *testing.B)  { benchmarkPow(1e4, 31, b) }
func BenchmarkPow31Prec100000(b *testing.B) { benchmarkPow(1e5, 31, b) }
