package coinmarketcap

import (
	"context"
	"errors"
	"fmt"

	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/domain/usecase/exchangerate"
	"github.com/vdrpkv/cointool/internal/domain/usecase/fiat"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap/api/fiatmap_v1"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap/api/quotelatests_v2"
)

type (
	APIKey      string
	Environment string

	Client struct {
		APIKey      APIKey
		Environment Environment
	}
)

func NewClient(
	apiKey APIKey,
	environment Environment,
) *Client {
	return &Client{
		APIKey:      apiKey,
		Environment: environment,
	}
}

var _ fiat.FiatCurrencyClient = (*Client)(nil)

func (c *Client) RecognizeFiatCurrency(ctx context.Context, symbol entity.CurrencySymbol) (bool, error) {
	fiatCurrencyClient := fiatCurrencyRecognizer{func() (*fiatmap_v1.FiatMapV1, error) {
		return fiatmap_v1.Do(
			ctx,
			coinmarketcap.APIKey(c.APIKey),
			coinmarketcap.Environment(c.Environment),
		)
	}}
	return fiatCurrencyClient.RecognizeFiatCurrency(ctx, symbol)
}

var _ exchangerate.ExchangeRateClient = (*Client)(nil)

func (c *Client) GetExchangeRate(ctx context.Context, from, to entity.CurrencySymbol) (entity.ExchangeRate, error) {
	var (
		fiatCurrencyClient = fiatCurrencyRecognizer{func() (*fiatmap_v1.FiatMapV1, error) {
			return fiatmap_v1.Do(
				ctx,
				coinmarketcap.APIKey(c.APIKey),
				coinmarketcap.Environment(c.Environment),
			)
		}}
		exchangeRateClientPlain = exchangeRateGetterPlain{func(from, to entity.CurrencySymbol) (*quotelatests_v2.QuotesLatestV2, error) {
			return quotelatests_v2.Do(
				ctx,
				coinmarketcap.APIKey(c.APIKey),
				coinmarketcap.Environment(c.Environment),
				from, to,
			)
		}}
		exchangeRateClientFiat = exchangeRateGetterFiat{
			fiatCurrencyClient: fiatCurrencyClient,
			exchangeRateClient: exchangeRateClientPlain,
		}
	)
	return exchangeRateClientFiat.GetExchangeRate(ctx, from, to)
}

type fiatCurrencyRecognizer struct {
	backend func() (*fiatmap_v1.FiatMapV1, error)
}

var _ fiat.FiatCurrencyClient = fiatCurrencyRecognizer{}

func (r fiatCurrencyRecognizer) RecognizeFiatCurrency(ctx context.Context, symbol entity.CurrencySymbol) (bool, error) {
	fiatmap_v1, err := r.backend()
	if err != nil {
		return false, fmt.Errorf("coinmarketcap: %w", err)
	}
	return fiatmap_v1.IsFiatCurrency(symbol), nil
}

type exchangeRateGetterPlain struct {
	backend func(from, to entity.CurrencySymbol) (*quotelatests_v2.QuotesLatestV2, error)
}

var _ exchangerate.ExchangeRateClient = exchangeRateGetterPlain{}

func (g exchangeRateGetterPlain) GetExchangeRate(ctx context.Context, from, to entity.CurrencySymbol) (entity.ExchangeRate, error) {
	// Get latest quotes from CoinMarktetCap.
	quotesLatest_v2, err := g.backend(from, to)
	if err != nil {
		return 0, fmt.Errorf("coinmarketcap: %w", err)
	}
	// Flip rate if first currency symbol is fiat one.
	rate, err := quotesLatest_v2.FindExchangeRate(from, to)
	if err != nil {
		switch {
		case errors.Is(err, quotelatests_v2.ErrQuoteNotFound):
			return 0, ErrExchangeRateNotFound
		default:
			return 0, fmt.Errorf("find exchange rate: %w", err)
		}
	}
	return rate, nil
}

var (
	ErrExchangeRateNotFound = errors.New("exchange rate is not found")
)

type exchangeRateGetterFiat struct {
	fiatCurrencyClient fiat.FiatCurrencyClient
	exchangeRateClient exchangerate.ExchangeRateClient
}

var _ exchangerate.ExchangeRateClient = exchangeRateGetterFiat{}

func (g exchangeRateGetterFiat) GetExchangeRate(ctx context.Context, from, to entity.CurrencySymbol) (entity.ExchangeRate, error) {
	// Find out is first symbol denotes fiat currency.
	isFiat, err := g.fiatCurrencyClient.RecognizeFiatCurrency(ctx, from)
	if err != nil {
		return 0, err
	}

	// Flip symbols if first one denotes fiat currency.
	if isFiat {
		from, to = to, from
	}

	rate, err := g.exchangeRateClient.GetExchangeRate(ctx, from, to)
	if err != nil {
		return 0, err
	}

	if isFiat {
		rate = rate.Flip()
	}

	return rate, nil
}
