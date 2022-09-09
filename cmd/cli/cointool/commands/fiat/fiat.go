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
		isFiat, err := run(args)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			return
		}
		fmt.Println(isFiat)
	},
}

func run(args []string) (bool, error) {
	isFiat, err := handlers.HandleRecognizeFiatCurrency(
		context.Background(),
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
