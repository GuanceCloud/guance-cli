package graph

import (
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/datasources/prometheus"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type ChartBuilder struct{}

func (builder *ChartBuilder) Build(m map[string]interface{}, opts chart.BuildOptions) (chart map[string]interface{}, err error) {
	delete(m, "datasource")

	panel := grafanaspec.Panel{}
	if err := types.Decode(m, &panel); err != nil {
		return chart, fmt.Errorf("failed to decode panel: %w", err)
	}

	queries, err := (&prometheus.Builder{
		Measurement: opts.Measurement,
		ChartType:   "sequence",
	}).BuildTargets(panel.Targets)
	if err != nil {
		return chart, fmt.Errorf("failed to build targets: %w", err)
	}

	chartType := ""
	if isTrue(m, "lines") {
		chartType = "line"
	} else if isTrue(m, "bars") {
		chartType = "bar"
	}

	return map[string]interface{}{
		"type": "sequence",
		"name": types.StringValue(panel.Title),
		"pos": map[string]interface{}{
			"h": panel.GridPos.H,
			"w": panel.GridPos.W,
			"x": panel.GridPos.X,
			"y": panel.GridPos.Y,
		},
		"queries": queries,
		"group": map[string]interface{}{
			"name": types.String(opts.Group),
		},
		"extend": map[string]interface{}{
			"settings": map[string]interface{}{
				"chartType":    chartType,
				"timeInterval": "auto",
			},
		},
	}, nil
}

func isTrue(m map[string]interface{}, key string) bool {
	if v, ok := m[key]; ok {
		if b, ok := v.(bool); ok {
			return b
		}
	}
	return false
}
