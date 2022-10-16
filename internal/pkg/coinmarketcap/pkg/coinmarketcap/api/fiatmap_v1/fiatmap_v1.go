package fiatmap_v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/internal/http/get"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap/payload"
)

type FiatMapV1 struct {
	payload.StatusPayload

	Data FiatMapV1Data `json:"data"`
}

type FiatMapV1Data []FiatMapV1DataItem

type FiatMapV1DataItem struct {
	Symbol string `json:"symbol"`
}

func (m FiatMapV1) IsFiatCurrency(symbol entity.CurrencySymbol) bool {
	for _, item := range m.Data {
		if symbol.String() == item.Symbol {
			return true
		}
	}
	return false
}

func Do(
	ctx context.Context,
	apiKey coinmarketcap.APIKey,
	env coinmarketcap.Environment,
) (
	*FiatMapV1,
	error,
) {
	config := coinmarketcap.Config{
		APIKey:      apiKey,
		APIPath:     "v1/fiat/map",
		Environment: env,
	}

	var fiatMap FiatMapV1
	if err := config.WrapCall(
		func() error {
			resp, err := get.Do(ctx, config, url.Values{})
			if err != nil {
				return fmt.Errorf("http: %w", err)
			}

			if err := json.NewDecoder(resp.Body).Decode(&fiatMap); err != nil {
				return fmt.Errorf("decode: %w", err)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}
	return &fiatMap, nil
}
