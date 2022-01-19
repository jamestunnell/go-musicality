package midi

import (
	"math/big"

	"gitlab.com/gomidi/midi/writer"
)

type NoteOnWriter struct {
	key, velocity uint8
}

type NoteOffWriter struct {
	key uint8
}

// NewNoteOnEvent makes a new note on event.
// Returns a non-nil error if the pitch is not in range for MIDI.
func NewNoteOnEvent(offset *big.Rat, key uint8, velocity uint8) *Event {
	return NewEvent(offset, &NoteOnWriter{key: key, velocity: velocity})
}

// NewNoteOffEvent makes a new note off event.
// Returns a non-nil error if the pitch is not in range for MIDI.
func NewNoteOffEvent(offset *big.Rat, key uint8) *Event {
	return NewEvent(offset, &NoteOffWriter{key: key})
}

func (e *NoteOnWriter) Write(wr *writer.SMF) error {
	return writer.NoteOn(wr, e.key, e.velocity)
}

func (e *NoteOffWriter) Write(wr *writer.SMF) error {
	return writer.NoteOff(wr, e.key)
}
