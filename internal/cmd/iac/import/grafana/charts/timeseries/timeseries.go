package timeseries

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

func (builder *ChartBuilder) Build(m map[string]interface{}, opts chart.BuildOptions) (chart map[string]interface{}, err error) {
	delete(m, "datasource")

	panel := grafanaspec.Panel{}
	if err := types.Decode(m, &panel); err != nil {
		return nil, fmt.Errorf("failed to decode panel: %w", err)
	}

	queries, err := (&prometheus.Builder{
		Measurement: opts.Measurement,
		ChartType:   "sequence",
	}).BuildTargets(panel.Targets)
	if err != nil {
		return nil, fmt.Errorf("failed to build targets: %w", err)
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
				"showTitle":        true,
				"titleDesc":        types.StringValue(panel.Description),
				"showFieldMapping": false,
				"timeInterval":     "auto",
			},
		},
	}, nil
}
