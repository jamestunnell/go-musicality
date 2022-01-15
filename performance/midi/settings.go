package midi

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jamestunnell/go-musicality/notation/score"
)

type MIDISettings struct {
	PartChannels map[string]uint8 `json:"partChannels"`
}

var errMIDISettingsNotFound = errors.New("MIDI settings not found")

func Settings(s *score.Score) (*MIDISettings, error) {
	obj, found := s.Settings["midi"]
	if !found {
		return nil, errMIDISettingsNotFound
	}

	d, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal MIDI settings obj: %w", err)
	}

	var midiSettings MIDISettings

	err = json.Unmarshal(d, &midiSettings)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal as MIDI settings: %w", err)
	}

	return &midiSettings, nil
}
