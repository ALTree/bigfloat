package bigfloat

import (
	"math/big"
)

// Tan is the tan function
// https://en.wikipedia.org/wiki/Trigonometric_functions
func Tan(z *big.Float) *big.Float {
	sin, cos := Sin(z), Cos(z)
	return sin.Quo(sin, cos)
}
