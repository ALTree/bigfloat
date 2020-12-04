package bigfloat_test

import (
	"math/big"
	"testing"

	"github.com/ALTree/bigfloat"
)

func TestSin(t *testing.T) {
	a := bigfloat.Sin(big.NewFloat(0).SetPrec(64))
	if value := a.Text('f', 64); value != "0.0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatalf("sin(0) should be 0 but is %s", value)
	}
	pi := bigfloat.PI(64)
	pi2 := big.NewFloat(2).SetPrec(64)
	pi2.Quo(pi, pi2)
	a = bigfloat.Sin(pi2)
	if value := a.Text('f', 64); value != "0.9999999999999999999457898913757247782996273599565029144287109375" {
		t.Fatalf("sin(pi/2) should be 1 but is %s", value)
	}
	a = bigfloat.Sin(pi)
	if value := a.Text('f', 64); value != "-0.0000000000000000002168404344971008868014905601739883422851562500" {
		t.Fatalf("sin(pi) should be 0 but is %s", value)
	}
}
