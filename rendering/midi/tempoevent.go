package midi

import (
	"fmt"
	"math/big"

	"gitlab.com/gomidi/midi/writer"
)

type TempoWriter struct {
	bpm float64
}

func NewTempoEvent(offset *big.Rat, bpm float64) *SMFEvent {
	return NewSMFEvent(offset, &TempoWriter{bpm: bpm})
}

func (e *TempoWriter) Summary() string {
	return fmt.Sprintf("tempo change(bpm=%v)", e.bpm)
}

func (e *TempoWriter) Write(wr *writer.SMF) error {
	return writer.TempoBPM(wr, e.bpm)
}
