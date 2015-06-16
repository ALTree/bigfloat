package floatutils_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floatutils"
)

func TestFloatSqrt(t *testing.T) {
	for _, test := range []struct {
		x    string
		prec uint
		want string
	}{
		// 80 decimal digits are enough to give us 250 binary digits when parsed from the Parse function
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

		if z.Prec() != test.prec {
			t.Errorf("sqrt(%v): got %d prec, want %d prec", x, z.Prec(), test.prec)
		}

		if want.Cmp(z) != 0 {
			t.Errorf("sqrt(%v):\ngot  %s\nwant %s", x, z.Text('e', 78), want.Text('e', 78))
		}
	}
}

func TestFloatSqrt32(t *testing.T) {
	for i := 2; i <= 1024; i++ {
		x := big.NewFloat(float64(i)).SetPrec(24)
		z, acc := floatutils.Sqrt(x).Float32()
		want := math.Sqrt(float64(i))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("sqrt(%g):\ngot  %b (%s)\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestFloatSqrt64(t *testing.T) {
	for i := 1; i <= 1024; i++ {
		x := big.NewFloat(float64(i)).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(float64(i))
		if z != want || acc != big.Exact {
			t.Errorf("sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)

		}
	}
}

func TestFloatSqrt64Random(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float64() * 1e4
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(r)
		if z != want || acc != big.Exact {
			t.Errorf("sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSpecialValues(t *testing.T) {
	for i, f := range []float64{
		+0.0,
		-0.0,
		math.Inf(+1),
	} {
		x := big.NewFloat(f).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(f)
		if z != want || acc != big.Exact {
			t.Errorf("%d) sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", i, f, z, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

var result *big.Float

func benchmarkSqrt(prec uint, b *testing.B) {
	x := new(big.Float).SetPrec(prec).SetFloat64(2.0)
	var f *big.Float
	for n := 0; n < b.N; n++ {
		f = floatutils.Sqrt(x)
	}

	result = f
}

func BenchmarkSqrtPrec10(b *testing.B)     { benchmarkSqrt(1e1, b) }
func BenchmarkSqrtPrec100(b *testing.B)    { benchmarkSqrt(1e2, b) }
func BenchmarkSqrtPrec1000(b *testing.B)   { benchmarkSqrt(1e3, b) }
func BenchmarkSqrtPrec10000(b *testing.B)  { benchmarkSqrt(1e4, b) }
func BenchmarkSqrtPrec100000(b *testing.B) { benchmarkSqrt(1e5, b) }
