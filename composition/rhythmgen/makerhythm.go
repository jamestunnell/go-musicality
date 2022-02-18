package rhythmgen

import "github.com/jamestunnell/go-musicality/common/rat"

func MakeRhythm(totalDur rat.Rat, g RhythmGenerator) rat.Rats {
	durs := rat.Rats{}

	for durs.Sum().Less(totalDur) {
		dur := g.NextDur()

		durs = append(durs, dur)
	}

	diff := durs.Sum().Sub(totalDur)
	if diff.Positive() {
		last := durs[len(durs)-1]

		durs[len(durs)-1] = last.Sub(diff)
	}

	return durs
}
