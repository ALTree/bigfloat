package bigfloat

import (
	"math/big"
)

// Tan is the tan function
func Tan(z *big.Float) *big.Float {
	sin, cos := Sin(z), Cos(z)
	return sin.Quo(sin, cos)
}
