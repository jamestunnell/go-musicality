package midi

import (
	"sort"

	"github.com/jamestunnell/go-musicality/notation/rat"
	"gitlab.com/gomidi/midi/writer"
)

type Event struct {
	Offset rat.Rat
	Writer EventWriter
}

type EventWriter interface {
	Summary() string
	Write(wr *writer.SMF) error
}

func NewEvent(offset rat.Rat, ew EventWriter) *Event {
	return &Event{
		Offset: offset,
		Writer: ew,
	}
}

func (e *Event) Write(wr *writer.SMF) error {
	return e.Writer.Write(wr)
}

// Sort by offset, keeping original order for equal elements.
func SortEvents(events []*Event) {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Offset.Less(events[j].Offset)
	})
}
