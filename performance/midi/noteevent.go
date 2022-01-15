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

func NoteOnEvent(offset *big.Rat, key, velocity uint8) *Event {
	w := &NoteOnWriter{key: key, velocity: velocity}

	return NewEvent(offset, w)
}

func NewNoteOffEvent(offset *big.Rat, key uint8) *Event {
	w := &NoteOffWriter{key: key}

	return NewEvent(offset, w)
}

func (e *NoteOnWriter) Write(wr *writer.SMF) error {
	return writer.NoteOn(wr, e.key, e.velocity)
}

func (e *NoteOffWriter) Write(wr *writer.SMF) error {
	return writer.NoteOff(wr, e.key)
}
