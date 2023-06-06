package grafana

import (
	"fmt"
	"sync"

	"github.com/hashicorp/go-multierror"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type Builder struct {
	groups []string
	charts []map[string]interface{}
	mu     sync.Mutex
}

func NewBuilder() *Builder {
	return &Builder{
		groups: make([]string, 0),
		charts: make([]map[string]interface{}, 0),
		mu:     sync.Mutex{},
	}
}

func (b *Builder) Build(src *dashboard.Spec) (map[string]interface{}, error) {
	b.reset()

	// Build panels
	var mErr error
	for _, panel := range src.Panels {
		if err := b.acceptPanel(panel); err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to build dashboard: %w", mErr)
	}

	// Build Variables
	variables := make([]map[string]interface{}, 0)
	for _, variable := range src.Templating.List {
		variables = append(variables, map[string]interface{}{
			"code":       variable.Name,
			"datasource": "dataflux",
			"definition": map[string]any{
				"defaultVal": map[string]any{
					"label": "",
					"value": "",
				},
				"field":  "",
				"metric": "",
				"object": "",
				"tag":    "",
				// TODO: implement syntax-directed translation
				"value": "SHOW_TAG_VALUE(from=['prom'], keyin=['instance']) LIMIT 50",
			},
			"hide":             0,
			"isHiddenAsterisk": 0,
			"name":             variable.Label,
			"seq":              2,
			"type":             "QUERY",
			"valueSort":        "asc",
		})
	}

	return map[string]interface{}{
		"dashboardBindSet":   []string{},
		"dashboardExtend":    map[string]interface{}{},
		"dashboardMapping":   []string{},
		"dashboardOwnerType": "node",
		"dashboardType":      "CUSTOM",
		"iconSet":            map[string]interface{}{},
		"main": map[string]interface{}{
			"charts": b.charts,
			"groups": b.groups,
			"type":   "template",
			"vars":   variables,
		},
		"summary":   types.StringValue(src.Description),
		"title":     types.StringValue(src.Title),
		"thumbnail": "",
		"tags":      src.Tags,
		"tagInfo":   []any{},
	}, nil
}

func (b *Builder) acceptPanel(v interface{}) error {
	m := v.(map[string]interface{})
	chartType, ok := m["type"].(string)
	if !ok {
		return fmt.Errorf("failed to get chart type")
	}
	delete(m, "datasource")

	// Convert Grafana row as Guance Cloud group
	if chartType == "row" {
		panel := &dashboard.RowPanel{}
		if err := types.Decode(m, &panel); err != nil {
			return fmt.Errorf("failed to decode row panel: %w", err)
		}
		b.groups = append(b.groups, types.StringValue(panel.Title))

		if len(panel.Panels) > 0 {
			for _, p := range panel.Panels {
				if err := b.acceptPanel(p); err != nil {
					return err
				}
			}
		}
		return nil
	}

	// Convert Grafana panel to Guance Cloud chart
	builder := charts.NewChartBuilder(chartType)
	chartManifest, err := builder.Build(m, chart.BuildOptions{
		Group:       b.currentGroup(),
		Measurement: "prom",
	})
	if err != nil {
		return fmt.Errorf("failed to build chart: %w", err)
	}
	b.charts = append(b.charts, chartManifest)
	return nil
}

func (b *Builder) currentGroup() string {
	if len(b.groups) == 0 {
		return ""
	}
	return b.groups[len(b.groups)-1]
}

func (b *Builder) reset() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.groups = nil
	b.charts = nil
}
