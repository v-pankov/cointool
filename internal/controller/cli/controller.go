package cli

import (
	"context"
	"errors"
)

type CliController interface {
	ExecCliController(ctx context.Context, args []string) (interface{}, error)
}

var (
	ErrUnexpectedNumberOfArguments = errors.New("unexpected number of arguments")
)
