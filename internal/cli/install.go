package cli

import (
	"fmt"
	"gover/internal/core"

	"github.com/spf13/cobra"
)

func installCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "install [version]",
		Short: "Install a specific Go version",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			version := "latest"
			if len(args) == 1 {
				version = args[0]
			}
			if err := core.Manager.Install(version); err != nil {
				return err
			}
			fmt.Printf("Go %s installed\n", version)
			return nil
		},
	}
}
