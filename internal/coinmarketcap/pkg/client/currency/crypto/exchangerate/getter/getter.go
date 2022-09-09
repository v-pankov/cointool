package getter

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/currency"

	currencyExchangeRateGetter "github.com/vdrpkv/cointool/internal/client/currency/exchangerate/getter"
	coinmarketcapExchangeRateApi "github.com/vdrpkv/cointool/internal/coinmarketcap/internal/api/http/get/exchangerate"
)

type getter struct {
	apiKey, apiPrefix string
}

var _ currencyExchangeRateGetter.CurrencyExchangeRateGetter = (*getter)(nil)

func New(
	apiKey, apiPrefix string,
) currencyExchangeRateGetter.CurrencyExchangeRateGetter {
	return &getter{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

func (r *getter) GetCurrencyExchangeRate(
	ctx context.Context,
	from, to currency.Symbol,
) (
	currency.ExchangeRate,
	error,
) {
	rate, err := coinmarketcapExchangeRateApi.Get(
		ctx,
		r.apiKey, r.apiPrefix,
		from, to,
	)

	if err != nil {
		return 0, fmt.Errorf("coinmarketcap: %w", err)
	}

	return rate, nil
}
