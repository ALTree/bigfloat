package floats_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floats"
)

// See note in sqrt_test.go about which numbers
// can we safely test this way.

func TestLog(t *testing.T) {
	for _, test := range []struct {
		x    string
		want string
	}{
		// 350 decimal digits are enough to give us up to 1000 binary digits
		{"0.5", "-0.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641868754200148102057068573368552023575813055703267075163507596193072757082837143519030703862389167347112335011536449795523912047517268157493206515552473413952588295045300709532636664265410423915781495204374043038550080194417064167151864471283996817178454696"},
		{"0.25", "-1.3862943611198906188344642429163531361510002687205105082413600189867872439393894312117266539928373750840029620411413714673710404715162611140653415032701519238614551416567428703806140772477833469422467002307289959104782409503453631498641303110494682790517659009060141906527332853082084783156299040874808607710016038883412833430372894256799363435690939"},
		{"0.0125", "-4.3820266346738816122696878190588939118276018917095387383953679294477534755864366270535871860788543609679722271039983058344660861723571984277642996240040095752750899208106689864147210106979082189417635556550588715983462075888842670124944153533207460860520530946333864410280342429017041970928492563533263928706772062013203262792640026952942261381891629"},

		{"1", "0.0"},
		{"2", "0.69314718055994530941723212145817656807550013436025525412068000949339362196969471560586332699641868754200148102057068573368552023575813055703267075163507596193072757082837143519030703862389167347112335011536449795523912047517268157493206515552473413952588295045300709532636664265410423915781495204374043038550080194417064167151864471283996817178454696"},
		{"10", "2.3025850929940456840179914546843642076011014886287729760333279009675726096773524802359972050895982983419677840422862486334095254650828067566662873690987816894829072083255546808437998948262331985283935053089653777326288461633662222876982198867465436674744042432743651550489343149393914796194044002221051017141748003688084012647080685567743216228355220"},
		{"4096", "8.3177661667193437130067854574981188169060016123230630494481601139207234636363365872703599239570242505040177722468482288042262428290975666843920490196209115431687308499404572222836844634867000816534802013843739754628694457020721788991847818662968096743105954054360851439163997118492508698937794245248851646260096233300477000582237365540796180614145634"},
		{"1e5", "11.512925464970228420089957273421821038005507443143864880166639504837863048386762401179986025447991491709838920211431243167047627325414033783331436845493908447414536041627773404218999474131165992641967526544826888663144230816831111438491099433732718337372021216371825775244671574696957398097022001110525508570874001844042006323540342783871608114177610"},
	} {
		for _, prec := range []uint{24, 53, 64, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000} {
			want := new(big.Float).SetPrec(prec)
			want.Parse(test.want, 10)

			x := new(big.Float).SetPrec(prec)
			x.Parse(test.x, 10)

			z := floats.Log(x)

			if z.Cmp(want) != 0 {
				t.Errorf("prec = %d, Log(%v) =\ngot %g;\n want %g", prec, test.x, z, want)
			}
		}
	}
}

func TestLog32(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float32() * 1e3
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floats.Log(x).Float32()
		want := math.Log(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Log(%f) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestLog32Small(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float32() * 1e-30
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floats.Log(x).Float32()
		want := math.Log(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Log(%f) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestLog32Big(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float32() * 1e30
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floats.Log(x).Float32()
		want := math.Log(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Log(%f) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

// Unfortunately, the Go math.Log function is not completely
// accurate, so it doesn't make sense to require 100%
// compatibility with it, since it happens that math.Log
// returns a result with the last bit off (see Issue #9546).
//
// For this reason, we just require that the result is
// within distance 1e-14 from what math.Log returns
// (1e-12 for very small values).
// TODO: figure out a good permitted error.
func TestLog64(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float64() * 1e3
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floats.Log(x).Float64()
		want := math.Log(r)
		if math.Abs(z-want) > 1e-14 || acc != big.Exact {
			t.Errorf("Log(%g) =\n got %g (%s);\nwant %g (Exact)", x, z, acc, want)
		}
	}
}

func TestLog64Small(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float64() * 1e-300
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floats.Log(x).Float64()
		want := math.Log(r)
		if math.Abs(z-want) > 1e-12 || acc != big.Exact { // 1e-12 for very small values
			t.Errorf("Log(%g) =\n got %g (%s);\nwant %g (Exact)", x, z, acc, want)
		}
	}
}

func TestLog64Big(t *testing.T) {
	for i := 0; i < 1e4; i++ {
		r := rand.Float64() * 1e300
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floats.Log(x).Float64()
		want := math.Log(r)
		if math.Abs(z-want) > 1e-14 || acc != big.Exact {
			t.Errorf("Log(%g) =\n got %g (%s);\nwant %g (Exact)", x, z, acc, want)
		}
	}
}

func TestLogSpecialValues(t *testing.T) {
	for i, f := range []float64{
		+0.0,
		-0.0,
		math.Inf(+1),
	} {
		x := big.NewFloat(f).SetPrec(53)
		z, acc := floats.Log(x).Float64()
		want := math.Log(f)
		if z != want || acc != big.Exact {
			t.Errorf("%d) Log(%f) =\n got %b (%s);\nwant %b (Exact)", i, f, z, acc, want)
		}
	}
}

// ---------- Benchmarks ----------

func benchmarkLog(num float64, prec uint, b *testing.B) {
	b.ReportAllocs()
	x := new(big.Float).SetPrec(prec).SetFloat64(num)
	for n := 0; n < b.N; n++ {
		floats.Log(x)
	}
}

func BenchmarkLog2Prec53(b *testing.B)    { benchmarkLog(2, 53, b) }
func BenchmarkLog2Prec64(b *testing.B)    { benchmarkLog(2, 64, b) }
func BenchmarkLog2Prec100(b *testing.B)   { benchmarkLog(2, 1e2, b) }
func BenchmarkLog2Prec250(b *testing.B)   { benchmarkLog(2, 250, b) }
func BenchmarkLog2Prec500(b *testing.B)   { benchmarkLog(2, 500, b) }
func BenchmarkLog2Prec1000(b *testing.B)  { benchmarkLog(2, 1e3, b) }
func BenchmarkLog2Prec10000(b *testing.B) { benchmarkLog(2, 1e4, b) }
