package timeseries

import (
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/base"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type Builder struct {
	base.Builder
}

func (builder Builder) Build(panel grafanaspec.Panel) (map[string]any, error) {
	return map[string]any{
		"type": builder.Meta().GuanceType,
		"name": types.StringValue(panel.Title),
		"pos": map[string]any{
			"h": panel.GridPos.H,
			"w": panel.GridPos.W,
			"x": panel.GridPos.X,
			"y": panel.GridPos.Y,
		},
		"group": map[string]any{
			"name": types.String(builder.Group),
		},
		"extend": map[string]any{
			"settings": map[string]any{
				"showTitle":        true,
				"titleDesc":        types.StringValue(panel.Description),
				"showFieldMapping": false,
				"timeInterval":     "auto",
				"chartType":        "areaLine",
			},
		},
	}, nil
}

const (
	ChartTypeTimeSeries     = "timeseries"
	GuanceChartTypeSequence = "sequence"
)

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypeSequence,
		GrafanaType: ChartTypeTimeSeries,
	}
}
