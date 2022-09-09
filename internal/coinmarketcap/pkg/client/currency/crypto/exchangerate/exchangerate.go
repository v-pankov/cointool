package exchangerate

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/currency"

	currencyExchangeRateClient "github.com/vdrpkv/cointool/internal/client/currency/exchangerate"
	coinmarketcapExchangeRateApi "github.com/vdrpkv/cointool/internal/coinmarketcap/internal/exchangerate"
)

type exchangerate struct {
	apiKey, apiPrefix string
}

var _ currencyExchangeRateClient.CurrencyExchangeRateGetter = (*exchangerate)(nil)

func New(
	apiKey, apiPrefix string,
) currencyExchangeRateClient.CurrencyExchangeRateGetter {
	return &exchangerate{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

func (r *exchangerate) GetExchangeRate(
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
