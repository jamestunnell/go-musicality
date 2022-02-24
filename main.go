package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/jamestunnell/go-musicality/application/ui"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Musicality")

	partManager := ui.NewPartManager(mainWindow)

	partManager.Monitor()

	music := container.NewVBox()
	musicScroll := container.NewVScroll(music)
	musicButtons := container.NewHBox(
		widget.NewButton("Play Parts", func() {
			dialog.ShowInformation("Play Parts", "This is a placeholder", mainWindow)
		}),
	)
	musicOuter := container.NewVSplit(musicButtons, musicScroll)

	// Give all available space to the bottom split element
	musicOuter.SetOffset(0.0)

	appTabs := container.NewAppTabs(
		partManager.BuildPartsTab(),
		container.NewTabItem("Music", musicOuter),
	)

	mainWindow.SetContent(appTabs)
	mainWindow.Resize(fyne.NewSize(320, 240))
	mainWindow.ShowAndRun()
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
