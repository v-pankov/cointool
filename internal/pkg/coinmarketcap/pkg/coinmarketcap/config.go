package coinmarketcap

import (
	"errors"
	"fmt"
	"net/http"
)

type (
	APIKey      string
	APIPath     string
	Environment string

	Config struct {
		APIKey      APIKey
		APIPath     APIPath
		Environment Environment
	}
)

func (c *Config) Validate() error {
	if c.APIKey == "" {
		return ErrEmptyAPIKey
	}

	if c.APIPath == "" {
		return ErrEmptyAPIPath
	}

	if c.Environment == "" {
		return ErrEmptyEnvironment
	}

	return nil
}

var (
	ErrEmptyAPIKey      = errors.New("api key is empty")
	ErrEmptyAPIPath     = errors.New("api path is empty")
	ErrEmptyEnvironment = errors.New("environment is empty")
)

func (c *Config) URL() string {
	return fmt.Sprintf(
		"https://%s-api.coinmarketcap.com/%s",
		c.Environment,
		c.APIPath,
	)
}

func (c *Config) SetHttpHeader(httpHeader http.Header) {
	httpHeader.Set("Accepts", "application/json")
	httpHeader.Add("X-CMC_PRO_API_KEY", string(c.APIKey))
}

func (c *Config) WrapCall(fn func() error) error {
	if err := c.Validate(); err != nil {
		return fmt.Errorf("invalid api: %w", err)
	}

	if err := fn(); err != nil {
		return fmt.Errorf("api [%s]: %w", c.APIPath, err)
	}

	return nil
}
