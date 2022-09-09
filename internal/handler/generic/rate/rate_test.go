package rate

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/currency"
	"github.com/vdrpkv/cointool/internal/handler/generic"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_HandleGenericCommand(t *testing.T) {
	var (
		stubCtx  = context.Background()
		stubFrom = currency.Symbol("STUB_FROM")
		stubTo   = currency.Symbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			rate currency.ExchangeRate
			err  error
		}

		testCaseGive struct {
			args  []string
			stubs testCaseGiveStubs
		}

		testCaseWantHandlerArgs struct {
			from currency.Symbol
			to   currency.Symbol
		}

		testCaseWantHandler struct {
			args   testCaseWantHandlerArgs
			called bool
		}

		testCaseWant struct {
			handler testCaseWantHandler
			rate    interface{}
			err     error
		}

		testCase struct {
			name string
			give testCaseGive
			want testCaseWant
		}
	)

	for _, testCase := range []testCase{
		{
			"no args",
			testCaseGive{},
			testCaseWant{
				err: generic.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"not enough args",
			testCaseGive{
				args: []string{"1"},
			},
			testCaseWant{
				err: generic.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"more than enough args",
			testCaseGive{
				args: []string{"1", "2", "3"},
			},
			testCaseWant{
				err: generic.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"error",
			testCaseGive{
				args: []string{stubFrom.String(), stubTo.String()},
				stubs: testCaseGiveStubs{
					err: stubErr,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						from: stubFrom,
						to:   stubTo,
					},
					called: true,
				},
				rate: 0,
				err:  stubErr,
			},
		},
		{
			"success",
			testCaseGive{
				args: []string{stubFrom.String(), stubTo.String()},
				stubs: testCaseGiveStubs{
					rate: 2.2,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						from: stubFrom,
						to:   stubTo,
					},
					called: true,
				},
				rate: currency.ExchangeRate(2.2),
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			rateCommandHandlerMock := mocks.NewRateCommandHandler(t)
			if testCase.want.handler.called {
				rateCommandHandlerMock.
					On(
						"HandleRateCommand",
						stubCtx,
						testCase.want.handler.args.from,
						testCase.want.handler.args.to,
					).
					Return(
						testCase.give.stubs.rate, testCase.give.stubs.err,
					)
			}

			genericRateCommandHandler := New(rateCommandHandlerMock)
			gotRate, gotErr := genericRateCommandHandler.HandleGenericCommand(
				stubCtx, testCase.give.args,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.rate, gotRate)
		})
	}
}
