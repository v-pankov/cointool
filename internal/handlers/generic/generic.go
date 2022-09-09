package generic

import (
	"context"
	"errors"
)

var (
	ErrNotEnoughArgs = errors.New("not enough args")
)

type GenericCommandHandler interface {
	HandleGenericCommand(
		ctx context.Context,
		args []string,
	) (
		interface{},
		error,
	)
}
