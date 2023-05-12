package iac

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

var (
	//go:embed template/terraform-module/versions.tf
	templateTerraformVersion []byte

	//go:embed template/terraform-module/dashboard/main.tf
	templateTerraformDashboard []byte

	//go:embed template/terraform-module/dashboard/outputs.tf
	templateTerraformDashboardOutput []byte

	//go:embed template/terraform-module/dashboard/variables.tf
	templateTerraformDashboardVar []byte

	//go:embed template/terraform-module/monitor/main.tf.gotpl
	templateTerraformMonitor []byte

	//go:embed template/terraform-module/monitor/outputs.tf.gotpl
	templateTerraformMonitorOutput []byte

	//go:embed template/terraform-module/monitor/variables.tf
	templateTerraformMonitorVar []byte
)

const (
	// SourceTypeConsole is the type from Guance Cloud Console
	SourceTypeConsole = "console"
)

const (
	// ResourceTypeDashboard is the type for Guance Cloud Dashboard
	ResourceTypeDashboard = "dashboard"
	// ResourceTypeMonitor is the type for Guance Cloud Monitor
	ResourceTypeMonitor = "monitor"
)

const (
	// TargetTypeTerraform is the type for Terraform
	TargetTypeTerraform = "terraform"
	// TargetTypeTerraformModule is the type for Terraform Module
	TargetTypeTerraformModule = "terraform-module"
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

type importOptions struct {
	Resource string
	Target   string
	File     string
	Out      string
}

type genFile struct {
	path    string
	content []byte
}

func newCmdImport() *cobra.Command {
	opts := importOptions{}
	cmd := &cobra.Command{
		Use:       "import",
		Short:     "Import external resource as Guance Cloud IaC resource",
		Args:      cobra.ExactArgs(1),
		ValidArgs: []string{SourceTypeConsole},
		RunE: func(cmd *cobra.Command, args []string) error {
			// The source must be imported from Guance Cloud Console now.
			// Add more checking at here if we support more source type in the future.

			var files []*genFile
			switch opts.Target {
			case TargetTypeTerraformModule:
				// Copy the manifest file import from console
				content, err := os.ReadFile(opts.File)
				if err != nil {
					return fmt.Errorf("read file %s error: %w", opts.File, err)
				}

				// Write the provider.tf
				files = append(files, &genFile{path: "versions.tf", content: templateTerraformVersion})

				// Write the resource file
				switch opts.Resource {
				case ResourceTypeDashboard:
					files = append(files, &genFile{path: "manifest.json", content: content})
					files = append(files, &genFile{path: "main.tf", content: templateTerraformDashboard})
					files = append(files, &genFile{path: "variables.tf", content: templateTerraformDashboardVar})
					files = append(files, &genFile{path: "outputs.tf", content: templateTerraformDashboardOutput})
				case ResourceTypeMonitor:
					var indices []string
					jsonModifiers := []jsonModifier{fixNoData, formatJSON}
					for i, value := range gjson.GetBytes(content, "checkers").Array() {
						valueContent := []byte(value.String())
						for _, modifier := range jsonModifiers {
							valueContent, err = modifier(valueContent)
							if err != nil {
								return fmt.Errorf("fix data error: %w", err)
							}
						}
						files = append(files, &genFile{path: fmt.Sprintf("manifest-%03d.json", i+1), content: valueContent})
						indices = append(indices, fmt.Sprintf("%03d", i+1))
					}
					outputFile, err := renderFile(templateTerraformMonitorOutput, indices)
					if err != nil {
						return fmt.Errorf("render outputs file error: %w", err)
					}
					files = append(files, &genFile{path: "outputs.tf", content: outputFile})
					mainFile, err := renderFile(templateTerraformMonitor, indices)
					if err != nil {
						return fmt.Errorf("render main file error: %w", err)
					}
					files = append(files, &genFile{path: "main.tf", content: mainFile})
					files = append(files, &genFile{path: "variables.tf", content: templateTerraformMonitorVar})
				default:
					return fmt.Errorf("resource type %s not supported", opts.Resource)
				}
			default:
				return fmt.Errorf("target type %s not supported", opts.Target)
			}
			for _, f := range files {
				if err := writeFile(fmt.Sprintf("%s/%s", opts.Out, f.path), f.content); err != nil {
					return fmt.Errorf("write file %s error: %w", f.path, err)
				}
			}
			return nil
		},
	}
	cmd.Flags().StringVarP(&opts.Resource, "resource", "r", "", "Source type, supports dashboard, monitor now.")
	cmd.Flags().StringVarP(&opts.File, "file", "f", "", "File path to import.")
	cmd.Flags().StringVarP(&opts.Target, "target", "t", "", "Target type, supports terraform-module now.")
	cmd.Flags().StringVarP(&opts.Out, "out", "o", "", "Output file path.")
	_ = cmd.MarkFlagRequired("target")
	_ = cmd.MarkFlagRequired("out")
	cmd.MarkFlagsRequiredTogether("resource", "file")
	return cmd
}

func writeFile(dst string, content []byte) error {
	if err := os.MkdirAll(filepath.Dir(dst), 0o755); err != nil {
		return fmt.Errorf("mkdir %s error: %w", dst, err)
	}
	if err := os.WriteFile(dst, content, 0o600); err != nil {
		return fmt.Errorf("write file %s error: %w", dst, err)
	}
	return nil
}

func renderFile(templateContent []byte, data interface{}) ([]byte, error) {
	tpl, err := template.New("guance").Parse(string(templateContent))
	if err != nil {
		return nil, fmt.Errorf("parse template error: %w", err)
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("execute template error: %w", err)
	}
	return buf.Bytes(), nil
}

type jsonModifier func([]byte) ([]byte, error)

func fixNoData(src []byte) ([]byte, error) {
	var err error
	if gjson.GetBytes(src, "jsonScript.noDataPeriodCount").Int() == 0 {
		src, err = sjson.DeleteBytes(src, "jsonScript.noDataPeriodCount")
		if err != nil {
			return nil, fmt.Errorf("delete noDataPeriodCount error: %w", err)
		}
	}
	if gjson.GetBytes(src, "extend.noDataPeriodCount").Int() == 0 {
		src, err = sjson.DeleteBytes(src, "extend.noDataPeriodCount")
		if err != nil {
			return nil, fmt.Errorf("delete noDataPeriodCount error: %w", err)
		}
	}
	if gjson.GetBytes(src, "jsonScript.noDataMessage").String() == "" {
		src, err = sjson.DeleteBytes(src, "jsonScript.noDataMessage")
		if err != nil {
			return nil, fmt.Errorf("delete noDataMessage error: %w", err)
		}
	}
	if gjson.GetBytes(src, "jsonScript.noDataTitle").String() == "" {
		src, err = sjson.DeleteBytes(src, "jsonScript.noDataTitle")
		if err != nil {
			return nil, fmt.Errorf("delete noDataTitle error: %w", err)
		}
	}
	if !gjson.GetBytes(src, "jsonScript.checkerOpt.infoEvent").Bool() {
		src, err = sjson.DeleteBytes(src, "jsonScript.checkerOpt.infoEvent")
		if err != nil {
			return nil, fmt.Errorf("delete infoEvent error: %w", err)
		}
	}
	return src, nil
}

func formatJSON(src []byte) ([]byte, error) {
	var m interface{}
	if err := json.Unmarshal(src, &m); err != nil {
		return nil, fmt.Errorf("unmarshal json error when formatting: %w", err)
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal json error when formatting: %w", err)
	}
	return b, nil
}
