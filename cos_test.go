package bigfloat_test

import (
	"math/big"
	"testing"

	"github.com/ALTree/bigfloat"
)

func TestCos(t *testing.T) {
	a := bigfloat.Cos(big.NewFloat(0).SetPrec(64))
	if value := a.Text('f', 64); value != "1.0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatalf("cos(0) should be 1 but is %s", value)
	}
	pi := bigfloat.PI(64)
	pi2 := big.NewFloat(2).SetPrec(64)
	pi2.Quo(pi, pi2)
	a = bigfloat.Cos(pi2)
	if value := a.Text('f', 64); value != "0.0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatalf("cos(pi/2) should be 0 but is %s", value)
	}
	a = bigfloat.Cos(pi)
	if value := a.Text('f', 64); value != "-1.0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatalf("cos(pi) should be -1 but is %s", value)
	}
}
