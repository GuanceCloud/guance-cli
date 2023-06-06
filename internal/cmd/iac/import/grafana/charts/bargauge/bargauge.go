package bargauge

import (
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/datasources/prometheus"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type ChartBuilder struct {
	Type string
}

func (builder *ChartBuilder) Build(m map[string]any, opts chart.BuildOptions) (chart map[string]any, err error) {
	delete(m, "datasource")

	panel := grafanaspec.Panel{}
	if err := types.Decode(m, &panel); err != nil {
		return nil, fmt.Errorf("failed to decode panel: %w", err)
	}

	queries, err := (&prometheus.Builder{
		Measurement: opts.Measurement,
		ChartType:   builder.Type,
	}).BuildTargets(panel.Targets)
	if err != nil {
		return nil, fmt.Errorf("failed to build targets: %w", err)
	}

	return map[string]any{
		"type": builder.Type,
		"name": types.StringValue(panel.Title),
		"pos": map[string]any{
			"h": panel.GridPos.H,
			"w": panel.GridPos.W,
			"x": panel.GridPos.X,
			"y": panel.GridPos.Y,
		},
		"queries": queries,
		"group": map[string]any{
			"name": types.String(opts.Group),
		},
		"extend": map[string]any{
			"settings": map[string]any{
				"showTitle":         true,
				"titleDesc":         types.StringValue(panel.Description),
				"showFieldMapping":  false,
				"isTimeInterval":    false,
				"fixedTime":         "",
				"timeInterval":      "default",
				"direction":         "vertical",
				"showTopSize":       false,
				"topSize":           10,
				"showTopWithMetric": "",
				"xAxisShowType":     "groupBy",
				"showLine":          false,
				"openCompare":       false,
				"openStack":         false,
				"stackContent":      "group",
				"stackType":         "time",
			},
		},
	}, nil
}
