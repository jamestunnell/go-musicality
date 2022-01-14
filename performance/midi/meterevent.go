package midi

import (
	"math/big"

	"github.com/jamestunnell/go-musicality/notation/meter"
	"gitlab.com/gomidi/midi/writer"
)

type MeterEvent struct {
	met    *meter.Meter
	offset *big.Rat
}

func NewMeterEvent(offset *big.Rat, met *meter.Meter) *MeterEvent {
	return &MeterEvent{
		offset: offset,
		met:    met,
	}
}

func (e *MeterEvent) Offset() *big.Rat {
	return e.offset
}

func (e *MeterEvent) Write(wr *writer.SMF) error {
	return writer.Meter(wr, e.met.Numerator, e.met.Denominator)
}
