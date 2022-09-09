package main

import (
	"coinconv/internal/coinmarketcap"
	"coinconv/internal/handlers"
	"context"
	"fmt"
)

func main() {
	rate, err := handlers.HandleGetExchangeRate(
		context.Background(),
		coinmarketcap.NewExchangeRateGetter(
			"b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c", "sandbox",
		),
		"BTC",
		"USD",
	)

	fmt.Printf("%f %v\n", rate, err)
}
