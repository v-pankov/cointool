// Package getter implements currency exchange getter client based on coinmarketcap.
package getter

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/currency"

	currencyExchangeRateGetter "github.com/vdrpkv/cointool/internal/client/currency/exchangerate/getter"
	fiatCurrencyRecognizer "github.com/vdrpkv/cointool/internal/client/currency/fiat/recognizer"
	coinmarketcapCryptoCurrencyExchangeRateGetter "github.com/vdrpkv/cointool/internal/coinmarketcap/pkg/client/currency/crypto/exchangerate/getter"
	coinmarketcapFiatCurrencyRecognizer "github.com/vdrpkv/cointool/internal/coinmarketcap/pkg/client/currency/fiat/recognizer"
)

type getter struct {
	currencyExchangeRateGetter currencyExchangeRateGetter.CurrencyExchangeRateGetter
	fiatCurrencyRecognizer     fiatCurrencyRecognizer.FiatCurrencyRecognizer
}

var _ currencyExchangeRateGetter.CurrencyExchangeRateGetter = (*getter)(nil)

func New(
	apiKey, apiPrefix string,
) currencyExchangeRateGetter.CurrencyExchangeRateGetter {
	return &getter{
		currencyExchangeRateGetter: coinmarketcapCryptoCurrencyExchangeRateGetter.New(
			apiKey, apiPrefix,
		),
		fiatCurrencyRecognizer: coinmarketcapFiatCurrencyRecognizer.New(
			apiKey, apiPrefix,
		),
	}
}

func (c *getter) GetCurrencyExchangeRate(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	// Check is FROM currency fiat one.
	isFiat, err := c.fiatCurrencyRecognizer.RecognizeFiatCurrency(
		ctx, from,
	)

	if err != nil {
		return 0, fmt.Errorf("recognize fiat currency: %w", err)
	}

	// Flip symbols because coinmarketcapExchangeRate.Get
	// does not find exchange rate from fiat currency.
	if isFiat {
		from, to = to, from
	}

	rate, err := c.currencyExchangeRateGetter.GetCurrencyExchangeRate(
		ctx, from, to,
	)
	if err != nil {
		return 0, fmt.Errorf("get currency exchange rate: %w", err)
	}

	// Flip exchange rate if first currency is fiat one.
	if isFiat {
		rate = rate.Flip()
	}

	return rate, nil
}
