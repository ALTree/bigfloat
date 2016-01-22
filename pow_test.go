package floatutils_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floatutils"
)

func TestPow(t *testing.T) {
	for _, test := range []struct {
		x    string
		n    int
		want string
	}{
		// 350 decimal digits are enough to give us up to 1000 binary digits
		{"1.0", 4, "1.0"},
		{"2.0", 8, "256.0"},
		{"2.5", 8, "1525.87890625"},
		{"3.0", 8, "6561.0"},
		{"3.4", 16, "3.189059870763703892770816e8"},
		{"4.6", 32, "1.61529040680870074100680119806799048214504294859145216e21"},
		{"5.7", 64, "2.3767897344134118845792411633735801064687423783821619363360084403117800446467034136609454101929141476456488896001e48"},
	} {
		for _, prec := range []uint{24, 53, 64, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000} {
			want := new(big.Float).SetPrec(prec)
			want.Parse(test.want, 10)

			x := new(big.Float).SetPrec(prec + 16) // see sqrt_test.go
			x.Parse(test.x, 10)

			z := floatutils.Pow(x, test.n)
			want.SetPrec(prec)

			wantMaxPrec, _, err := big.ParseFloat(test.want, 0, maxPrec, big.ToNearestEven)
			if err != nil {
				t.Errorf("prec = %d, parse(%s): %v", maxPrec, test.want, err)
			}
			acc := big.Accuracy(want.Cmp(wantMaxPrec))

			z.SetPrec(prec)

			if z.Cmp(want) != 0 || z.Acc() != acc {
				t.Errorf("prec = %d, Pow(%v, %d) = %g (%v); want %g (%v)", prec, test.x, test.n, z, z.Acc(), want, acc)
			}
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
			t.Errorf("Pow(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
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
			t.Errorf("%d) Pow(%g) =\n got %b (%s);\nwant %b (Exact)", i, test.f, z, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

func benchmarkPow(prec uint, exp int, b *testing.B) {
	b.ReportAllocs()
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
