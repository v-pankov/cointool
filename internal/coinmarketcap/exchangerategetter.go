package coinmarketcap

import (
	"context"
	"errors"

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
) handler.CryptoCurrencyExchangeRateGetter {
	return &exchangeRateGetter{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

var (
	ErrCurrencySymbolNotFound = errors.New("currency symbol not found")
	ErrExchangeRateNotFound   = errors.New("exchange rate not found")
)

func (g *exchangeRateGetter) GetCryptoCurrencyExchangeRate(
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
