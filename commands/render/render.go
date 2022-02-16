package render

import (
	"fmt"

	"github.com/jamestunnell/go-musicality/commands"
	"github.com/jamestunnell/go-musicality/commands/render/midismf"
)

func NewSubcommand(subCmdName string, cliArgs ...string) (commands.Command, error) {
	switch subCmdName {
	case midismf.Name:
		return midismf.NewFromArgs(cliArgs...)
	}

	return nil, fmt.Errorf("unknown subcommand '%s'", subCmdName)
}
