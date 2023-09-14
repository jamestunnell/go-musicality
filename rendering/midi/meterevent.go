package midi

import (
	"fmt"
	"math/big"

	"gitlab.com/gomidi/midi/writer"
)

type MeterWriter struct {
	num, denom uint8
}

func NewMeterEvent(offset *big.Rat, num, denom uint8) *SMFEvent {
	return NewSMFEvent(offset, &MeterWriter{num: num, denom: denom})
}

func (e *MeterWriter) Summary() string {
	return fmt.Sprintf("meter change(num=%d, denom=%d)", e.num, e.denom)
}

func (e *MeterWriter) Write(wr *writer.SMF) error {
	return writer.Meter(wr, e.num, e.denom)
}
