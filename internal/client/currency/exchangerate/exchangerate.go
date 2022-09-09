package exchangerate

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
)

type (
	// ExchangeRateGetter gets exchange rate for given currency pair.
	ExchangeRateGetter interface {
		GetExchangeRate(
			ctx context.Context,
			from, to currency.Symbol,
		) (
			currency.ExchangeRate,
			error,
		)
	}
)
