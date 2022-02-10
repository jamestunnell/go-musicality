package model

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
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

	secOffset := rat.Zero()

	for _, sec := range sections {
		secDur := sec.Duration()

		for _, partName := range sec.PartNames() {
			part, partFound := fs.Parts[partName]
			if !partFound && secOffset.Positive() {
				part = note.Notes{note.New(secOffset.Clone())}
			}

			part = append(part, sec.PartNotes(partName)...)

			fs.Parts[partName] = part
		}

		for _, change := range sec.DynamicChanges(secOffset.Clone()) {
			fs.DynamicChanges = append(fs.DynamicChanges, change)
		}

		for _, change := range sec.TempoChanges(secOffset.Clone()) {
			fs.TempoChanges = append(fs.TempoChanges, change)
		}

		secOffset.Accum(secDur)
	}

	return fs, nil
}
