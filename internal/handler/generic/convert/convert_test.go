package convert

import (
	"context"
	"strconv"
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
	)

	type (
		testCaseGiveStubs struct {
			amount currency.Amount
			err    error
		}

		testCaseGive struct {
			args  []string
			stubs testCaseGiveStubs
		}

		testCaseWantHandlerArgs struct {
			amount currency.Amount
			from   currency.Symbol
			to     currency.Symbol
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
				args: []string{"1", "2", "3", "4"},
			},
			testCaseWant{
				err: generic.ErrUnexpectedNumberOfArguments,
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
				amount: currency.Amount(4.4),
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			convertCommandHandlerMock := mocks.NewConvertCommandHandler(t)
			if testCase.want.handler.called {
				convertCommandHandlerMock.
					On(
						"HandleConvertCommand",
						stubCtx,
						testCase.want.handler.args.amount,
						testCase.want.handler.args.from,
						testCase.want.handler.args.to,
					).
					Return(
						testCase.give.stubs.amount, testCase.give.stubs.err,
					)
			}

			genericConvertCommandHandler := New(convertCommandHandlerMock)
			gotAmount, gotErr := genericConvertCommandHandler.HandleGenericCommand(
				stubCtx, testCase.give.args,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.amount, gotAmount)
		})
	}
}
