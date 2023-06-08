package table

import (
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/chart"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/datasources/prometheus"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type ChartBuilder struct {
	Type string
}

func (builder *ChartBuilder) Build(m map[string]any, opts chart.BuildOptions) (chart map[string]any, err error) {
	delete(m, "datasource")

	panel := grafanaspec.Panel{}
	if err := types.Decode(m, &panel); err != nil {
		return chart, fmt.Errorf("failed to decode panel: %w", err)
	}

	queries, err := (&prometheus.Builder{
		Measurement: opts.Measurement,
		ChartType:   builder.Type,
	}).BuildTargets(panel.Targets)
	if err != nil {
		return chart, fmt.Errorf("failed to build targets: %w", err)
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
				"showTitle":                true,
				"titleDesc":                types.StringValue(panel.Description),
				"showFieldMapping":         false,
				"isTimeInterval":           false,
				"fixedTime":                "",
				"timeInterval":             "default",
				"queryMode":                "toGroupColumn",
				"pageSize":                 0,
				"mainMeasurementQueryCode": "A",
				"mainMeasurementSort":      "top",
				"mainMeasurementLimit":     20,
			},
		},
	}, nil
}
