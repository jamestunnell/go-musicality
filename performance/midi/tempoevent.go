package midi

import (
	"fmt"

	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

type TempoWriter struct {
	bpm float64
}

func NewTempoEvent(offset rat.Rat, bpm float64) *Event {
	return NewEvent(offset, &TempoWriter{bpm: bpm})
}

func (e *TempoWriter) Summary() string {
	return fmt.Sprintf("tempo change(bpm=%v)", e.bpm)
}

func (e *TempoWriter) Write(wr *writer.SMF) error {
	return writer.TempoBPM(wr, e.bpm)
}