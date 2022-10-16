// Package rate implements 'cointool rate' command.
package fiat

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"

	"github.com/vdrpkv/cointool/internal/infrastructure/coinmarketcap"

	fiatCliController "github.com/vdrpkv/cointool/internal/controller/cli/fiat"
	fiatUseCase "github.com/vdrpkv/cointool/internal/domain/usecase/fiat"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "fiat symbol",
		Short: "Recognize fiat currency symbol",
		Run: func(cmd *cobra.Command, args []string) {
			command.ExecCliController(
				cmd, args,
				fiatCliController.New(
					fiatUseCase.NewUseCaseRecognizeFiatCurrency(
						coinmarketcap.NewClient(
							coinmarketcap.APIKey(variable.ApiKey),
							coinmarketcap.Environment(variable.ApiPrefix),
						),
					),
				),
			)
		},
	}
}
