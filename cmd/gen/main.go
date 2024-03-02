package main

import (
	CustomCommand "sre-tool/app/command"
	_ "sre-tool/internal/bootstrap"
	"sre-tool/internal/command"
)

func main() {
	cmd := command.New()
	cmd.AddCommand(CustomCommand.NewCommand(cmd.Root())).Execute()
}
