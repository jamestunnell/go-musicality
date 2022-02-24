package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"

	"github.com/jamestunnell/go-musicality/application"
)

type PartManager struct {
	parts    []*application.PartInfo
	partsBox *fyne.Container
	addParts chan *application.PartInfo
}

func NewPartManager(partsBox *fyne.Container) *PartManager {
	return &PartManager{
		parts:    []*application.PartInfo{},
		partsBox: partsBox,
		addParts: make(chan *application.PartInfo),
	}
}

func (pm *PartManager) Monitor() {
	go pm.monitor()
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

func (pm *PartManager) monitor() {
	for {
		partInfo := <-pm.addParts

		partForm := widget.NewForm(
			widget.NewFormItem("Name", widget.NewLabel(partInfo.Name)),
		)

		pm.parts = append(pm.parts, partInfo)

		pm.partsBox.Add(partForm)
	}
}
