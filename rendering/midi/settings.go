package midi

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/score"
)

type MIDISettings struct {
	PartChannels      map[string]uint8 `json:"partChannels"`
	TempoSamplePeriod *big.Rat         `json:"tempoSamplePeriod"`
}

func DefaultTempoSamplePeriod() *big.Rat {
	return big.NewRat(1, 4)
}

func Settings(s *score.Score) (*MIDISettings, error) {
	var midiSettings MIDISettings

	obj, found := s.Settings["midi"]
	if found {
		d, err := json.Marshal(obj)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal MIDI settings obj: %w", err)
		}

		err = json.Unmarshal(d, &midiSettings)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal as MIDI settings: %w", err)
		}
	} else {
		partChannels := map[string]uint8{}
		channel := uint8(1)

		for _, sec := range s.Sections {
			for _, partName := range sec.PartNames() {
				if _, found := partChannels[partName]; !found {
					partChannels[partName] = channel

					switch channel {
					case 9: // reserve channel 10 for percussion
						channel = 11
					case 16: // wrap around once last channel is reached
						channel = 1
					default:
						channel++
					}
				}
			}
		}

		midiSettings = MIDISettings{
			PartChannels:      partChannels,
			TempoSamplePeriod: DefaultTempoSamplePeriod(),
		}
	}

	return &midiSettings, nil
}
