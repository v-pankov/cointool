// Package variable stores cointool flags global variables.
package variable

import "time"

var (
	ApiKey                string
	ApiPrefix             string
	Timeout               time.Duration
	ExchangeRateZeroValue float64
)
