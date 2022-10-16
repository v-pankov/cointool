package quotelatests_v2

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"

	"github.com/vdrpkv/cointool/internal/domain/entity"
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

func (qs QuotesLatestV2) FindExchangeRate(from, to entity.CurrencySymbol) (entity.ExchangeRate, error) {
	// Get items.
	items, ok := qs.Data[from.String()]
	if !ok {
		return 0, ErrQuotesItemsNotFound
	}

	// Exit with error if no quotes are received.
	if len(items) == 0 {
		return 0, ErrQuotesItemsEmpty
	}

	// Take first.
	quote, ok := items[0].Quote[to.String()]
	if !ok {
		return 0, ErrQuoteNotFound
	}

	return entity.ExchangeRate(quote.Price), nil
}

var (
	ErrQuotesItemsNotFound = errors.New("quotes items not found")
	ErrQuotesItemsEmpty    = errors.New("quotes items empty")
	ErrQuoteNotFound       = errors.New("quote not found")
)

func Do(
	ctx context.Context,
	apiKey coinmarketcap.APIKey,
	env coinmarketcap.Environment,
	from, to entity.CurrencySymbol,
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
