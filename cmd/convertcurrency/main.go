package main

import (
	"coinconv/internal/coinmarketcap"
	"coinconv/internal/handlers"
	"context"
	"fmt"
)

func main() {
	amount, err := handlers.HandleConvertCurrency(
		context.Background(),
		coinmarketcap.NewFiatCurrencyRecognizer(
			"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c", "sandbox",
		),
		coinmarketcap.NewExchangeRateGetter(
			"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c", "sandbox",
		),
		1000,
		"BTC",
		"USD",
	)

	fmt.Printf("%f %v\n", amount, err)
}
