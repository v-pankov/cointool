package exchangerate

import (
	"context"
	"fmt"

	exchangeRateClient "github.com/vdrpkv/cointool/internal/client/currency/exchangerate"

	coinmarketcapExchangeRate "github.com/vdrpkv/cointool/internal/coinmarketcap/internal/exchangerate"
	coinmarketcapFiatRecognizer "github.com/vdrpkv/cointool/internal/coinmarketcap/internal/fiat"

	"github.com/vdrpkv/cointool/internal/currency"
)

type client struct {
	apiKey, apiPrefix string
}

var _ exchangeRateClient.CurrencyExchangeRateGetter = (*client)(nil)

func NewCurrencyExchangeRateGetter(
	apiKey, apiPrefix string,
) exchangeRateClient.CurrencyExchangeRateGetter {
	return &client{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

func (c *client) GetExchangeRate(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	// Check is FROM currency fiat one.
	isFiat, err := coinmarketcapFiatRecognizer.Get(
		ctx,
		c.apiKey, c.apiPrefix,
		from,
	)

	if err != nil {
		return 0, fmt.Errorf("recognize fiat currency: %w", err)
	}

	// Flip symbols because coinmarketcapExchangeRate.Get
	// does not find exchange rate from fiat currency.
	if isFiat {
		from, to = to, from
	}

	rate, err := coinmarketcapExchangeRate.Get(
		ctx,
		c.apiKey, c.apiPrefix,
		from, to,
	)
	if err != nil {
		return 0, fmt.Errorf("get exchange rate: %w", err)
	}

	// Flip exchange rate if first currency is fiat one.
	if isFiat {
		rate = rate.Flip()
	}

	return rate, nil
}
