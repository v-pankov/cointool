package coinmarketcap

import (
	"coinconv/internal/coinmarketcap/internal/requests"
	"coinconv/internal/currency"
	"coinconv/internal/handlers"
	"context"
)

type fiatCurrencyRecognizer struct {
	apiKey    string
	apiPrefix string
}

func NewFiatCurrencyRecognizer(
	apiKey string,
	apiPrefix string,
) handlers.FiatCurrencyRecognizer {
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
