package midi

import (
	"fmt"
	"math/big"

	"gitlab.com/gomidi/midi/writer"
)

type NoteEvent struct {
	offset      *big.Rat
	eventWriter NoteEventWriter
}

type NoteEventWriter interface {
	Summary() string
	Write(wr writer.ChannelWriter) error
}

type NoteOnWriter struct {
	key, velocity uint8
}

type NoteOffWriter struct {
	key uint8
}

func NewNoteEvent(offset *big.Rat, ew NoteEventWriter) *NoteEvent {
	return &NoteEvent{
		offset:      offset,
		eventWriter: ew,
	}
}

func (e *NoteEvent) Offset() *big.Rat {
	return e.offset
}

func (e *NoteEvent) Summary() string {
	return e.eventWriter.Summary()
}

func (e *NoteEvent) Write(wr writer.ChannelWriter) error {
	return e.eventWriter.Write(wr)
}

// NewNoteOnEvent makes a new note on event.
// Returns a non-nil error if the pitch is not in range for MIDI.
func NewNoteOnEvent(offset *big.Rat, key uint8, velocity uint8) *NoteEvent {
	return NewNoteEvent(offset, &NoteOnWriter{key: key, velocity: velocity})
}

// NewNoteOffEvent makes a new note off event.
// Returns a non-nil error if the pitch is not in range for MIDI.
func NewNoteOffEvent(offset *big.Rat, key uint8) *NoteEvent {
	return NewNoteEvent(offset, &NoteOffWriter{key: key})
}

func (e *NoteOnWriter) Write(wr writer.ChannelWriter) error {
	return writer.NoteOn(wr, e.key, e.velocity)
}

func (e *NoteOnWriter) Summary() string {
	return fmt.Sprintf("note on(key=%d, vel=%d)", e.key, e.velocity)
}

func (e *NoteOffWriter) Summary() string {
	return fmt.Sprintf("note off(key=%d)", e.key)
}

func (e *NoteOffWriter) Write(wr writer.ChannelWriter) error {
	return writer.NoteOff(wr, e.key)
}
