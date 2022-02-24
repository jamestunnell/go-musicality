package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type Item interface {
	Name() string
	MakeUIObject() fyne.CanvasObject
}

type ItemFormHelper interface {
	MakeItem() Item
	FormItems() []*widget.FormItem
}

type MakeItemFormHelperFunc func(m *ItemManager) ItemFormHelper

type ItemManager struct {
	mainWindow         fyne.Window
	itemName           string
	items              []Item
	itemsBox           *fyne.Container
	addItem            chan Item
	makeItemFormHelper MakeItemFormHelperFunc
}

func NewItemManager(mainWindow fyne.Window, itemName string, makeItemFormHelper MakeItemFormHelperFunc) *ItemManager {
	return &ItemManager{
		itemName:           itemName,
		items:              []Item{},
		itemsBox:           container.NewVBox(),
		addItem:            make(chan Item),
		mainWindow:         mainWindow,
		makeItemFormHelper: makeItemFormHelper,
	}
}

func (pm *ItemManager) Monitor() {
	go pm.monitor()
}

func (pm *ItemManager) BuildTab() *container.TabItem {
	scroll := container.NewVScroll(pm.itemsBox)
	buttons := container.NewHBox(
		widget.NewButton(fmt.Sprintf("Add %s", pm.itemName), func() {
			pm.ShowAddItemDialog()
		}),
	)
	outer := container.NewVSplit(buttons, scroll)

	// Give all available space to the bottom split element
	outer.SetOffset(0.0)

	return container.NewTabItem(fmt.Sprintf("%ss", pm.itemName), outer)
}

func (pm *ItemManager) HasItem(name string) bool {
	for _, item := range pm.items {
		if item.Name() == name {
			return true
		}
	}

	return false
}

func (m *ItemManager) ShowAddItemDialog() {
	h := m.makeItemFormHelper(m)
	cb := func(ok bool) {
		if ok {
			item := h.MakeItem()

			// log.Info().Interface("item", item).Msg("adding item")

			m.addItem <- item
		}
	}
	title := fmt.Sprintf("Add %s", m.itemName)

	dialog.ShowForm(title, "Create", "Cancel", h.FormItems(), cb, m.mainWindow)
}

func (pm *ItemManager) monitor() {
	for {
		item := <-pm.addItem

		pm.items = append(pm.items, item)

		itemUI := container.NewVBox(item.MakeUIObject())

		// editButton := widget.NewButton("Edit", func() {

		// })
		deleteButton := widget.NewButton("Delete", func() {
			pm.itemsBox.Remove(itemUI)
		})
		buttons := container.NewHBox(deleteButton)

		itemUI.Add(buttons)

		pm.itemsBox.Add(itemUI)
	}
}
