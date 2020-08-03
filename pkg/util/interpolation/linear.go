package interpolation

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/pkg/util"
)

var (
	percentRange = util.NewRange(0.0, 1.0)
)

// Linear interpolates interpolates between the given values, with xPercent in range [0,1].
func Linear(y0, y1, xPercent float64) (float64, error) {
	if !percentRange.IncludesValue(xPercent) {
		return 0.0, fmt.Errorf("value %f is not in range [0,1]", xPercent)
	}

	y := y0 + xPercent*(y1-y0)
	return y, nil
}
