package midi

import (
	"math/big"
	"sort"

	"gitlab.com/gomidi/midi/writer"
)

type Event struct {
	Offset *big.Rat
	Writer EventWriter
}

type EventWriter interface {
	Write(wr *writer.SMF) error
}

func NewEvent(offset *big.Rat, ew EventWriter) *Event {
	return &Event{
		Offset: offset,
		Writer: ew,
	}
}

func (e *Event) Write(wr *writer.SMF) error {
	return e.Writer.Write(wr)
}

// Sort by offset, keeping original order or equal elements.
func SortEvents(events []*Event) {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Offset.Cmp(events[j].Offset) == -1
	})
}
