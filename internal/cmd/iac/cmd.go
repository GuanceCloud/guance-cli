package iac

import (
	_ "embed"

	"github.com/spf13/cobra"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/console"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana"
)

// NewCmd create a new iac command
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "iac",
		Short: "Infrastructure as Code",
	}
	cmd.AddCommand(newCmdImport())
	return cmd
}

func newCmdImport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import external resource as Guance Cloud IaC resource",
	}
	cmd.AddCommand(console.NewCmd())
	cmd.AddCommand(grafana.NewCmd())
	return cmd
}
