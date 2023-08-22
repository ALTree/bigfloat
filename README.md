Package bigfloat provides arbitrary-precision natural logarithm and
exponentiation for the standard library's `big.Float` type.

[![GoDoc](https://godoc.org/github.com/ALTree/bigfloat?status.png)](https://godoc.org/github.com/ALTree/bigfloat)

The package requires Go 1.10 or newer.

#### Example

```go
package main

import (
	"fmt"
	"math/big"

	"github.com/ALTree/bigfloat"
)

// We'll compute the value of the transcendental number 2^√2, also
// known as the Gelfond–Schneider constant, to 1000 bits.
func main() {
	const prec = 1000 // in binary digits
	two := big.NewFloat(2).SetPrec(prec)
	sqrtTwo := new(big.Float).SetPrec(prec).Sqrt(two)

	// Pow uses the first argument's precision.
	gsc := bigfloat.Pow(two, sqrtTwo) // 2^√2
	fmt.Printf("gsc = %.60f\n", gsc)
}
```

outputs:
```
gsc = 2.665144142690225188650297249873139848274211313714659492835980
```

#### Documentation

See https://godoc.org/github.com/ALTree/bigfloat
