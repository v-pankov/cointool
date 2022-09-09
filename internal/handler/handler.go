package handler

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
)

type (
	// CryptoCurrencyExchangeRateGetter gets exchange rate for given cryptocurrency pair.
	// CryptoCurrencyExchangeRateGetter accepts cryptocurrency symbol as FROM and any symbol as TO.
	// CryptoCurrencyExchangeRateGetter may not find exchange rate if fiat currency symbol is passed as FROM.
	CryptoCurrencyExchangeRateGetter interface {
		GetCryptoCurrencyExchangeRate(
			ctx context.Context,
			from, to currency.Symbol,
		) (
			currency.ExchangeRate,
			error,
		)
	}

	// FiatCurrencyRecognizer checks is given currency symbol is fiat one.
	FiatCurrencyRecognizer interface {
		RecognizeFiatCurrency(
			ctx context.Context,
			symbol currency.Symbol,
		) (
			bool,
			error,
		)
	}
)
