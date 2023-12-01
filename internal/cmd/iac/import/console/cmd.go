package console

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"

	"github.com/GuanceCloud/guance-cli/internal/generator/tfmod/resources/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/generator/tfmod/resources/monitor"
	"github.com/GuanceCloud/guance-cli/internal/helpers/osfs"
)

const (
	// ResourceTypeDashboard is the type for Guance Cloud Dashboard
	ResourceTypeDashboard = "dashboard"
	// ResourceTypeMonitor is the type for Guance Cloud Monitor
	ResourceTypeMonitor = "monitor"
)

const (
	// TargetTypeTerraformModule is the type for Terraform Module
	TargetTypeTerraformModule = "terraform-module"
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
		Use:   "console",
		Short: "Import Guance Cloud Console resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch opts.Target {
			case TargetTypeTerraformModule:
				return generateTerraformModule(opts)
			default:
				return fmt.Errorf("target type %s not supported", opts.Target)
			}
		},
	}
	cmd.Flags().StringVarP(&opts.Resource, "resource", "r", "", "Source type, supports dashboard, monitor now.")
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "File path to import.")
	cmd.Flags().StringVarP(&opts.Target, "target", "t", "terraform-module", "Target type, supports terraform-module now.")
	cmd.Flags().StringVarP(&opts.Out, "out", "o", "out", "Output file path.")
	cmd.MarkFlagsRequiredTogether("resource", "file")
	return cmd
}

func generateTerraformModule(opts importOptions) error {
	content, err := os.ReadFile(opts.File)
	if err != nil {
		return fmt.Errorf("read file %s error: %w", opts.File, err)
	}
	var files osfs.Files
	switch opts.Resource {
	case ResourceTypeDashboard:
		files, err = dashboard.Generate(dashboard.Options{Manifests: []dashboard.Manifest{
			{
				Name:    "manifest",
				Title:   gjson.GetBytes(content, "title").String(),
				Content: content,
			},
		}})
	case ResourceTypeMonitor:
		var monitors []json.RawMessage
		for _, value := range gjson.GetBytes(content, "checkers").Array() {
			monitors = append(monitors, json.RawMessage(value.Raw))
		}
		files, err = monitor.Generate(monitor.Options{Manifests: monitors})
	default:
		return fmt.Errorf("resource type %s not supported", opts.Resource)
	}
	if err != nil {
		return fmt.Errorf("generate terraform module error: %w", err)
	}
	if err = files.Save(opts.Out); err != nil {
		return fmt.Errorf("save terraform module error: %w", err)
	}
	return nil
}
