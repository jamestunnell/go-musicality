package stats

import (
	"math"

	"github.com/jamestunnell/go-musicality/pkg/util"
	"gonum.org/v1/gonum/stat/distuv"
)

// GetIntProbs get the probabilities of each integer value in the given range
func GetIntProbs(dist distuv.Normal, r util.Range) []float64 {
	min := int(math.Round(r.Start))
	max := int(math.Round(r.End))

	probs := make([]float64, max-min)
	for i := 0; i < len(probs); i++ {
		x := min + i
		probs[i] = dist.Prob(float64(x))
	}

	return probs
}
