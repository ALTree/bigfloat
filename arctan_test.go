package bigfloat_test

import (
	"math/big"
	"testing"

	"github.com/ALTree/bigfloat"
)

func TestArctan(t *testing.T) {
	a := big.NewFloat(2.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	if a.String() != "1.19028995" {
		t.Fatalf("not equal %s != %s", a.String(), "1.19028995")
	}

	a = big.NewFloat(1.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	if a.String() != "0.9827937232" {
		t.Fatalf("not equal %s != %s", a.String(), "0.9827937232")
	}

	a = big.NewFloat(.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	if a.String() != "0.463647609" {
		t.Fatalf("not equal %s != %s", a.String(), "0.463647609")
	}

	a = big.NewFloat(-.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	if a.String() != "-0.463647609" {
		t.Fatalf("not equal %s != %s", a.String(), "-0.463647609")
	}

	a = big.NewFloat(-1.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	if a.String() != "-0.9827937232" {
		t.Fatalf("not equal %s != %s", a.String(), "-0.9827937232")
	}

	a = big.NewFloat(-2.5).SetPrec(64)
	a = bigfloat.Arctan(a)
	if a.String() != "-1.19028995" {
		t.Fatalf("not equal %s != %s", a.String(), "-1.19028995")
	}
}
