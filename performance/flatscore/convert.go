package flatscore

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/performance/computer"
)

const (
	DefaultStartDynamic = note.ControlNormal
)

func Convert(s *score.Score) (*FlatScore, error) {
	if result := s.Validate(); result != nil {
		return nil, fmt.Errorf("score is invalid: %w", result)
	}

	sections := s.ProgramSections()
	startDynamic := DefaultStartDynamic
	startTempo := section.DefaultStartTempo
	startBeatDur := DefaultStartBeatDur

	if len(sections) > 0 {
		startTempo = sections[0].StartTempo
		startBeatDur, _ = sections[0].StartMeter.BeatDuration.Float64()
	}

	dynamicChanges := change.Changes{}
	beatDurChanges := change.Changes{}
	tempoChanges := change.Changes{}
	parts := map[string]note.Notes{}
	secOffset := rat.Zero()

	for _, sec := range sections {
		beatDur, _ := sec.StartMeter.BeatDuration.Float64()

		tempoChanges = append(tempoChanges, change.NewImmediate(secOffset, sec.StartTempo))
		beatDurChanges = append(beatDurChanges, change.NewImmediate(secOffset, beatDur))

		tempoChanges = append(tempoChanges, sec.TempoChanges(secOffset)...)
		beatDurChanges = append(beatDurChanges, sec.BeatDurChanges(secOffset)...)

		for _, partName := range sec.PartNames() {
			part, partFound := parts[partName]
			if !partFound && rat.IsPositive(secOffset) {
				// insert a rest for previous sections, before adding notes in current section
				part = note.Notes{note.New(secOffset)}
			}

			part = append(part, sec.PartNotes(partName)...)

			parts[partName] = part
		}

		secOffset = rat.Add(secOffset, sec.Duration())
	}

	dc, err := computer.New(startDynamic, dynamicChanges)
	if err != nil {
		return nil, fmt.Errorf("failed to make dynamic computer: %w", err)
	}

	tc, err := computer.New(startTempo, tempoChanges)
	if err != nil {
		return nil, fmt.Errorf("failed to make tempo computer: %w", err)
	}

	bdc, err := computer.New(startBeatDur, beatDurChanges)
	if err != nil {
		return nil, fmt.Errorf("failed to make beat duration computer: %w", err)
	}

	fs := &FlatScore{
		Parts:           parts,
		BeatDurComputer: bdc,
		TempoComputer:   tc,
		DyamicComputer:  dc,
	}

	return fs, nil
}
