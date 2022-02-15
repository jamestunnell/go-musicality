package computer

import "github.com/jamestunnell/go-musicality/notation/change"

func SimplifyChanges(startVal float64, changes change.Changes) change.Changes {
	lastVal := startVal
	simplified := change.Changes{}

	for _, c := range changes {
		if c.EndValue != lastVal {
			simplified = append(simplified, c)

			lastVal = c.EndValue
		}
	}

	return simplified
}
