package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"
)

type Item interface {
	Name() string
	MakeUIObject() fyne.CanvasObject
}

type ItemFormHelper interface {
	MakeItem() Item
	FormItems() []*widget.FormItem
}

type ItemUpdate struct {
	Index int
	Item  Item
}

type MakeItemFormHelperFunc func(im *ItemManager, item Item) ItemFormHelper

type ItemManager struct {
	mainWindow         fyne.Window
	itemName           string
	items              []Item
	itemsBox           *fyne.Container
	addItem            chan Item
	removeItem         chan string
	updateItem         chan *ItemUpdate
	makeItemFormHelper MakeItemFormHelperFunc
}

func NewItemManager(mainWindow fyne.Window, itemName string, makeItemFormHelper MakeItemFormHelperFunc) *ItemManager {
	return &ItemManager{
		itemName:           itemName,
		items:              []Item{},
		itemsBox:           container.NewVBox(),
		addItem:            make(chan Item),
		removeItem:         make(chan string),
		updateItem:         make(chan *ItemUpdate),
		mainWindow:         mainWindow,
		makeItemFormHelper: makeItemFormHelper,
	}
}

func (im *ItemManager) Monitor() {
	go im.monitor()
}

func (im *ItemManager) BuildTab() *container.TabItem {
	scroll := container.NewVScroll(im.itemsBox)
	buttons := container.NewHBox(
		widget.NewButton(fmt.Sprintf("Add %s", im.itemName), func() {
			im.ShowAddItemDialog()
		}),
	)
	outer := container.NewVSplit(buttons, scroll)

	// Give all available space to the bottom split element
	outer.SetOffset(0.0)

	return container.NewTabItem(fmt.Sprintf("%ss", im.itemName), outer)
}

func (im *ItemManager) HasItem(name string) bool {
	return im.ItemIndex(name) != -1
}

func (im *ItemManager) RemoveItem(idx int) {
	im.items = append(im.items[:idx], im.items[idx+1:]...)
}

func (im *ItemManager) ItemIndex(name string) int {
	for i, item := range im.items {
		if item.Name() == name {
			return i
		}
	}

	return -1
}

func (im *ItemManager) ShowAddItemDialog() {
	h := im.makeItemFormHelper(im, nil)
	cb := func(ok bool) {
		if ok {
			item := h.MakeItem()

			// log.Info().Interface("item", item).Msg("adding item")

			im.addItem <- item
		}
	}
	title := fmt.Sprintf("Add %s", im.itemName)

	dialog.ShowForm(title, "Create", "Cancel", h.FormItems(), cb, im.mainWindow)
}

func (im *ItemManager) ShowEditItemDialog(item Item) {
	h := im.makeItemFormHelper(im, item)
	cb := func(ok bool) {
		if ok {
			updatedItem := h.MakeItem()

			// log.Info().Interface("item", item).Msg("adding item")

			update := &ItemUpdate{
				Index: im.ItemIndex(item.Name()),
				Item:  updatedItem,
			}
			im.updateItem <- update
		}
	}
	title := fmt.Sprintf("Edit %s", im.itemName)

	dialog.ShowForm(title, "Modify", "Cancel", h.FormItems(), cb, im.mainWindow)
}

func (im *ItemManager) monitor() {
	for {
		select {
		case item := <-im.addItem:
			im.items = append(im.items, item)

			itemUI := im.createItemUI(item)

			im.itemsBox.Add(itemUI)
		case update := <-im.updateItem:
			im.items[update.Index] = update.Item

			itemUI := im.createItemUI(update.Item)

			im.itemsBox.Objects[update.Index] = itemUI

			im.itemsBox.Refresh()
		case name := <-im.removeItem:
			if idx := im.ItemIndex(name); idx != -1 {
				im.RemoveItem(idx)
				im.itemsBox.Remove(im.itemsBox.Objects[idx])
			}

		}
	}
}

func (im *ItemManager) createItemUI(item Item) fyne.CanvasObject {
	itemUI := container.NewVBox(item.MakeUIObject())

	editButton := widget.NewButton("Edit", func() {
		log.Debug().Str("name", item.Name()).Msg("editing item")

		im.ShowEditItemDialog(item)
	})
	deleteButton := widget.NewButton("Delete", func() {
		log.Debug().Str("name", item.Name()).Msg("removing item")

		im.removeItem <- item.Name()
	})
	buttons := container.NewHBox(editButton, deleteButton)

	itemUI.Add(buttons)

	return itemUI
}
