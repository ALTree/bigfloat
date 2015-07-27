### Floatutils

Package floatutils provides the implementation of a few additional operations (square root, exponentiation) for the standard library `big.Float` type.

#### Install

```
go get github.com/ALTree/floatutils
```

Please note that `floatutils` requires Go 1.5 (since the `big.Float` type is not available in previous versions). 

#### Example

```go
package main

import (
	"fmt"
	"math/big"

	fu "github.com/ALTree/floatutils"
)

func main() {
	const prec = 200 // 200 binary digits
	two := new(big.Float).SetPrec(prec).SetInt64(2)

	z := fu.Sqrt(two) // z.Prec is automatically set to 200

	fmt.Printf("sqrt(2) = %.60f...\n", z) // print the first 60 decimal digits
	// sqrt(2) = 1.414213562373095048801688724209698078569671875376948073176680...
}
```

#### Documentation

See https://godoc.org/github.com/ALTree/floatutils
