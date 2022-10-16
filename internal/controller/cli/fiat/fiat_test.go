package fiat

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/vdrpkv/cointool/internal/controller/cli"
	"github.com/vdrpkv/cointool/internal/domain/entity"
	"github.com/vdrpkv/cointool/internal/mocks"
)

func Test_FiatCliController(t *testing.T) {
	var (
		stubCtx    = context.Background()
		stubSymbol = entity.CurrencySymbol("STUB_SYMBOL")
		stubErr    = errors.New("stub error")
	)

	type (
		testCaseGiveStubs struct {
			fiat bool
			err  error
		}

		testCaseGive struct {
			args  []string
			stubs testCaseGiveStubs
		}

		testCaseWantHandlerArgs struct {
			symbol entity.CurrencySymbol
		}

		testCaseWantHandler struct {
			args   testCaseWantHandlerArgs
			called bool
		}

		testCaseWant struct {
			handler testCaseWantHandler
			fiat    interface{}
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
			"more than enough args",
			testCaseGive{
				args: []string{"1", "2"},
			},
			testCaseWant{
				err: cli.ErrUnexpectedNumberOfArguments,
			},
		},
		{
			"error",
			testCaseGive{
				args: []string{stubSymbol.String()},
				stubs: testCaseGiveStubs{
					err: stubErr,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						symbol: stubSymbol,
					},
					called: true,
				},
				fiat: nil,
				err:  stubErr,
			},
		},
		{
			"success",
			testCaseGive{
				args: []string{stubSymbol.String()},
				stubs: testCaseGiveStubs{
					fiat: false,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						symbol: stubSymbol,
					},
					called: true,
				},
				fiat: false,
			},
		},
		{
			"success",
			testCaseGive{
				args: []string{stubSymbol.String()},
				stubs: testCaseGiveStubs{
					fiat: true,
				},
			},
			testCaseWant{
				handler: testCaseWantHandler{
					args: testCaseWantHandlerArgs{
						symbol: stubSymbol,
					},
					called: true,
				},
				fiat: true,
			},
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			useCaseMock := mocks.NewUseCaseRecognizeFiatCurrency(t)
			if testCase.want.handler.called {
				useCaseMock.
					On(
						"DoUseCaseRecognizeFiatCurrency",
						stubCtx,
						testCase.want.handler.args.symbol,
					).
					Return(
						testCase.give.stubs.fiat, testCase.give.stubs.err,
					)
			}

			cliController := New(useCaseMock)
			gotRate, gotErr := cliController.ExecCliController(
				stubCtx, testCase.give.args,
			)

			require.ErrorIs(t, gotErr, testCase.want.err)
			require.Equal(t, testCase.want.fiat, gotRate)
		})
	}
}
