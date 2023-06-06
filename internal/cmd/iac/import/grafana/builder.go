package grafana

import (
	"fmt"
	"sync"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/datasources/prometheus"

	"github.com/hashicorp/go-multierror"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

const DefaultMeasurement = "prom"

// Builder is the builder of Grafana dashboard
type Builder struct {
	Measurement string
	groups      []string
	charts      []map[string]any
	mu          sync.Mutex
}

// NewBuilder creates a new builder
func NewBuilder() *Builder {
	return &Builder{
		Measurement: DefaultMeasurement,
		groups:      make([]string, 0),
		charts:      make([]map[string]any, 0),
		mu:          sync.Mutex{},
	}
}

// Build will build Guance Cloud dashboard from Grafana dashboard
func (b *Builder) Build(src *dashboard.Spec) (map[string]any, error) {
	b.reset()

	// Build prometheus builder
	promBuilder := &prometheus.Builder{
		Measurement: b.Measurement,
	}

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
	variables, err := promBuilder.BuildVariables(src.Templating.List)
	if err != nil {
		return nil, fmt.Errorf("failed to build variables: %w", err)
	}

	return map[string]any{
		"dashboardBindSet":   []string{},
		"dashboardExtend":    map[string]any{},
		"dashboardMapping":   []string{},
		"dashboardOwnerType": "node",
		"dashboardType":      "CUSTOM",
		"iconSet":            map[string]any{},
		"main": map[string]any{
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

func (b *Builder) acceptPanel(v any) error {
	m := v.(map[string]any)
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
