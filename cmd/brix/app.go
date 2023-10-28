package brix

import (
	"fmt"
	"github.com/brix-go/brix/config"
	"github.com/brix-go/brix/internal/command/init"
	"github.com/brix-go/brix/internal/command/run"
	"github.com/spf13/cobra"
)

var text = "A tool for generating golang project"
var coloredText = "\x1B[38;2;66;211;146m" + text + "\x1B[39m"

var CmdRoot = &cobra.Command{
	Use:     "brix",
	Example: "brix new demo-api",
	Short:   fmt.Sprintf("\n                               \n,-----.         ,--.           \n|  |) /_ ,--.--.`--',--.  ,--. \n|  .-.  \\|  .--',--. \\  `'  /  \n|  '--' /|  |   |  | /  /.  \\  \n`------' `--'   `--''--'  '--' \n                               \n%s", coloredText),
	Version: fmt.Sprintf("\n                               \n,-----.         ,--.           \n|  |) /_ ,--.--.`--',--.  ,--. \n|  .-.  \\|  .--',--. \\  `'  /  \n|  '--' /|  |   |  | /  /.  \\  \n`------' `--'   `--''--'  '--' \n                               \nBrix %s - Copyright (c) 2023 Nunu\nReleased under the MIT License.\n\n", config.Version),
}

func init() {
	CmdRoot.AddCommand(new.CmdNew)
	//CmdRoot.AddCommand(create.CmdCreate)
	CmdRoot.AddCommand(run.CmdRun)

	//CmdRoot.AddCommand(upgrade.CmdUpgrade)
	//create.CmdCreate.AddCommand(create.CmdCreateHandler)
	//create.CmdCreate.AddCommand(create.CmdCreateService)
	//create.CmdCreate.AddCommand(create.CmdCreateRepository)
	//create.CmdCreate.AddCommand(create.CmdCreateModel)
	//create.CmdCreate.AddCommand(create.CmdCreateAll)

	//CmdRoot.AddCommand(wire.CmdWire)
	//wire.CmdWire.AddCommand(wire.CmdWireAll)
}

// Execute executes the root command.
func Execute() error {
	return CmdRoot.Execute()
}
