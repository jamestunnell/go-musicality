package midi

import (
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/jamestunnell/go-musicality/performance/model"
)

func CollectNoteEvents(fs *model.FlatScore, dc *model.Computer, part string) ([]*Event, error) {
	events := []*Event{}

	notes := fs.Parts[part]
	converter := model.NewNoteConverter(model.OptionReplaceSlursAndGlides())

	notes2, err := converter.Process(notes)
	if err != nil {
		return []*Event{}, fmt.Errorf("failed to convert notes: %w", err)
	}

	for i, n := range notes2 {
		if len(n.PitchDurs) > 1 {
			return []*Event{}, fmt.Errorf("note %d has multiple pitch durs", i)
		}

		pd := n.PitchDurs[0]
		p := pd.Pitch

		key, err := Key(p)
		if err != nil {
			err = fmt.Errorf("failed to get MIDI note for pitch '%s': %w", p.String(), err)

			return []*Event{}, err
		}

		dynamic, err := function.At(dc, n.Start)
		if err != nil {
			err = fmt.Errorf("failed compute dynamic at note start %s: %w", n.Start, err)

			return []*Event{}, err
		}

		vel, err := Velocity(n.Attack * dynamic)
		if err != nil {
			err = fmt.Errorf("failed to get MIDI velocity: %w", err)

			return []*Event{}, err
		}

		newDur, err := model.AdjustDuration(pd.Duration, n.Separation)
		if err != nil {
			err = fmt.Errorf("failed to adjust duration: %w", err)

			return []*Event{}, err
		}

		endOffset := new(big.Rat).Add(n.Start, newDur)

		events = append(events, NewNoteOnEvent(n.Start, key, vel))
		events = append(events, NewNoteOffEvent(endOffset, key))
	}

	return events, nil
}
