package coinmarketcap

import (
	"context"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/requests"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler"
)

type fiatCurrencyRecognizer struct {
	apiKey    string
	apiPrefix string
}

func NewFiatCurrencyRecognizer(
	apiKey string,
	apiPrefix string,
) handler.FiatCurrencyRecognizer {
	return &fiatCurrencyRecognizer{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

func (r *fiatCurrencyRecognizer) RecognizeFiatCurrency(
	ctx context.Context,
	symbol currency.Symbol,
) (
	bool,
	error,
) {
	fiatMap, err := requests.RequestFiatMapV1(ctx, r.apiKey, r.apiPrefix)
	if err != nil {
		return false, err
	}

	for _, item := range fiatMap.Data {
		if symbol.String() == item.Symbol {
			return true, nil
		}
	}

	return false, nil
}
