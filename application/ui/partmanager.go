package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/jamestunnell/go-musicality/application"
	"github.com/rs/zerolog/log"
)

type PartManager struct {
	mainWindow fyne.Window
	parts      []*application.PartInfo
	partsBox   *fyne.Container
	addParts   chan *application.PartInfo
}

func NewPartManager(mainWindow fyne.Window) *PartManager {
	return &PartManager{
		parts:      []*application.PartInfo{},
		partsBox:   container.NewVBox(),
		addParts:   make(chan *application.PartInfo),
		mainWindow: mainWindow,
	}
}

func (pm *PartManager) Monitor() {
	go pm.monitor()
}

func (pm *PartManager) BuildPartsTab() *container.TabItem {
	partsScroll := container.NewVScroll(pm.partsBox)
	partsButtons := container.NewHBox(
		widget.NewButton("Add Part", func() {
			pm.ShowAddPartDialog()
		}),
	)
	partsOuter := container.NewVSplit(partsButtons, partsScroll)

	// Give all available space to the bottom split element
	partsOuter.SetOffset(0.0)

	return container.NewTabItem("Parts", partsOuter)
}

func (pm *PartManager) AddParts() chan<- *application.PartInfo {
	return pm.addParts
}

func (pm *PartManager) HasPart(name string) bool {
	for _, p := range pm.parts {
		if p.Name == name {
			return true
		}
	}

	return false
}

func (pm *PartManager) ShowAddPartDialog() {
	x := NewPartInfoFormItems(pm)
	cb := func(ok bool) {
		if ok {
			partInfo := x.GetPartInfo()

			log.Info().Interface("part info", partInfo).Msg("adding part")

			pm.AddParts() <- partInfo
		}
	}

	dialog.ShowForm("Add Part", "Create", "Cancel", x.FormItems, cb, pm.mainWindow)
}

func (pm *PartManager) monitor() {
	for {
		partInfo := <-pm.addParts

		partForm := widget.NewForm(
			widget.NewFormItem("Name", widget.NewLabel(partInfo.Name)),
			widget.NewFormItem("MIDI Channel", widget.NewLabel(strconv.Itoa(partInfo.MIDIChannel))),
			widget.NewFormItem("MIDI Instrument", widget.NewLabel(strconv.Itoa(partInfo.MIDIInstrument))),
		)

		pm.parts = append(pm.parts, partInfo)

		pm.partsBox.Add(partForm)
	}
}
