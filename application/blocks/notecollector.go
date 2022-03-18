package blocks

import (
	"github.com/jamestunnell/go-musicality/application/block"
	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/common/value"
	"github.com/jamestunnell/go-musicality/notation/note"
)

type NoteCollector struct {
	Collected []*note.Note
	Notes     *block.Port
}

func NewNoteCollector() *NoteCollector {
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

func (b *NoteCollector) Initialize(paramVals value.Map) error {
	b.Collected = []*note.Note{}

	return nil
}

func (b *NoteCollector) Configure(controlVals value.Map) {
}

func (b *NoteCollector) Process(offset rat.Rat) {
	n := b.Notes.CurrentValue.(*note.Note)

	if n != nil {
		b.Collected = append(b.Collected, n)
	}
}
