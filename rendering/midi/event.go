package midi

import (
	"math/big"
	"sort"

	"github.com/jamestunnell/go-musicality/common/rat"
	"gitlab.com/gomidi/midi/writer"
)

type Event interface {
	Offset() *big.Rat
	Write(wr writer.ChannelWriter) error
	Summary() string
}

// Sort by offset, keeping original order for equal elements.
func SortEvents(events []Event) {
	sort.SliceStable(events, func(i, j int) bool {
		return rat.IsLess(events[i].Offset(), events[j].Offset())
	})
}
