package convert

import (
	"context"
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/mocks"
)

func Test_ConvertCliController(t *testing.T) {
	var (
		stubCtx  = context.Background()
		stubFrom = entity.CurrencySymbol("STUB_FROM")
		stubTo   = entity.CurrencySymbol("STUB_TO")
		stubErr  = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			amount entity.CurrencyAmount
			err    error
		}

		testCaseGive struct {
			args  []string
			stubs testCaseGiveStubs
		}

		testCaseWantHandlerArgs struct {
			amount entity.CurrencyAmount
			from   entity.CurrencySymbol
			to     entity.CurrencySymbol
		}

		testCaseWantHandler struct {
			args   testCaseWantHandlerArgs
			called bool
		}

		testCaseWant struct {
			handler testCaseWantHandler
			amount  interface{}
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
				args: []string{"1", "2", "3", "4"},
			},
			testCaseWant{
				err: cli.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"parse amount error",
			testCaseGive{
				args: []string{"_", stubFrom.String(), stubTo.String()},
			},
			testCaseWant{
				amount: 0,
				err:    strconv.ErrSyntax,
			},
		},
		{
			"error",
			testCaseGive{
				args: []string{"2.1", stubFrom.String(), stubTo.String()},
				stubs: testCaseGiveStubs{
					err: stubErr,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						amount: 2.1,
						from:   stubFrom,
						to:     stubTo,
					},
					called: true,
				},
				amount: nil,
				err:    stubErr,
			},
		},
		{
			"success",
			testCaseGive{
				args: []string{"2.2", stubFrom.String(), stubTo.String()},
				stubs: testCaseGiveStubs{
					amount: 4.4,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						amount: 2.2,
						from:   stubFrom,
						to:     stubTo,
					},
					called: true,
				},
				amount: entity.CurrencyAmount(4.4),
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			useCaseMock := mocks.NewUseCaseConvertCurrency(t)
			if testCase.want.handler.called {
				useCaseMock.
					On(
						"DoUseCaseConvertCurrency",
						stubCtx,
						testCase.want.handler.args.amount,
						testCase.want.handler.args.from,
						testCase.want.handler.args.to,
					).
					Return(
						testCase.give.stubs.amount, testCase.give.stubs.err,
					)
			}

			cliController := New(useCaseMock)
			gotAmount, gotErr := cliController.ExecCliController(
				stubCtx, testCase.give.args,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.amount, gotAmount)
		})
	}
}
