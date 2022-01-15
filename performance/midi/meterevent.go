package midi

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"gitlab.com/gomidi/midi/writer"
)

type MeterWriter struct {
	met *meter.Meter
}

func NewMeterEvent(offset *big.Rat, met *meter.Meter) *Event {
	return NewEvent(offset, &MeterWriter{met: met})
}

func (e *MeterWriter) Write(wr *writer.SMF) error {
	return writer.Meter(wr, e.met.Numerator, e.met.Denominator)
}
