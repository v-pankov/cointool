// Package command contains cointool commands along with utility functions shared by them.
package command

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/controller/cli"
)

func ExecCliController(
	cmd *cobra.Command,
	args []string,
	controller cli.CliController,
) {
	ctx, ctxCancel := context.WithTimeout(
		cmd.Context(), variable.Timeout,
	)
	defer ctxCancel()

	result, err := controller.ExecCliController(ctx, args)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Println(result)
}
