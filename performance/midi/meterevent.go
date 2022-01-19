package midi

import (
	"math/big"

	"gitlab.com/gomidi/midi/writer"
)

type MeterWriter struct {
	num, denom uint8
}

func NewMeterEvent(offset *big.Rat, num, denom uint8) *Event {
	return NewEvent(offset, &MeterWriter{num: num, denom: denom})
}

func (e *MeterWriter) Write(wr *writer.SMF) error {
	return writer.Meter(wr, e.num, e.denom)
}
