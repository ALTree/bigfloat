### Floats 

Package floats provides the implementation of a few additional operations (square root, exponentiation, natural logarithm) for the standard library `big.Float` type.

[![GoDoc](https://godoc.org/github.com/ALTree/floats?status.png)](https://godoc.org/github.com/ALTree/floats)

#### Install

```
go get github.com/ALTree/floats
```

Please note that `floats` requires Go > 1.5 (since the `big.Float` type is not available in previous versions). 

#### Example

```go
package main

import (
	"fmt"
	"math/big"

	"github.com/ALTree/floats"
)

func main() {
	const prec = 200 // 200 binary digits
	two := new(big.Float).SetPrec(prec).SetInt64(2)

	z := floats.Sqrt(two) // z.Prec is automatically set to 200

	fmt.Printf("sqrt(2) = %.50f...\n", z) // print the first 50 decimal digits
}
```
outputs
```
sqrt(2) = 1.41421356237309504880168872420969807856967187537694...
```

#### Documentation

See https://godoc.org/github.com/ALTree/floats
