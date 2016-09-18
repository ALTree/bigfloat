### Floats 

Package bigfloat provides the implementation of a few additional operations (square root, exponentiation, natural logarithm, exponential function) for the standard library `big.Float` type.

[![GoDoc](https://godoc.org/github.com/ALTree/bigfloat?status.png)](https://godoc.org/github.com/ALTree/bigfloat)

#### Install

```
go get github.com/ALTree/bigfloat
```

Please note that `bigfloat` requires Go >= 1.5 (since the `big.Float` type is not available in previous versions). 

#### Example

```go
package main

import (
	"fmt"
	"math/big"

	"github.com/ALTree/bigfloat"
)

// In this example, we'll compute the value of the
// trascendental number 2 ** √2, also known as
// the Gelfond–Schneider constant, to 1000 bits.
func main() {
	// Work with 1000 binary digits of precision.
	const prec = 1000

	two := big.NewFloat(2).SetPrec(prec)

	// Compute √2.
	// Sqrt uses the argument's precision.
	sqrtTwo := bigfloat.Sqrt(two)

	// Compute 2 ** √2
	// Pow uses the first argument's precision.
	gsc := bigfloat.Pow(two, sqrtTwo)

	// Print gsc, truncated to 60 decimal digits.
	fmt.Printf("gsc = %.60f...\n", gsc)
}
```

outputs
```
gsc = 2.665144142690225188650297249873139848274211313714659492835980...
```

#### Documentation

See https://godoc.org/github.com/ALTree/bigfloat
