package midi

import (
	"errors"

	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/common/rat"
)

type SMFEvent struct {
	offset      rat.Rat
	eventWriter SMFEventWriter
}

type SMFEventWriter interface {
	Summary() string
	Write(wr *writer.SMF) error
}

var errNotSMFWriter = errors.New("not an SMF writer")

func NewSMFEvent(offset rat.Rat, ew SMFEventWriter) *SMFEvent {
	return &SMFEvent{
		offset:      offset,
		eventWriter: ew,
	}
}

func (e *SMFEvent) Offset() rat.Rat {
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
