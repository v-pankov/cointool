package exchangerate

import (
	"context"
	"errors"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/request"
	"github.com/vdrpkv/cointool/internal/currency"
)

var (
	ErrCurrencySymbolNotFound = errors.New("currency symbol not found")
	ErrExchangeRateNotFound   = errors.New("exchange rate not found")
)

// Get returns exchange rate for given cryptocurrency pair.
// Get does not return exchange rate from fiat currency because quotes/latest doesn't provide it.
func Get(
	ctx context.Context,
	apiKey, apiPrefix string,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	quotesLatest, err := request.RequestQuotesLatestV2(
		ctx,
		apiKey, apiPrefix,
		from, to,
	)
	if err != nil {
		return 0, err
	}

	items, ok := quotesLatest.Data[from.String()]
	if !ok {
		return 0, ErrCurrencySymbolNotFound
	}

	if len(items) == 0 {
		return 0, ErrExchangeRateNotFound
	}

	// take first
	quote, ok := items[0].Quote[to.String()]
	if !ok {
		return 0, ErrExchangeRateNotFound
	}

	return currency.ExchangeRate(quote.Price), nil
}
