package midi

import (
	"errors"
	"math/big"

	"gitlab.com/gomidi/midi/writer"
)

type SMFEvent struct {
	offset      *big.Rat
	eventWriter SMFEventWriter
}

type SMFEventWriter interface {
	Summary() string
	Write(wr *writer.SMF) error
}

var errNotSMFWriter = errors.New("not an SMF writer")

func NewSMFEvent(offset *big.Rat, ew SMFEventWriter) *SMFEvent {
	return &SMFEvent{
		offset:      offset,
		eventWriter: ew,
	}
}

func (e *SMFEvent) Offset() *big.Rat {
	return e.offset
}

func (e *SMFEvent) Summary() string {
	return e.eventWriter.Summary()
}

func (e *SMFEvent) Write(wr writer.ChannelWriter) error {
	wrSMF, ok := wr.(*writer.SMF)
	if !ok {
		return errNotSMFWriter
	}

	return e.eventWriter.Write(wrSMF)
}
