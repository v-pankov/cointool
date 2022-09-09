package request

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
	const apiPath = "v1/fiat/map"

	fiatMap, err := requestFiatMapV1(ctx, apiKey, apiPrefix, apiPath)

	if err != nil {
		return nil, fmt.Errorf("api [%s]: %w", apiPath, err)
	}

	return fiatMap, nil
}

func requestFiatMapV1(
	ctx context.Context,
	apiKey, apiPrefix, apiPath string,
) (
	*response.FiatMapV1,
	error,
) {
	q := url.Values{}

	resp, err := sendGetRequest(
		ctx, apiKey, apiPrefix, apiPath, q,
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
