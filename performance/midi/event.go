package midi

import (
	"math/big"
	"sort"

	"gitlab.com/gomidi/midi/writer"
)

type Event interface {
	Offset() *big.Rat
	Write(wr *writer.SMF) error
}

// Sort by offset, keeping original order or equal elements.
func SortEvents(events []Event) {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Offset().Cmp(events[j].Offset()) == -1
	})
}
