package flatscore

import (
	"math/big"
	"sort"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/performance/computer"
)

type FlatScore struct {
	Parts           map[string]note.Notes
	BeatDurComputer *computer.Computer
	TempoComputer   *computer.Computer
	DyamicComputer  *computer.Computer
}

const DefaultStartBeatDur = 0.25

var zero = big.NewRat(0, 1)

// Duration returns the duration of the longest part.
func (s *FlatScore) Duration() *big.Rat {
	if len(s.Parts) == 0 {
		return rat.Zero()
	}

	durs := rat.Rats{}

	for _, notes := range s.Parts {
		durs = append(durs, notes.TotalDuration())
	}

	sort.Sort(durs)

	return durs[durs.Len()-1]
}
