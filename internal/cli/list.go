package cli

import (
	"gover/internal/core"

	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List installed versions",
		RunE: func(cmd *cobra.Command, args []string) error {
			versions, err := core.Manager.List()
			if err != nil {
				return err
			}
			for _, v := range versions {
				println(v)
			}
			return nil
		},
	}
}
