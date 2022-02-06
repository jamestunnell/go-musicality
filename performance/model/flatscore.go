package model

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
)

type FlatScore struct {
	Parts                        map[string]note.Notes
	StartDynamic, StartTempo     float64
	DynamicChanges, TempoChanges change.Map
}

var zero = big.NewRat(0, 1)

func NewFlatScore(s *score.Score) (*FlatScore, error) {
	if result := s.Validate(); result != nil {
		return nil, fmt.Errorf("score is invalid: %w", result)
	}

	var startDynamic float64
	var startTempo float64

	sections := s.ProgramSections()

	if len(sections) == 0 {
		startDynamic = section.DefaultStartDynamic
		startTempo = section.DefaultStartTempo
	} else {
		startDynamic = sections[0].StartDynamic
		startTempo = sections[0].StartTempo
	}

	fs := &FlatScore{
		Parts:          map[string]note.Notes{},
		StartDynamic:   startDynamic,
		StartTempo:     startTempo,
		DynamicChanges: change.Map{},
		TempoChanges:   change.Map{},
	}

	secOffset := big.NewRat(0, 1)

	for _, sec := range sections {
		secDur := sec.Duration()

		for _, partName := range sec.PartNames() {
			part, partFound := fs.Parts[partName]
			if !partFound && secOffset.Cmp(zero) == 1 {
				part = note.Notes{note.New(secOffset)}
			}

			part = append(part, sec.PartNotes(partName)...)

			fs.Parts[partName] = part
		}

		for offset, change := range sec.DynamicChanges(secOffset) {
			fs.DynamicChanges[offset] = change
		}

		for offset, change := range sec.TempoChanges(secOffset) {
			fs.TempoChanges[offset] = change
		}

		secOffset = new(big.Rat).Add(secOffset, secDur)
	}

	return fs, nil
}

func (s *FlatScore) DynamicComputer() (*Computer, error) {
	return NewComputer(s.StartDynamic, s.DynamicChanges)
}

func (s *FlatScore) TempoComputer() (*Computer, error) {
	return NewComputer(s.StartTempo, s.TempoChanges)
}
