package grafana

import (
	"fmt"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts"
	chart2 "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	"sync"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
	"github.com/hashicorp/go-multierror"
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

	var mErr error
	for _, panel := range src.Panels {
		if err := b.acceptPanel(panel); err != nil {
			mErr = multierror.Append(mErr, err)
		}
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to build dashboard: %w", mErr)
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
		},
		"summary":   types.StringValue(src.Description),
		"title":     types.StringValue(src.Title),
		"thumbnail": "",
		"tags":      []string{},
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
	chart, err := builder.Build(m, chart2.BuildOptions{Group: b.currentGroup()})
	if err != nil {
		return fmt.Errorf("failed to build chart: %w", err)
	}
	b.charts = append(b.charts, chart)
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
