// Package getter provides currency exchange getter client interface.
package getter

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
)

type (
	// CurrencyExchangeRateGetter gets exchange rate for given currency pair.
	CurrencyExchangeRateGetter interface {
		GetCurrencyExchangeRate(
			ctx context.Context,
			from, to currency.Symbol,
		) (
			currency.ExchangeRate,
			error,
		)
	}
)
