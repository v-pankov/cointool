package generic

import (
	"context"
	"errors"
)

var (
	ErrUnexpectedNumberOfArguments = errors.New("unexpected number of arguments")
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
