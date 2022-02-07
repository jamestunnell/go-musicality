package midi

import (
	"fmt"

	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/notation/rat"
)

type MeterWriter struct {
	num, denom uint8
}

func NewMeterEvent(offset rat.Rat, num, denom uint8) *Event {
	return NewEvent(offset, &MeterWriter{num: num, denom: denom})
}

func (e *MeterWriter) Summary() string {
	return fmt.Sprintf("meter change(num=%d, denom=%d)", e.num, e.denom)
}

func (e *MeterWriter) Write(wr *writer.SMF) error {
	return writer.Meter(wr, e.num, e.denom)
}
