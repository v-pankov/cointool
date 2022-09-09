// Package recognizer implements fiat currency recognizer client based on coinmarketcap.
package recognizer

import (
	"context"
	"fmt"

	"github.com/vdrpkv/cointool/internal/currency"

	fiatCurrencyRecognizer "github.com/vdrpkv/cointool/internal/client/currency/fiat/recognizer"
	coinmarketcapFiatApi "github.com/vdrpkv/cointool/internal/coinmarketcap/internal/api/http/get/fiat"
)

type recognizer struct {
	apiKey, apiPrefix string
}

var _ fiatCurrencyRecognizer.FiatCurrencyRecognizer = (*recognizer)(nil)

func New(
	apiKey, apiPrefix string,
) fiatCurrencyRecognizer.FiatCurrencyRecognizer {
	return &recognizer{
		apiKey:    apiKey,
		apiPrefix: apiPrefix,
	}
}

func (r *recognizer) RecognizeFiatCurrency(
	ctx context.Context,
	from currency.Symbol,
) (
	bool,
	error,
) {
	isFiat, err := coinmarketcapFiatApi.Get(
		ctx,
		r.apiKey, r.apiPrefix,
		from,
	)

	if err != nil {
		return false, fmt.Errorf("coinmarketcap: %w", err)
	}

	return isFiat, nil
}
