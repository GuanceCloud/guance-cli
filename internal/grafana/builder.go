package grafana

import (
	"fmt"
	"sync"

	"github.com/GuanceCloud/guance-cli/internal/grafana/addons"
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons/variables"
	"github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"

	"github.com/hashicorp/go-multierror"
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
func (b *Builder) Build(src *spec.Spec) (map[string]any, error) {
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

	dashboard := map[string]any{
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
		},
		"summary":   types.StringValue(src.Description),
		"title":     types.StringValue(src.Title),
		"thumbnail": "",
		"tags":      src.Tags,
		"tagInfo":   []any{},
	}

	var err error
	for _, addon := range []addons.DashboardAddon{
		&variables.Addon{Measurement: b.Measurement},
	} {
		dashboard, err = addon.PatchDashboard(src, dashboard)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to patch dashboard: %w", err))
		}
	}
	return dashboard, mErr
}

func (b *Builder) acceptPanel(v any) error {
	m := v.(map[string]any)
	chartType, ok := m["type"].(string)
	if !ok {
		return fmt.Errorf("failed to get chart type")
	}
	delete(m, "datasource")

	// Convert Grafana row as Guance Cloud group
	switch chartType {
	case "row":
		panel := &spec.RowPanel{}
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
	default:
		panel := spec.Panel{}
		if err := types.Decode(m, &panel); err != nil {
			return fmt.Errorf("failed to decode panel: %w", err)
		}

		builder, err := NewChartBuilder(
			WithType(chartType),
			WithGroup(b.currentGroup()),
			WithMeasurement(b.Measurement),
		)
		if err != nil {
			return fmt.Errorf("failed to create chart builder: %w", err)
		}

		chart, err := builder.Build(panel)
		if err != nil {
			return fmt.Errorf("failed to build chart: %w", err)
		}

		var mErr error
		for _, addon := range builder.Addons() {
			chart, err = addon.PatchChart(&panel, chart)
			if err != nil {
				mErr = multierror.Append(mErr, err)
			}
		}
		if mErr != nil {
			return fmt.Errorf("failed to patch chart: %w", mErr)
		}

		b.charts = append(b.charts, chart)
	}
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

	b.groups = make([]string, 0)
	b.charts = make([]map[string]any, 0)
}
