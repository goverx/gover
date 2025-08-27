package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gover",
	Short: "Gover is a Go version manager",
	Long:  "Lightweight Go version manager, like nvm but for Go.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(installCmd())
	rootCmd.AddCommand(useCmd())
	rootCmd.AddCommand(listCmd())
	rootCmd.AddCommand(currentCmd())
}
