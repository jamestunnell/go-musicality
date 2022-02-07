package function

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

func At(f Function, x rat.Rat) (float64, error) {
	d := f.Domain()

	if !d.IncludesValue(x) {
		return 0.0, fmt.Errorf("domain %v does not include value %v", d, x)
	}

	return f.At(x), nil
}
