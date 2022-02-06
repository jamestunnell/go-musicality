package stats

import (
	"gonum.org/v1/gonum/stat/distuv"
)

// GetIntProbs get the probabilities of each integer value in the given range
func GetIntProbs(dist distuv.Normal, min, max int) []float64 {
	probs := make([]float64, max-min)
	for i := 0; i < len(probs); i++ {
		x := min + i
		probs[i] = dist.Prob(float64(x))
	}

	return probs
}
