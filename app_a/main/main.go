package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"cli-example/app_a"
	"cli-example/commands"
	"cli-example/module"
)

var appA = cli.App{
	Name: "app-a",
	Commands: []*cli.Command{
		&app_a.CmdSpecificToA,
		&commands.CmdShared,
	},
}

func main() {
	module.CommandSetup(appA.Commands)

	_ = appA.Run([]string{"/path/arg", "cmd-specific-to-a"})
	fmt.Println()
	_ = appA.Run([]string{"/path/arg", "cmd-shared"})
}
