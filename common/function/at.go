package function

import (
	"fmt"
	"math/big"
)

func At(f Function, x *big.Rat) (float64, error) {
	d := f.Domain()

	if !d.IncludesValue(x) {
		return 0.0, fmt.Errorf("domain %v does not include value %v", d, x)
	}

	return f.At(x), nil
}
