package exchangerate

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
)

type (
	// CurrencyExchangeRateGetter gets exchange rate for given currency pair.
	CurrencyExchangeRateGetter interface {
		GetExchangeRate(
			ctx context.Context,
			from, to currency.Symbol,
		) (
			currency.ExchangeRate,
			error,
		)
	}
)
