package fiat

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
	Use:   "fiat symbol",
	Short: "check is coin fiat or not",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ctx, ctxCancel := context.WithTimeout(
			cmd.Context(), variables.Timeout,
		)
		defer ctxCancel()

		isFiat, err := run(ctx, args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}

		fmt.Println(isFiat)
	},
}

func run(ctx context.Context, args []string) (bool, error) {
	isFiat, err := handlers.HandleRecognizeFiatCurrency(
		ctx,
		coinmarketcap.NewFiatCurrencyRecognizer(
			variables.ApiKey, variables.ApiPrefix,
		),
		currency.Symbol(args[0]),
	)

	if err != nil {
		return false, err
	}

	return isFiat, nil
}
