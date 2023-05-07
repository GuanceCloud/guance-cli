package cmd

import (
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "guance",
		Short: "Guance Cloud Command-Line Interface",
	}
	rootCmd.AddCommand(iac.NewCmd())
	return rootCmd
}
