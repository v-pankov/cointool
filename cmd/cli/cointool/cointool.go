package cointool

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/variables"
)

var (
	Command = &cobra.Command{
		Use: "cointool",
		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}
)

func init() {
	Command.PersistentFlags().StringVarP(
		&variables.ApiKey,
		"api-key", "k", "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c",
		"coinmarketcap API key",
	)

	Command.PersistentFlags().StringVarP(
		&variables.ApiPrefix,
		"api-prefix", "p", "sandbox",
		"coinmarketcap API prefix: sandbox or pro",
	)

	Command.PersistentFlags().DurationVarP(
		&variables.Timeout,
		"timeout", "t", 7*time.Second, "command timeout duration",
	)
}
