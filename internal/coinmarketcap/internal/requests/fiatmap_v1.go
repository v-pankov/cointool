package requests

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vdrpkv/cointool/internal/coinmarketcap/internal/response"
)

func RequestFiatMapV1(
	ctx context.Context,
	apiKey, apiPrefix string,
) (
	*response.FiatMapV1,
	error,
) {
	q := url.Values{}

	resp, err := sendGetRequest(
		ctx, apiKey, apiPrefix, "v1/fiat/map", q,
	)

	if err != nil {
		return nil, err
	}

	var fiatMap response.FiatMapV1

	if err := json.NewDecoder(resp.Body).Decode(&fiatMap); err != nil {
		return nil, fmt.Errorf("parse response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmtBadStatusError(resp.StatusCode, &fiatMap.Status)
	}

	return &fiatMap, nil
}
