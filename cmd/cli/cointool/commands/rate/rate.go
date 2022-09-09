package rate

import (
	"coinconv/cmd/cli/cointool/variables"
	"coinconv/internal/coinmarketcap"
	"coinconv/internal/currency"
	"coinconv/internal/handlers"
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "rate symbol-from symbol-to",
	Short: "get coin exchange rate",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		rate, err := run(args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		fmt.Println(rate)
	},
}

func run(args []string) (currency.ExchangeRate, error) {
	rate, err := handlers.HandleGetExchangeRate(
		context.Background(),
		coinmarketcap.NewExchangeRateGetter(
			variables.ApiKey, variables.ApiPrefix,
		),
		currency.Symbol(args[0]),
		currency.Symbol(args[1]),
	)

	if err != nil {
		return 0, err
	}

	return rate, nil
}
