package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"github.com/jamestunnell/go-musicality/application/ui"
)

func main() {
	myApp := app.New()
	mainWindow := myApp.NewWindow("Musicality")

	partManager := ui.NewItemManager(mainWindow, "Part", ui.NewPartInfoFormHelper)
	rhythmGenManager := ui.NewItemManager(mainWindow, "Rhythm Generator", ui.NewRhythmGenFormHelper)

	partManager.Monitor()
	rhythmGenManager.Monitor()

	appTabs := container.NewAppTabs(
		partManager.BuildTab(),
		rhythmGenManager.BuildTab(),
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
