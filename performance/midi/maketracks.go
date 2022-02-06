package midi

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/score"
	"github.com/jamestunnell/go-musicality/performance/model"
)

func MakeTracks(s *score.Score, settings *MIDISettings) ([]*Track, error) {
	fs, err := (&model.ScoreConverter{}).Process(s)
	if err != nil {
		err = fmt.Errorf("failed to convert to flat score: %w", err)

		return []*Track{}, err
	}

	for partName := range fs.Parts {
		if _, found := settings.PartChannels[partName]; !found {
			return []*Track{}, fmt.Errorf("part '%s' channel is missing from MIDI settings", partName)
		}
	}

	dc, err := fs.DynamicComputer()
	if err != nil {
		err = fmt.Errorf("failed to make dynamic computer: %w", err)

		return []*Track{}, err
	}

	// metEvents, err := collectMeterEvents(fs)
	// if err != nil {
	// 	return []*Track{}, fmt.Errorf("failed to collect meter events: %w", err)
	// }

	tracks := []*Track{}

	for partName := range fs.Parts {
		noteEvents, err := CollectNoteEvents(fs, dc, partName)
		if err != nil {
			return []*Track{}, fmt.Errorf("failed to collect notes for part '%s': %w", partName, err)
		}

		events := []*Event{}

		// events = append(events, metEvents...)

		events = append(events, noteEvents...)

		if len(events) == 0 {
			break
		}

		SortEvents(events)

		track := &Track{
			Name:       partName,
			Channel:    settings.PartChannels[partName],
			Events:     events,
			Instrument: 1,
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}
