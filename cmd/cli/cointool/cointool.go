package cointool

import (
	"coinconv/cmd/cli/cointool/variables"
	"time"

	"github.com/spf13/cobra"
)

var (
	Command = &cobra.Command{
		Use: "cointool",
	}
)

func init() {
	Command.PersistentFlags().StringVarP(
		&variables.ApiKey,
		"api-key", "K", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
		"coinmarketcap API key",
	)

	Command.PersistentFlags().StringVarP(
		&variables.ApiPrefix,
		"api-prefix", "P", "sandbox",
		"coinmarketcap API prefix: sandbox or pro",
	)

	Command.PersistentFlags().DurationVarP(
		&variables.Timeout,
		"timeout", "t", 7*time.Second, "command timeout duration",
	)
}
