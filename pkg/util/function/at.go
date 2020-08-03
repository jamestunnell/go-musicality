package function

import "fmt"

func At(f Function, x float64) (float64, error) {
	d := f.Domain()

	if !d.IncludesValue(x) {
		return 0.0, fmt.Errorf("domain %v does not include value %f", d, x)
	}

	return f.At(x), nil
}
