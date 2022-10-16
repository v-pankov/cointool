package get

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/internal/fmterr"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap"
	"github.com/vdrpkv/cointool/internal/pkg/coinmarketcap/pkg/coinmarketcap/payload"
)

func Do(
	ctx context.Context,
	config coinmarketcap.Config,
	urlValues url.Values,
) (
	*http.Response,
	error,
) {
	client := &http.Client{}

	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		config.URL(),
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	config.SetHttpHeader(req.Header)
	req.URL.RawQuery = urlValues.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		var statusPayload payload.StatusPayload
		if err := json.NewDecoder(resp.Body).Decode(&statusPayload); err != nil {
			return nil, fmt.Errorf("parse response: %w", err)
		}
		return nil, fmterr.FormatError(resp.StatusCode, &statusPayload.Status)
	}

	return resp, nil
}
