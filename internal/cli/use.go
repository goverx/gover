package cli

import (
	"fmt"
	"gover/internal/core"

	"github.com/spf13/cobra"
)

func useCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "use [version]",
		Short: "Switch to a specific Go version",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			version := args[0]
			if err := core.Manager.Use(version); err != nil {
				return err
			}
			fmt.Printf("Now using Go %s\n", version)
			return nil
		},
	}
}
