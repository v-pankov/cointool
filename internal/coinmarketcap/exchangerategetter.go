package coinmarketcap

import (
	"context"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/requests"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handlers"
)

type exchangeRateGetter struct {
	apiKey    string
	apiPrefix string
}

func NewExchangeRateGetter(
	apiKey, apiPrefix string,
) handlers.ExchangeRateGetter {
	return &exchangeRateGetter{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

func (g *exchangeRateGetter) GetExchangeRate(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	quotesLatest, err := requests.RequestQuotesLatestV2(
		ctx,
		g.apiKey, g.apiPrefix,
		from, to,
	)
	if err != nil {
		return 0, err
	}

	items, ok := quotesLatest.Data[from.String()]
	if !ok {
		return 0, handlers.ErrCurrencySymbolNotFound
	}

	if len(items) == 0 {
		return 0, handlers.ErrExchangeRateNotFound
	}

	// take first
	quote, ok := items[0].Quote[to.String()]
	if !ok {
		return 0, handlers.ErrExchangeRateNotFound
	}

	return currency.ExchangeRate(quote.Price), nil
}
