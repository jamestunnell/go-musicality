package model

import (
	"math/big"
	"sort"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/section"
)

type FlatScore struct {
	Parts                        map[string]note.Notes
	StartDynamic, StartTempo     float64
	DynamicChanges, TempoChanges change.Map
}

var zero = big.NewRat(0, 1)

func NewFlatScore() *FlatScore {
	return &FlatScore{
		Parts:          map[string]note.Notes{},
		StartDynamic:   section.DefaultStartDynamic,
		StartTempo:     section.DefaultStartTempo,
		DynamicChanges: change.Map{},
		TempoChanges:   change.Map{},
	}
}

func (s *FlatScore) Duration() rat.Rat {
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

func (s *FlatScore) DynamicComputer() (*Computer, error) {
	return NewComputer(s.StartDynamic, s.DynamicChanges)
}

func (s *FlatScore) TempoComputer() (*Computer, error) {
	return NewComputer(s.StartTempo, s.TempoChanges)
}
