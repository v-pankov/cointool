package exchangerate

import (
	"context"
	"fmt"

	currencyExchangeRateClient "github.com/vdrpkv/cointool/internal/client/currency/exchangerate"
	fiatRecognizerClient "github.com/vdrpkv/cointool/internal/client/currency/fiat/recognizer"

	coinmarketcapCurrencyExchangeRate "github.com/vdrpkv/cointool/internal/coinmarketcap/pkg/client/currency/crypto/exchangerate"
	coinmarketcapFiatCurrencyRecognizer "github.com/vdrpkv/cointool/internal/coinmarketcap/pkg/client/currency/fiat/recognizer"

	"github.com/vdrpkv/cointool/internal/currency"
)

type client struct {
	currencyExchangeRateGetter currencyExchangeRateClient.CurrencyExchangeRateGetter
	fiatCurrencyRecognizer     fiatRecognizerClient.FiatCurrencyRecognizer
}

var _ currencyExchangeRateClient.CurrencyExchangeRateGetter = (*client)(nil)

func NewCurrencyExchangeRateGetter(
	apiKey, apiPrefix string,
) currencyExchangeRateClient.CurrencyExchangeRateGetter {
	return &client{
		currencyExchangeRateGetter: coinmarketcapCurrencyExchangeRate.NewCurrencyExchangeRateGetter(
			apiKey, apiPrefix,
		),
		fiatCurrencyRecognizer: coinmarketcapFiatCurrencyRecognizer.New(
			apiKey, apiPrefix,
		),
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

	rate, err := c.currencyExchangeRateGetter.GetExchangeRate(
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
