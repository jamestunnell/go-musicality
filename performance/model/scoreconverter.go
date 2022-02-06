package model

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
)

type ScoreConverter struct {
}

func (sc *ScoreConverter) Process(s *score.Score) (*FlatScore, error) {
	if result := s.Validate(); result != nil {
		return nil, fmt.Errorf("score is invalid: %w", result)
	}

	sections := s.ProgramSections()
	fs := NewFlatScore()

	if len(sections) == 0 {
		fs.StartDynamic = section.DefaultStartDynamic
		fs.StartTempo = section.DefaultStartTempo
	} else {
		fs.StartDynamic = sections[0].StartDynamic
		fs.StartTempo = sections[0].StartTempo
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
