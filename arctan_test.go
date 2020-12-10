package bigfloat_test

import (
	"math/big"
	"testing"

	"github.com/ALTree/bigfloat"
)

func TestArctan(t *testing.T) {
	a := big.NewFloat(1.5).SetPrec(64)
	b := big.NewFloat(-.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	b = bigfloat.Arctan(b)
	if a.Cmp(b) != 0 {
		t.Fatalf("not equal %s != %s", a.String(), b.String())
	}

	a = big.NewFloat(2.5).SetPrec(64)
	b = big.NewFloat(.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	b = bigfloat.Arctan(b)
	if a.Cmp(b) != 0 {
		t.Fatalf("not equal %s != %s", a.String(), b.String())
	}
}
