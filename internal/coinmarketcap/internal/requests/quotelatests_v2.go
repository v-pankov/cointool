package requests

import (
	"coinconv/internal/coinmarketcap/internal/responses"
	"coinconv/internal/currency"

	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func RequestQuotesLatestV2(
	ctx context.Context,
	apiKey, apiPrefix string,
	from, to currency.Symbol,
) (
	*responses.QuotesLatestV2,
	error,
) {
	q := url.Values{}
	q.Add("symbol", from.String())
	q.Add("convert", to.String())

	resp, err := sendGetRequest(
		ctx, apiKey, apiPrefix, "v2/cryptocurrency/quotes/latest", q,
	)

	if err != nil {
		return nil, err
	}

	var quotesLatest responses.QuotesLatestV2

	if err := json.NewDecoder(resp.Body).Decode(&quotesLatest); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmtBadStatusError(resp.StatusCode, &quotesLatest.Status)
	}

	return &quotesLatest, nil
}
