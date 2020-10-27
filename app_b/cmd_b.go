package app_b

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"cli-example/module"
)

/* LEGACY START */

type ModuleB struct {
	// Application specific data is put on this and must be run against legacy startup code
	Database interface{}
}

func (m *ModuleB) Start() {
	fmt.Println("Start B")
	m.Database = nil // Setup a DB specific to this command
}

func (m *ModuleB) Stop() {
	fmt.Println("Stop B")
}

/* LEGACY STOP */

var CmdSpecificToB = cli.Command{
	Name: "cmd-specific-to-b",
	Metadata: map[string]interface{}{
		module.CLIModuleKey: &ModuleB{},
	},
	Before: func(ctx *cli.Context) error {
		fmt.Println("Before B")
		return nil
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("Action B")
		return nil
	},
	After: func(ctx *cli.Context) error {
		fmt.Println("After B")
		return nil
	},
}
