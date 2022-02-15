package flatscore

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/change"
	"github.com/jamestunnell/go-musicality/notation/note"
	"github.com/jamestunnell/go-musicality/notation/rat"
	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/notation/section"
	"github.com/jamestunnell/go-musicality/performance/computer"
)

type Converter struct {
}

func (sc *Converter) Process(s *score.Score) (*FlatScore, error) {
	if result := s.Validate(); result != nil {
		return nil, fmt.Errorf("score is invalid: %w", result)
	}

	sections := s.ProgramSections()
	startDynamic := section.DefaultStartDynamic
	startTempo := section.DefaultStartTempo
	startBeatDur := DefaultStartBeatDur

	if len(sections) > 0 {
		startDynamic = sections[0].StartDynamic
		startTempo = sections[0].StartTempo
		startBeatDur = sections[0].StartMeter.BeatDuration.Float64()
	}

	dynamicChanges := change.Changes{}
	beatDurChanges := change.Changes{}
	tempoChanges := change.Changes{}
	parts := map[string]note.Notes{}
	secOffset := rat.Zero()

	for _, sec := range sections {
		dynamicChanges = append(dynamicChanges, change.NewImmediate(secOffset.Clone(), sec.StartDynamic))
		tempoChanges = append(tempoChanges, change.NewImmediate(secOffset.Clone(), sec.StartTempo))
		beatDurChanges = append(beatDurChanges, change.NewImmediate(secOffset.Clone(), sec.StartMeter.BeatDuration.Float64()))

		dynamicChanges = append(dynamicChanges, sec.DynamicChanges(secOffset.Clone())...)
		tempoChanges = append(tempoChanges, sec.TempoChanges(secOffset.Clone())...)
		beatDurChanges = append(beatDurChanges, sec.BeatDurChanges(secOffset.Clone())...)

		for _, partName := range sec.PartNames() {
			part, partFound := parts[partName]
			if !partFound && secOffset.Positive() {
				// insert a rest for previous sections, before adding notes in current section
				part = note.Notes{note.New(secOffset.Clone())}
			}

			part = append(part, sec.PartNotes(partName)...)

			parts[partName] = part
		}

		secOffset.Accum(sec.Duration())
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
