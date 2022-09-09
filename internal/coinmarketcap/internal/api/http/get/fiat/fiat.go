package fiat

import (
	"context"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/request"
	"github.com/vdrpkv/cointool/internal/currency"
)

func Get(
	ctx context.Context,
	apiKey, apiPrefix string,
	symbol currency.Symbol,
) (
	bool,
	error,
) {
	fiatMap, err := request.RequestFiatMapV1(ctx, apiKey, apiPrefix)
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
