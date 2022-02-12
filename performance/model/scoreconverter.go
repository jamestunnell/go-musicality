package model

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/change"
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
		// section start dynamic can count as a change
		if lastDC := fs.DynamicChanges.Last(); lastDC != nil && lastDC.EndValue != sec.StartDynamic {
			c := change.NewImmediate(secOffset, sec.StartDynamic)

			fs.DynamicChanges = append(fs.DynamicChanges, c)
		}

		// section start meter can count as a change
		if lastBDC := fs.BeatDurChanges.Last(); lastBDC != nil && lastBDC.EndValue != sec.StartMeter.BeatDur().Float64() {
			c := change.NewImmediate(secOffset, sec.StartMeter.BeatDur().Float64())

			fs.BeatDurChanges = append(fs.BeatDurChanges, c)
		}

		fs.BeatDurChanges = append(fs.BeatDurChanges, sec.BeatDurChanges(secOffset)...)

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

		secOffset.Accum(sec.Duration())
	}

	return fs, nil
}
