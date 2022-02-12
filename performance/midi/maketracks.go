package midi

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/performance/model"
)

func MakeTracks(fs *model.FlatScore, settings *MIDISettings) ([]*Track, error) {
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

	tracks := []*Track{}

	for partName := range fs.Parts {
		noteEvents, err := CollectNoteEvents(fs, dc, partName)
		if err != nil {
			return nil, fmt.Errorf("failed to collect notes for part '%s': %w", partName, err)
		}

		if len(noteEvents) == 0 {
			break
		}

		SortEvents(noteEvents)

		track := &Track{
			Name:       partName,
			Channel:    settings.PartChannels[partName],
			Events:     noteEvents,
			Instrument: 1,
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}
