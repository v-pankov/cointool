package main

import (
	"coinconv/internal/coinmarketcap"
	"coinconv/internal/handlers"
	"context"
	"fmt"
)

func main() {
	isFiat, err := handlers.HandleRecognizeFiatCurrency(
		context.Background(),
		coinmarketcap.NewFiatCurrencyRecognizer(
			"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c", "sandbox",
		),
		"BTC",
	)

	fmt.Printf("%t %v\n", isFiat, err)
}
