package stat

import (
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons"
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons/queries"
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons/units"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/base"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"

	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type Builder struct {
	base.Builder
}

func (builder Builder) Build(panel grafanaspec.Panel) (map[string]any, error) {
	height := 5 // min-height
	if panel.GridPos.H > height {
		height = panel.GridPos.H
	}
	return map[string]any{
		"name": types.StringValue(panel.Title),
		"pos": map[string]any{
			"h": height,
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
		"type": builder.Meta().GuanceType,
	}, nil
}

const (
	ChartTypeStat             = "stat"
	GuanceChartTypeSingleStat = "singlestat"
)

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypeSingleStat,
		GrafanaType: ChartTypeStat,
	}
}

func (builder Builder) Addons() []addons.ChartAddon {
	return []addons.ChartAddon{
		&queries.Addon{Measurement: builder.Measurement},
		&units.Addon{Measurement: builder.Measurement},
	}
}