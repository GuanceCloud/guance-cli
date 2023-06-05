package grafana

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	dashboardtfmod "github.com/GuanceCloud/guance-cli/internal/generator/tfmod/resources/dashboard"
)

type importOptions struct {
	Resource string
	Target   string
	File     string
	Out      string
}

func NewCmd() *cobra.Command {
	opts := importOptions{}
	cmd := &cobra.Command{
		Use:   "grafana",
		Short: "(Alpha) Import grafana resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			content, err := os.ReadFile(opts.File)
			if err != nil {
				return fmt.Errorf("read file error: %w", err)
			}

			grafanaDashboard, err := ParseGrafana(content)
			if err != nil {
				return fmt.Errorf("parse grafana dashboard error: %w", err)
			}

			builder := NewBuilder()
			adt, err := builder.Build(grafanaDashboard)
			if err != nil {
				return fmt.Errorf("generate dashboard error: %w", err)
			}

			manifest, err := json.Marshal(adt)
			if err != nil {
				return fmt.Errorf("marshal dashboard error: %w", err)
			}

			fmt.Println(string(manifest))

			files, err := dashboardtfmod.Generate(dashboardtfmod.Options{Manifest: manifest})
			if err != nil {
				return fmt.Errorf("generate dashboard error: %w", err)
			}
			return files.Save(opts.Out)
		},
	}
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "File path to import.")
	cmd.Flags().StringVarP(&opts.Target, "target", "t", "", "Target type, supports terraform-module now.")
	cmd.Flags().StringVarP(&opts.Out, "out", "o", "", "Output file path.")
	_ = cmd.MarkFlagRequired("target")
	_ = cmd.MarkFlagRequired("out")
	cmd.MarkFlagsRequiredTogether("file")
	return cmd
}
