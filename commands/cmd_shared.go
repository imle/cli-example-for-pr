package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"cli-example/module"
)

/* LEGACY START */

type ModuleShared struct {
	// Application specific data is put on this and must be run against legacy startup code
	Database interface{}
}

func (m *ModuleShared) Start() {
	fmt.Println("Start Shared")
	m.Database = nil // Setup a DB specific to this command
}

func (m *ModuleShared) Stop() {
	fmt.Println("Stop Shared")
}

func (m *ModuleShared) DoAThing(something ...interface{}) {
	fmt.Println("Doing Thing")
}

var mod = &ModuleShared{}

/* LEGACY STOP */

var CmdShared = cli.Command{
	Name: "cmd-shared",
	Metadata: map[string]interface{}{
		module.CLIModuleKey: mod,
	},
	Before: func(ctx *cli.Context) error {
		fmt.Println("Before Shared")
		return nil
	},
	Action: func(ctx *cli.Context) error {
		fmt.Println("Action Shared")
		mod.DoAThing()
		return nil
	},
	After: func(ctx *cli.Context) error {
		fmt.Println("After Shared")
		return nil
	},
}
