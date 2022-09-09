// Package convert provides 'cointool convert' command handler.
package convert

import (
	"context"

	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler/rate"
)

type ConvertCommandHandler interface {
	HandleConvertCommand(
		ctx context.Context,
		amount currency.Amount,
		from, to currency.Symbol,
	) (
		currency.Amount,
		error,
	)
}

type convertHandler struct {
	rateHandler rate.RateCommandHandler
}

var _ ConvertCommandHandler = (*convertHandler)(nil)

func New(
	rateHandler rate.RateCommandHandler,
) ConvertCommandHandler {
	return &convertHandler{
		rateHandler: rateHandler,
	}
}

func (h *convertHandler) HandleConvertCommand(
	ctx context.Context,
	amount currency.Amount,
	from, to currency.Symbol,
) (
	currency.Amount,
	error,
) {
	rate, err := h.rateHandler.HandleRateCommand(
		ctx,
		from, to,
	)

	if err != nil {
		return 0, err
	}

	return rate.Convert(amount), nil
}
