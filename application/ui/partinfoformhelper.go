package ui

import (
	"strconv"

	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
)

type PartInfoFormHelper struct {
	NameEntry, MIDIChanEntry, MIDIInstrEntry *widget.Entry
	formItems                                []*widget.FormItem
}

func NewPartInfoFormHelper(m *ItemManager) ItemFormHelper {
	nameEntry := widget.NewEntry()
	midiChanEntry := widget.NewEntry()
	midiInstrEntry := widget.NewEntry()

	nameEntry.Validator = MakeNameValidator(m)
	midiChanEntry.Validator = MakeIntValidator(1, 16)
	midiInstrEntry.Validator = MakeIntValidator(1, 128)

	nameItem := widget.NewFormItem("Name", nameEntry)
	midiChanItem := widget.NewFormItem("MIDI Channel", midiChanEntry)
	midiInstrItem := widget.NewFormItem("MIDI Instrument", midiInstrEntry)

	return &PartInfoFormHelper{
		NameEntry:      nameEntry,
		MIDIChanEntry:  midiChanEntry,
		MIDIInstrEntry: midiInstrEntry,
		formItems:      []*widget.FormItem{nameItem, midiChanItem, midiInstrItem},
	}
}

func (h *PartInfoFormHelper) FormItems() []*widget.FormItem {
	return h.formItems
}

func (h *PartInfoFormHelper) MakeItem() Item {
	c, err := strconv.Atoi(h.MIDIChanEntry.Text)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to convert MIDI channel text '%s' to int", h.MIDIChanEntry.Text)
	}

	i, err := strconv.Atoi(h.MIDIInstrEntry.Text)
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to convert MIDI instrument text '%s' to int", h.MIDIInstrEntry.Text)
	}

	return &PartInfo{
		name:           h.NameEntry.Text,
		MIDIChannel:    c,
		MIDIInstrument: i,
	}
}
