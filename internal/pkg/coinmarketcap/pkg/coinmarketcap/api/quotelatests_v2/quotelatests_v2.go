package quotelatests_v2

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/vdrpkv/cointool/internal/domain/entity/currency"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/internal/http/get"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap/payload"
)

type QuotesLatestV2 struct {
	payload.StatusPayload

	Data QuotesLatestV2Data `json:"data"`
}

type QuotesLatestV2Data map[string][]QuotesLatestV2DataItem

type QuotesLatestV2DataItem struct {
	Quote QuotesLatestV2DataQuote `json:"quote"`
}

type QuotesLatestV2DataQuote map[string]QuotesLatestV2DataQuoteItem

type QuotesLatestV2DataQuoteItem struct {
	Price float64 `json:"price"`
}

func Do(
	ctx context.Context,
	apiKey coinmarketcap.APIKey,
	env coinmarketcap.Environment,
	from, to currency.Symbol,
) (
	*QuotesLatestV2,
	error,
) {
	config := coinmarketcap.Config{
		APIKey:      apiKey,
		APIPath:     "v2/cryptocurrency/quotes/latest",
		Environment: env,
	}

	var quotesLatestV2 QuotesLatestV2
	if err := config.WrapCall(
		func() error {
			q := url.Values{}
			q.Add("symbol", from.String())
			q.Add("convert", to.String())

			resp, err := get.Do(ctx, config, q)
			if err != nil {
				return fmt.Errorf("http: %w", err)
			}

			if err := json.NewDecoder(resp.Body).Decode(&quotesLatestV2); err != nil {
				return fmt.Errorf("decode: %w", err)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}
	return &quotesLatestV2, nil
}
