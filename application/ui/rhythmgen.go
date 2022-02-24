package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/jamestunnell/go-musicality/common/rat"
	"github.com/jamestunnell/go-musicality/notation/meter"
)

type RhythmGen struct {
	name        string
	Meter       *meter.Meter
	SmallestDur rat.Rat
}

func (rg *RhythmGen) Name() string {
	return rg.name
}

func (rg *RhythmGen) MakeUIObject() fyne.CanvasObject {
	return widget.NewForm(
		widget.NewFormItem("Name", widget.NewLabel(rg.name)),
		widget.NewFormItem("Meter", widget.NewLabel(rg.Meter.String())),
		widget.NewFormItem("SmallestDur", widget.NewLabel(rg.SmallestDur.String())),
	)
}
