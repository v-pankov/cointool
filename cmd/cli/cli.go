package main

import (
	"coinconv/cmd/cli/cointool"
	"coinconv/cmd/cli/cointool/commands/convert"
	"coinconv/cmd/cli/cointool/commands/fiat"
	"coinconv/cmd/cli/cointool/commands/rate"
)

func init() {
	cointool.Command.AddCommand(convert.Command)
	cointool.Command.AddCommand(fiat.Command)
	cointool.Command.AddCommand(rate.Command)
}

func main() {
	cointool.Command.Execute()
}
