package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/rs/zerolog/log"

	"github.com/jamestunnell/go-musicality/application"
	"github.com/jamestunnell/go-musicality/application/ui"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Musicality")

	partsBox := container.NewVBox()
	partManager := ui.NewPartManager(partsBox)

	partManager.Monitor()

	partsScroll := container.NewVScroll(partsBox)
	partsButtons := container.NewHBox(
		widget.NewButton("Add Part", func() {
			showAddPartDialog(mainWindow, partManager)
		}),
	)
	partsOuter := container.NewVSplit(partsButtons, partsScroll)

	music := container.NewVBox()
	musicScroll := container.NewVScroll(music)
	musicButtons := container.NewHBox(
		widget.NewButton("Play Parts", func() {
			dialog.ShowInformation("Play Parts", "This is a placeholder", mainWindow)
		}),
	)
	musicOuter := container.NewVSplit(musicButtons, musicScroll)

	appTabs := container.NewAppTabs(
		container.NewTabItem("Parts", partsOuter),
		container.NewTabItem("Music", musicOuter),
	)

	mainWindow.SetContent(appTabs)
	mainWindow.Resize(fyne.NewSize(320, 240))
	mainWindow.ShowAndRun()
}

func showAddPartDialog(parent fyne.Window, pm *ui.PartManager) {
	nameEntry := widget.NewEntry()
	nameEntry.Validator = func(s string) error {
		if len(s) == 0 {
			return fmt.Errorf("name is empty")
		}

		if pm.HasPart(s) {
			return fmt.Errorf("part '%s' already exists", s)
		}

		return nil
	}
	nameItem := widget.NewFormItem("Name", nameEntry)
	formItems := []*widget.FormItem{
		nameItem,
		widget.NewFormItem("MIDI Channel", widget.NewEntry()),
		widget.NewFormItem("MIDI Instrument", widget.NewEntry()),
	}
	cb := func(added bool) {
		if !added {
			log.Info().Msg("canceled add part")

			return
		}

		partInfo := &application.PartInfo{
			Name: nameEntry.Text,
		}

		log.Info().Msg("adding part")

		pm.AddParts() <- partInfo

		log.Info().Msg("added part")
	}

	dialog.ShowForm("Add Part", "Create", "Cancel", formItems, cb, parent)
}

// func main() {
// 	if len(os.Args) == 1 {
// 		fmt.Println("no command")
// 		os.Exit(1)
// 	}

// 	if len(os.Args) == 2 {
// 		fmt.Println("no subcommand")
// 		os.Exit(1)
// 	}

// 	var cmd commands.Command
// 	var err error

// 	cmdName := os.Args[1]
// 	subCmdName := os.Args[2]
// 	args := os.Args[3:]

// 	switch cmdName {
// 	case "render":
// 		if cmd, err = render.NewSubcommand(subCmdName, args...); err != nil {
// 			fmt.Printf("%v\n", err)

// 			os.Exit(1)
// 		}
// 	case "generate":
// 		if cmd, err = generate.NewSubcommand(subCmdName, args...); err != nil {
// 			fmt.Printf("%v\n", err)

// 			os.Exit(1)
// 		}
// 	case "play":
// 		if cmd, err = play.NewSubcommand(subCmdName, args...); err != nil {
// 			fmt.Printf("%v\n", err)

// 			os.Exit(1)
// 		}
// 	case "validate":
// 		if cmd, err = validate.NewSubcommand(subCmdName, args...); err != nil {
// 			fmt.Printf("%v\n", err)

// 			os.Exit(1)
// 		}
// 	default:
// 		fmt.Printf("unknown command '%s'", cmdName)

// 		os.Exit(1)
// 	}

// 	if err := cmd.Execute(); err != nil {
// 		fmt.Printf("%s command failed: %v\n", cmd.Name(), err)

// 		os.Exit(1)
// 	}

// 	os.Exit(0)
// }
