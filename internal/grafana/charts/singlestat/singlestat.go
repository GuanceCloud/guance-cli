package singlestat

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
		"name": types.StringValue(panel.Title),
		"pos": map[string]any{
			"h": 5,
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
	ChartTypeSingleStat       = "singlestat"
	GuanceChartTypeSingleStat = "singlestat"
)

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypeSingleStat,
		GrafanaType: ChartTypeSingleStat,
	}
}
