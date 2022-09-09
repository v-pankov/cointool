package request

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

func sendGetRequest(
	ctx context.Context,
	apiKey, apiPrefix, apiSuffix string,
	urlValues url.Values,
) (
	*http.Response,
	error,
) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(
		ctx,
		"GET",
		fmt.Sprintf(
			"https://%s-api.coinmarketcap.com/%s",
			apiPrefix,
			apiSuffix,
		),
		nil,
	)

	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", apiKey)
	req.URL.RawQuery = urlValues.Encode()

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send request: %w", err)
	}

	return resp, nil
}
