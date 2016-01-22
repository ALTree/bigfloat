package floatutils_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/floatutils"
)

const maxPrec uint = 1100

func TestSqrt(t *testing.T) {
	for _, test := range []struct {
		x    string
		want string
	}{
		// 350 decimal digits are enough to give us up to 1000 binary digits
		{"0.5", "0.70710678118654752440084436210484903928483593768847403658833986899536623923105351942519376716382078636750692311545614851246241802792536860632206074854996791570661133296375279637789997525057639103028573505477998580298513726729843100736425870932044459930477616461524215435716072541988130181399762570399484362669827316590441482031030762917619752737287514"},
		{"2.0", "1.4142135623730950488016887242096980785696718753769480731766797379907324784621070388503875343276415727350138462309122970249248360558507372126441214970999358314132226659275055927557999505011527820605714701095599716059702745345968620147285174186408891986095523292304843087143214508397626036279952514079896872533965463318088296406206152583523950547457503"},
		{"3.0", "1.7320508075688772935274463415058723669428052538103806280558069794519330169088000370811461867572485756756261414154067030299699450949989524788116555120943736485280932319023055820679748201010846749232650153123432669033228866506722546689218379712270471316603678615880190499865373798593894676503475065760507566183481296061009476021871903250831458295239598"},
		{"4.0", "2.0"},

		{"5e10", "223606.79774997896964091736687312762354406183596115257242708972454105209256378048994144144083787822749695081761507737835042532677244470738635863601215334527088667781731918791658112766453226398565805357613504175337850034233924140644420864325390972525926272288762995174024406816117759089094984923713907297288984820886415426898940991316935770197486788844"},
		{"6e100", "2.4494897427831780981972840747058913919659474806566701284326925672509603774573150265398594331046402348185946012266141891248588654598377573416257839512372785528289127475276765712476301052709117702234813106789866908536324433525456040338088089393745855678465747243613041442702702161742018383000815898078380130897007286939936308371580944008004437386875492e50"},
		{"7e1000", "2.6457513110645905905016157536392604257102591830824501803683344592010688232302836277603928864745436106150645783384974630957435298886272147844273905558801077227171507297283238922996895948650872607009780542037238280237159411003419391160015785255963059457410351523968027164073737990740415815199044034743194536713997305970050513996922375456160971190273782e500"},

		{"10e-10", "0.000031622776601683793319988935444327185337195551393252168268575048527925944386392382213442481083793002951873472841528400551485488560304538800146905195967001539033449216571792599406591501534741133394841240853169295770904715764610443692578790620378086099418283717115484063285529991185968245642033269616046913143361289497918902665295436126761787813500613882"},
		{"11e-100", "3.3166247903553998491149327366706866839270885455893535970586821461164846426090438467088433991282906509070125578495274565922754397848575474797793249330447288473028739748286556825773944446120980444771931123571441329715210988326604957100372485207381068208074875839658949945251593152984006827197105182895572484245567663700205081298422367705978420397374274e-50"},
		{"13e-1000", "3.6055512754639892931192212674704959462512965738452462127104530562271669482930104452046190820184907176735141820240635403760306782646978077051630171668927097577426905642741526332338303949623469447962732299962880032688564272130721127331690722052975017855588384448146538689210753953924825633102446828366387325286968101892633130498355399216211744108618057e-500"},
	} {
		for _, prec := range []uint{24, 53, 64, 100, 200, 300, 400, 500, 600, 700, 800, 900, 1000} {
			want := new(big.Float).SetPrec(prec)
			want.Parse(test.want, 10)

			// We can't guarantee that the result will have *prec* precision
			// if we call Sqrt with an argument with *prec* precision, because
			// the Newton's iteration will actually converge to a number that
			// is not the square root of x with *prec+64* precision, but to a
			// number that is the square root of x with *prec* precision.
			// If we want Sqrt(x) with *prec* precision and correct rounding,
			// we need to call Sqrt with an argument having precision greater
			// than *prec*.
			// TODO: document this
			x := new(big.Float).SetPrec(prec + 16) // TODO: find a sensible value
			x.Parse(test.x, 10)

			z := floatutils.Sqrt(x)
			want.SetPrec(prec)

			wantMaxPrec, _, err := big.ParseFloat(test.want, 0, maxPrec, big.ToNearestEven)
			if err != nil {
				t.Errorf("prec = %d, parse(%s): %v", maxPrec, test.want, err)
			}
			acc := big.Accuracy(want.Cmp(wantMaxPrec))

			z.SetPrec(prec)

			if z.Cmp(want) != 0 || z.Acc() != acc {
				t.Errorf("prec = %d, Sqrt(%v) = %g (%v); want %g (%v)", prec, test.x, z, z.Acc(), want, acc)
			}
		}
	}
}

func TestSqrt32(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float32() * 1e1
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Sqrt(x).Float32()
		want := math.Sqrt(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt32Small(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float32() * 1e-30
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Sqrt(x).Float32()
		want := math.Sqrt(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}
func TestSqrt32Big(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float32() * 1e30
		x := big.NewFloat(float64(r)).SetPrec(24)
		z, acc := floatutils.Sqrt(x).Float32()
		want := math.Sqrt(float64(r))
		if z != float32(want) || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt64(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float64() * 1e1
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(r)
		if z != want || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt64Small(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float64() * 1e-300
		x := big.NewFloat(r).SetPrec(53)
		z, acc := floatutils.Sqrt(x).Float64()
		want := math.Sqrt(r)
		if z != want || acc != big.Exact {
			t.Errorf("Sqrt(%g) =\n got %b (%s);\nwant %b (Exact)", x, z, acc, want)
		}
	}
}

func TestSqrt64Big(t *testing.T) {
	for i := 0; i < 1e5; i++ {
		r := rand.Float64() * 1e300
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
