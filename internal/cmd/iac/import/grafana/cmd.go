package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/grafana-tools/sdk"
	"github.com/hashicorp/go-multierror"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"

	dashboardtfmod "github.com/GuanceCloud/guance-cli/internal/generator/tfmod/resources/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/grafana"
)

type importOptions struct {
	Target         string
	Out            string
	Measurement    string
	Files          []string
	TemplateID     string
	Search         bool
	SearchID       int
	SearchFolderID int
	SearchQuery    string
	SearchTag      string
}

func NewCmd() *cobra.Command {
	opts := importOptions{}
	cmd := &cobra.Command{
		Use:   "grafana",
		Short: "Import Grafana Dashboard resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			var err error
			var dashboards []dashboardtfmod.Manifest

			if opts.Search {
				dashboards, err = searchDashboards(context.Background(), &opts)
			} else if len(opts.Files) != 0 {
				dashboards, err = readDashboards(context.Background(), &opts)
			} else if opts.TemplateID != "" {
				dashboards, err = importDashboardTemplate(context.Background(), &opts)
			} else {
				return fmt.Errorf("file, template or search must be specified")
			}
			if err != nil {
				return fmt.Errorf("discovery dashboards error: %w", err)
			}

			var result []dashboardtfmod.Manifest
			for _, d := range dashboards {
				grafanaDashboard, err := grafana.ParseGrafana(d.Content)
				if err != nil {
					return fmt.Errorf("parse grafana dashboard error: %w", err)
				}

				builder := grafana.NewBuilder()
				builder.Measurement = opts.Measurement
				adt, err := builder.Build(grafanaDashboard)
				if err != nil {
					return fmt.Errorf("generate dashboard error: %w", err)
				}

				manifest, err := json.Marshal(adt)
				if err != nil {
					return fmt.Errorf("marshal dashboard error: %w", err)
				}
				result = append(result, dashboardtfmod.Manifest{
					Name:    d.Name,
					Title:   d.Title,
					Content: manifest,
				})
			}

			files, err := dashboardtfmod.Generate(dashboardtfmod.Options{Manifests: result})
			if err != nil {
				return fmt.Errorf("generate dashboard error: %w", err)
			}
			return files.Save(opts.Out)
		},
	}

	cmd.Flags().StringSliceVarP(&opts.Files, "file", "f", nil, "File path to import.")
	cmd.Flags().StringVar(&opts.TemplateID, "template-id", "", "Template ID to import.")
	cmd.Flags().IntVar(&opts.SearchID, "search-id", 0, "Dashboard id to import.")
	cmd.Flags().IntVar(&opts.SearchFolderID, "search-folder-id", 0, "Folder id to import.")
	cmd.Flags().StringVar(&opts.SearchQuery, "search-query", "", "Query to search dashboard.")
	cmd.Flags().StringVar(&opts.SearchTag, "search-tag", "", "Tag to search dashboard.")
	cmd.Flags().BoolVar(&opts.Search, "search", false, "Search dashboard.")
	cmd.Flags().StringVarP(&opts.Target, "target", "t", "terraform-module", "Target type, supports terraform-module now.")
	cmd.Flags().StringVarP(&opts.Out, "out", "o", "out", "Output file path.")
	cmd.Flags().StringVarP(&opts.Measurement, "measurement", "m", "", "Measurement (default is prom).")
	cmd.MarkFlagsMutuallyExclusive("file", "search")
	return cmd
}

func readDashboards(ctx context.Context, opts *importOptions) ([]dashboardtfmod.Manifest, error) {
	var mErr error
	result := make([]dashboardtfmod.Manifest, 0, len(opts.Files))
	for i, filePath := range opts.Files {
		content, err := os.ReadFile(filePath)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("read file %s error: %w", filePath, err))
			continue
		}

		result = append(result, dashboardtfmod.Manifest{
			Name:    fmt.Sprintf("dashboard_%d", i),
			Title:   gjson.GetBytes(content, "title").String(),
			Content: content,
		})
	}
	return result, nil
}

func importDashboardTemplate(ctx context.Context, opts *importOptions) ([]dashboardtfmod.Manifest, error) {
	resp, err := http.Get(fmt.Sprintf("https://grafana.com/api/dashboards/%s/revisions/latest/download", opts.TemplateID))
	if err != nil {
		return nil, fmt.Errorf("download template %s failed: %w", opts.TemplateID, err)
	}
	fmt.Println("Downloaded Grafana Template", opts.TemplateID)
	defer func() { _ = resp.Body.Close() }()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read stream of template %s failed", opts.TemplateID)
	}
	return []dashboardtfmod.Manifest{
		{
			Name:    fmt.Sprintf("template_%s", opts.TemplateID),
			Title:   gjson.GetBytes(content, "title").String(),
			Content: content,
		},
	}, nil
}

func searchDashboards(ctx context.Context, opts *importOptions) ([]dashboardtfmod.Manifest, error) {
	c, err := sdk.NewClient(
		os.Getenv("GRAFANA_URL"),
		os.Getenv("GRAFANA_AUTH"),
		sdk.DefaultHTTPClient)
	if err != nil {
		return nil, fmt.Errorf("failed to create a client: %s", err)
	}

	params := []sdk.SearchParam{
		sdk.SearchType(sdk.SearchTypeDashboard),
	}
	if opts.SearchID != 0 {
		params = append(params, sdk.SearchDashboardID(opts.SearchID))
	}
	if opts.SearchFolderID != 0 {
		params = append(params, sdk.SearchFolderID(opts.SearchFolderID))
	}
	if opts.SearchQuery != "" {
		params = append(params, sdk.SearchQuery(opts.SearchQuery))
	}
	if opts.SearchTag != "" {
		params = append(params, sdk.SearchTag(opts.SearchTag))
	}

	boardLinks, err := c.Search(ctx, params...)
	if err != nil {
		return nil, err
	}

	var mErr error
	result := make([]dashboardtfmod.Manifest, 0, len(boardLinks))
	for _, link := range boardLinks {
		rawBoard, meta, err := c.GetRawDashboardByUID(ctx, link.UID)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to get dashboard %s: %s", link.UID, err))
			continue
		}
		result = append(result, dashboardtfmod.Manifest{
			Name:    meta.Slug,
			Content: rawBoard,
			Title:   link.Title,
		})
		fmt.Printf("Downloaded %q\n", link.Title)
	}
	return result, nil
}
