// Package convert implements 'cointool convert' command.
package convert

import (
	"github.com/spf13/cobra"

	"github.com/vdrpkv/cointool/cmd/cli/cointool/command"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"

	"github.com/vdrpkv/cointool/internal/infrastructure/coinmarketcap"

	convertCliController "github.com/vdrpkv/cointool/internal/controller/cli/convert"
	convertUseCase "github.com/vdrpkv/cointool/internal/domain/usecase/currency/convert"
	exchangeRateUseCase "github.com/vdrpkv/cointool/internal/domain/usecase/currency/exchangerate"
)

func NewCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "convert amount symbol-from symbol-to",
		Short: "Convert currency",
		Run: func(cmd *cobra.Command, args []string) {
			command.ExecCliController(
				cmd, args,
				convertCliController.New(
					convertUseCase.NewUseCaseConvertCurrency(
						exchangeRateUseCase.NewUseCaseGetExchangeRate(
							coinmarketcap.NewClient(
								coinmarketcap.APIKey(variable.ApiKey),
								coinmarketcap.Environment(variable.ApiPrefix),
							),
							variable.ExchangeRateZeroValue,
						),
					),
				),
			)
		},
	}
}
