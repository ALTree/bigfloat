package bigfloat_test

import (
	"math"
	"math/big"
	"math/rand"
	"testing"

	"github.com/ALTree/bigfloat"
)

func TestTan(t *testing.T) {
	rand.Seed(1)
	for i := 0; i < 256; i++ {
		a := big.NewFloat((rand.Float64() - .5) * math.Pi).SetPrec(64)
		b := bigfloat.Tan(a)
		t.Log("b=", b.String())
		c := bigfloat.Arctan(b)
		if a.String() != c.String() {
			t.Fatalf("not equal %s != %s", a.String(), c.String())
		}
	}
}
