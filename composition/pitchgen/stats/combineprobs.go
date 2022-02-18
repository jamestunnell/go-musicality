package stats

import "errors"

//
func CombineAndNormalizeProbs(probArrays [][]float64) ([]float64, error) {
	// Sanity check #1
	if len(probArrays) == 0 {
		err := errors.New("no probability arrays given")
		return []float64{}, err
	}

	n := len(probArrays[0])

	// Sanity check #2
	for i := 1; i < len(probArrays); i++ {
		if len(probArrays[i]) != n {
			err := errors.New("probability arrays do not all have the same length")
			return []float64{}, err
		}
	}

	// Compute the un-normalized probability of each semitone by multiplying all
	// the component probabilities.
	total := 0.0
	probs := make([]float64, n)

	for i := 0; i < n; i++ {
		prob := 1.0

		for j := 0; j < len(probArrays); j++ {
			prob *= probArrays[j][i]
		}

		probs[i] = prob
		total += prob
	}

	// Normalize the probabilities so they sum to 1
	for i := 0; i < n; i++ {
		probs[i] /= total
	}

	return probs, nil
}
