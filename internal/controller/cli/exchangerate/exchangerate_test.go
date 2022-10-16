package exchangerate

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/mocks"
)

func Test_ExchangeRateCliController(t *testing.T) {
	var (
		stubCtx  = context.Background()
		stubFrom = entity.CurrencySymbol("STUB_FROM")
		stubTo   = entity.CurrencySymbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			rate entity.ExchangeRate
			err  error
		}

		testCaseGive struct {
			args  []string
			stubs testCaseGiveStubs
		}

		testCaseWantHandlerArgs struct {
			from entity.CurrencySymbol
			to   entity.CurrencySymbol
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
				err: cli.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"not enough args",
			testCaseGive{
				args: []string{"1"},
			},
			testCaseWant{
				err: cli.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"more than enough args",
			testCaseGive{
				args: []string{"1", "2", "3"},
			},
			testCaseWant{
				err: cli.ErrUnexpectedNumberOfArguments,
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
				rate: nil,
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
				rate: entity.ExchangeRate(2.2),
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			useCaseMock := mocks.NewUseCaseGetExchangeRate(t)
			if testCase.want.handler.called {
				useCaseMock.
					On(
						"DoUseCaseGetExchangeRate",
						stubCtx,
						testCase.want.handler.args.from,
						testCase.want.handler.args.to,
					).
					Return(
						testCase.give.stubs.rate, testCase.give.stubs.err,
					)
			}

			cliController := New(useCaseMock)
			gotRate, gotErr := cliController.ExecCliController(
				stubCtx, testCase.give.args,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.rate, gotRate)
		})
	}
}
