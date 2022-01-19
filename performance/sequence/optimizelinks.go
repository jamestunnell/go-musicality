package sequence

import (
	"math"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func OptimizeLinks(unlinked, untargeted *pitch.Set) PitchMap {
	n := min(unlinked.Len(), untargeted.Len())
	bestScore := math.MaxInt

	var bestComb pitch.Pitches

	var bestPerm pitch.Pitches

	unlinked.Pitches().Combination(n, func(comb pitch.Pitches) {
		untargeted.Pitches().Permutation(n, func(perm pitch.Pitches) {
			if newScore := ScoreLinking(comb, perm); newScore < bestScore {
				bestComb = comb.Clone()
				bestPerm = perm.Clone()
				bestScore = newScore
			}
		})
	})

	pm := PitchMap{}

	for i := 0; i < n; i++ {
		pm[bestComb[i]] = bestPerm[i]
	}

	return pm
}

func ScoreLinking(a, b pitch.Pitches) int {
	score := 0
	n := min(len(a), len(b))

	for i := 0; i < n; i++ {
		score += abs(a[i].Diff(b[i]))
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
