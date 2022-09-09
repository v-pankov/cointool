package request

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/response"
	"github.com/vdrpkv/cointool/internal/currency"
)

func RequestQuotesLatestV2(
	ctx context.Context,
	apiKey, apiPrefix string,
	from, to currency.Symbol,
) (
	*response.QuotesLatestV2,
	error,
) {
	const apiPath = "v2/cryptocurrency/quotes/latest"

	quotesLatest, err := requestQuotesLatestV2(
		ctx,
		apiKey, apiPrefix, apiPath,
		from, to,
	)

	if err != nil {
		return nil, fmt.Errorf("api [%s]: %w", apiPath, err)
	}

	return quotesLatest, nil
}

func requestQuotesLatestV2(
	ctx context.Context,
	apiKey, apiPrefix, apiPath string,
	from, to currency.Symbol,
) (
	*response.QuotesLatestV2,
	error,
) {
	q := url.Values{}
	q.Add("symbol", from.String())
	q.Add("convert", to.String())

	resp, err := sendGetRequest(
		ctx, apiKey, apiPrefix, apiPath, q,
	)

	if err != nil {
		return nil, err
	}

	var quotesLatest response.QuotesLatestV2

	if err := json.NewDecoder(resp.Body).Decode(&quotesLatest); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmtBadStatusError(resp.StatusCode, &quotesLatest.Status)
	}

	return &quotesLatest, nil
}
