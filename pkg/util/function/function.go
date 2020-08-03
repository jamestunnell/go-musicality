package function

import (
	"math"

	"github.com/jamestunnell/go-musicality/pkg/util"
)

var (
	DomainAllFloat64 = util.NewRange(-math.MaxFloat64, math.MaxFloat64)
)

type Function interface {
	At(x float64) float64
	Domain() util.Range
}
