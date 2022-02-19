package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

func TrimRhythm(totalDur rat.Rat, rhythmDurs rat.Rats) {
	diff := rhythmDurs.Sum().Sub(totalDur)
	if diff.Positive() {
		last := rhythmDurs[len(rhythmDurs)-1]

		rhythmDurs[len(rhythmDurs)-1] = last.Sub(diff)
	}
}
