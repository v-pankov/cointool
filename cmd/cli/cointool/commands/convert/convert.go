package convert

import (
	"coinconv/cmd/cli/cointool/variables"
	"coinconv/internal/coinmarketcap"
	"coinconv/internal/currency"
	"coinconv/internal/handlers"
	"context"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var Command = &cobra.Command{
	Use:   "convert amount symbol-from symbol-to",
	Short: "convert coins",
	Args:  cobra.ExactArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, ctxCancel := context.WithTimeout(
			cmd.Context(), variables.Timeout,
		)
		defer ctxCancel()

		amount, err := run(ctx, args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		fmt.Println(amount)
	},
}

func run(ctx context.Context, args []string) (currency.Amount, error) {
	argAmount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, fmt.Errorf("invalid amount: %w", err)
	}

	var (
		argFrom = args[1]
		argTo   = args[2]
	)

	amount, err := handlers.HandleConvertCurrency(
		ctx,
		coinmarketcap.NewFiatCurrencyRecognizer(
			variables.ApiKey, variables.ApiPrefix,
		),
		coinmarketcap.NewExchangeRateGetter(
			variables.ApiKey, variables.ApiPrefix,
		),
		currency.Amount(argAmount),
		currency.Symbol(argFrom),
		currency.Symbol(argTo),
	)

	if err != nil {
		return 0, err
	}

	return amount, nil
}
