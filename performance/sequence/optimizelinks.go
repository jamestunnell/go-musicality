package sequence

import (
	"math"

	"github.com/jamestunnell/go-musicality/notation/pitch"
)

func OptimizeLinks(unlinked, untargeted *pitch.Set) PitchMap {
	n := min(unlinked.Len(), untargeted.Len())
	bestScore := math.MaxInt

	var bestComb1 pitch.Pitches

	var bestComb2 pitch.Pitches

	unlinked.Pitches().Combination(n, func(comb1 pitch.Pitches) {
		untargeted.Pitches().Combination(n, func(comb2 pitch.Pitches) {
			comb1.Sort()
			comb2.Sort()

			if newScore := ScoreLinking(comb1, comb2); newScore < bestScore {
				bestComb1 = comb1.Clone()
				bestComb2 = comb2.Clone()
				bestScore = newScore
			}
		})
	})

	pm := PitchMap{}

	for i := 0; i < n; i++ {
		pm[bestComb1[i]] = bestComb2[i]
	}

	return pm
}

func ScoreLinking(a, b pitch.Pitches) int {
	score := 0
	n := min(len(a), len(b))

	for i := 0; i < n; i++ {
		score += Abs(a[i].Diff(b[i]))
	}

	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
