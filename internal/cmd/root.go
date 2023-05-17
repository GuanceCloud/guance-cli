package cmd

import (
	"github.com/spf13/cobra"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "guance",
		Short: "Guance Cloud Command-Line Interface",
	}
	rootCmd.AddCommand(iac.NewCmd())
	return rootCmd
}
