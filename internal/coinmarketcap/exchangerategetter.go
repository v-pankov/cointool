package coinmarketcap

import (
	"context"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/request"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
)

type exchangeRateGetter struct {
	apiKey    string
	apiPrefix string
}

func NewExchangeRateGetter(
	apiKey, apiPrefix string,
) handler.ExchangeRateGetter {
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
	quotesLatest, err := request.RequestQuotesLatestV2(
		ctx,
		g.apiKey, g.apiPrefix,
		from, to,
	)
	if err != nil {
		return 0, err
	}

	items, ok := quotesLatest.Data[from.String()]
	if !ok {
		return 0, handler.ErrCurrencySymbolNotFound
	}

	if len(items) == 0 {
		return 0, handler.ErrExchangeRateNotFound
	}

	// take first
	quote, ok := items[0].Quote[to.String()]
	if !ok {
		return 0, handler.ErrExchangeRateNotFound
	}

	return currency.ExchangeRate(quote.Price), nil
}
