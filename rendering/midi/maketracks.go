package midi

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/performance/flatscore"
)

func MakeTracks(fs *flatscore.FlatScore, settings *MIDISettings) ([]*Track, error) {
	for partName := range fs.Parts {
		if _, found := settings.PartChannels[partName]; !found {
			return []*Track{}, fmt.Errorf("part '%s' channel is missing from MIDI settings", partName)
		}
	}

	tracks := []*Track{}

	for partName := range fs.Parts {
		noteEvents, err := CollectNoteEvents(fs, fs.DyamicComputer, partName)
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
