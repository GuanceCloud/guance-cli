package stat

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
	panel := grafanaspec.Panel{}
	if err := types.Decode(m, &panel); err != nil {
		return chart, fmt.Errorf("failed to decode panel: %w", err)
	}

	queries, err := prometheus.BuildTargets(panel.Targets, "singlestat")
	if err != nil {
		return chart, fmt.Errorf("failed to build targets: %w", err)
	}

	return map[string]interface{}{
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
				"isTimeInterval":   false,
				"fixedTime":        "",
				"timeInterval":     "default",
				"downsample":       "last",
				"showLine":         false,
				"showLineAxis":     false,
				"lineColor":        "#3AB8FF",
				"openCompare":      false,
				"compareType":      "",
				"bgColor":          "",
				"fontColor":        "",
				"precision":        "2",
			},
		},
	}, nil
}
