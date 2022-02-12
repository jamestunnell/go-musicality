package midi

import (
	"fmt"
	"math"

	"github.com/jamestunnell/go-musicality/performance/function"
	"github.com/jamestunnell/go-musicality/performance/model"
)

func CollectNoteEvents(fs *model.FlatScore, dc *model.Computer, part string) ([]Event, error) {
	events := []Event{}

	notes := fs.Parts[part]
	converter := model.NewNoteConverter(model.OptionReplaceSlursAndGlides())

	notes2, err := converter.Process(notes)
	if err != nil {
		return []Event{}, fmt.Errorf("failed to convert notes: %w", err)
	}

	for i, n := range notes2 {
		if len(n.PitchDurs) > 1 {
			return []Event{}, fmt.Errorf("note %d has multiple pitch durs", i)
		}

		pd := n.PitchDurs[0]
		p := pd.Pitch

		key, err := Key(p)
		if err != nil {
			err = fmt.Errorf("failed to get MIDI note for pitch '%s': %w", p.String(), err)

			return []Event{}, err
		}

		dynamic, err := function.At(dc, n.Start)
		if err != nil {
			err = fmt.Errorf("failed compute dynamic at note start %s: %w", n.Start, err)

			return []Event{}, err
		}

		vel := Velocity(n.Attack * dynamic)
		newDur := model.AdjustDuration(pd.Duration, n.Separation)
		endOffset := n.Start.Add(newDur)

		events = append(events, NewNoteOnEvent(n.Start, key, vel))
		events = append(events, NewNoteOffEvent(endOffset, key))
	}

	return events, nil
}

// Key converts the pitch to a MIDI note number.
// Returns a non-nil error if the pitch is not in the range [C-1, G9].
func Key(p *model.Pitch) (uint8, error) {
	const (
		// minTotalSemitone is the total semitone value of MIDI note 0 (octave below C0)
		minTotalSemitone = -12
		// maxTotalSemitone is the total semitone value of MIDI note 127 (G9)
		maxTotalSemitone = 115
	)

	totalSemitone := p.TotalSemitone()

	if totalSemitone < minTotalSemitone || totalSemitone > maxTotalSemitone {
		return 0, fmt.Errorf("pitch %s is outside of MIDI note number range", p.String())
	}

	return uint8(totalSemitone + 12), nil
}

// Velocity converts the attack to a MIDI velocity.
// Returns a non-nil error if the attack is not in the range [0.0, 1.0]
func Velocity(attack float64) uint8 {
	return 31 + uint8(math.Round(attack*96))
}
