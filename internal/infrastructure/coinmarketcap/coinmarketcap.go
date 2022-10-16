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
	fiatmap_v1, err := fiatmap_v1.Do(
		ctx,
		coinmarketcap.APIKey(c.APIKey),
		coinmarketcap.Environment(c.Environment),
	)
	if err != nil {
		return false, fmt.Errorf("coinmarketcap: %w", err)
	}

	return fiatmap_v1.IsFiatCurrency(symbol), nil
}

var _ exchangerate.ExchangeRateClient = (*Client)(nil)

func (c *Client) GetExchangeRate(ctx context.Context, from, to entity.CurrencySymbol) (entity.ExchangeRate, error) {
	// Find out is first symbol denotes fiat currency.
	isFiat, err := c.RecognizeFiatCurrency(ctx, from)
	if err != nil {
		return 0, err
	}

	// Flip symbols if first one denotes fiat currency.
	if isFiat {
		from, to = to, from
	}

	// Get latest quotes from CoinMarktetCap.
	quotesLatest_v2, err := quotelatests_v2.Do(
		ctx,
		coinmarketcap.APIKey(c.APIKey),
		coinmarketcap.Environment(c.Environment),
		from, to,
	)
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

	if isFiat {
		rate = rate.Flip()
	}

	return rate, nil
}

var (
	ErrExchangeRateNotFound = errors.New("exchange rate is not found")
)
