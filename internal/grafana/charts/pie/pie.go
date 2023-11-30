package pie

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
		"group": map[string]any{
			"name": types.String(builder.Group),
		},
		"extend": map[string]any{
			"settings": map[string]any{
				"showTitle":                true,
				"titleDesc":                types.StringValue(panel.Description),
				"showFieldMapping":         false,
				"isTimeInterval":           false,
				"fixedTime":                "",
				"timeInterval":             "default",
				"chartType":                "pie",
				"legendPostion":            "bottom",
				"onlyShowGroupName":        false,
				"otherColor":               "#F56610",
				"chartCombineDefaultColor": "#F56610",
			},
		},
	}, nil
}

const (
	ChartTypePie       = "piechart"
	GuanceChartTypePie = "pie"
)

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypePie,
		GrafanaType: ChartTypePie,
	}
}
