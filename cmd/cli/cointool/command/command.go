// Package command contains cointool commands along with utility functions shared by them.
package command

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vdrpkv/cointool/cmd/cli/cointool/variable"
	"github.com/vdrpkv/cointool/internal/handler/generic"
)

func RunGenericCommandHandler(
	cmd *cobra.Command,
	args []string,
	handler generic.GenericCommandHandler,
) {
	ctx, ctxCancel := context.WithTimeout(
		cmd.Context(), variable.Timeout,
	)
	defer ctxCancel()

	result, err := handler.HandleGenericCommand(ctx, args)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	fmt.Println(result)
}
