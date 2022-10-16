// Package rate implements 'cointool rate' command.
package rate

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"

	"github.com/vdrpkv/cointool/internal/infrastructure/coinmarketcap"

	exchangeRateCliController "github.com/vdrpkv/cointool/internal/controller/cli/exchangerate"
	exchangeRateUseCase "github.com/vdrpkv/cointool/internal/domain/usecase/exchangerate"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "rate symbol-from symbol-to",
		Short: "Get exchange rate",
		Run: func(cmd *cobra.Command, args []string) {
			command.ExecCliController(
				cmd, args,
				exchangeRateCliController.New(
					exchangeRateUseCase.NewUseCaseGetExchangeRate(
						coinmarketcap.NewClient(
							coinmarketcap.APIKey(variable.ApiKey),
							coinmarketcap.Environment(variable.ApiPrefix),
						),
						variable.ExchangeRateZeroValue,
					),
				),
			)
		},
	}
}
