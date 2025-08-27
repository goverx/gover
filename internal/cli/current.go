package cli

import (
	"fmt"
	"gover/internal/core"

	"github.com/spf13/cobra"
)

func currentCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "current",
		Short: "Show current Go version",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := core.Manager.Current()
			if err != nil {
				return err
			}
			fmt.Println(version)
			return nil
		},
	}
}
