package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type PartInfo struct {
	name           string
	MIDIChannel    int
	MIDIInstrument int
}

func (pi *PartInfo) Name() string {
	return pi.name
}

func (pi *PartInfo) MakeUIObject() fyne.CanvasObject {
	return widget.NewForm(
		widget.NewFormItem("Name", widget.NewLabel(pi.name)),
		widget.NewFormItem("MIDI Channel", widget.NewLabel(strconv.Itoa(pi.MIDIChannel))),
		widget.NewFormItem("MIDI Instrument", widget.NewLabel(strconv.Itoa(pi.MIDIInstrument))),
	)
}
