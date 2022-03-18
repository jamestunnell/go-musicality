package notecollector

import (
	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/jamestunnell/go-musicality/notation/note"
)

type NoteCollector struct {
	Collected []*note.Note
	Notes     *block.Port
}

const NotesName = "Notes"

func New() *NoteCollector {
	var n *note.Note

	return &NoteCollector{
		Collected: []*note.Note{},
		Notes:     block.NewInput(n),
	}
}

func (b *NoteCollector) Params() map[string]*block.Param {
	return map[string]*block.Param{}
}

func (b *NoteCollector) Ports() map[string]*block.Port {
	return map[string]*block.Port{
		NotesName: b.Notes,
	}
}

func (b *NoteCollector) Initialize() error {
	b.Collected = []*note.Note{}

	return nil
}

func (b *NoteCollector) Configure() {
}

func (b *NoteCollector) Process() {
	n := b.Notes.CurrentValue.(*note.Note)

	if n != nil {
		b.Collected = append(b.Collected, n)
	}
}
