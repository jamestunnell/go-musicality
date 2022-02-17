package midi

import (
	"sort"

	"gitlab.com/gomidi/midi/writer"

	"github.com/jamestunnell/go-musicality/common/rat"
)

type Event interface {
	Offset() rat.Rat
	Write(wr writer.ChannelWriter) error
	Summary() string
}

// Sort by offset, keeping original order for equal elements.
func SortEvents(events []Event) {
	sort.SliceStable(events, func(i, j int) bool {
		return events[i].Offset().Less(events[j].Offset())
	})
}
