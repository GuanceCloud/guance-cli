package bar

import (
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/base"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type Builder struct {
	base.Builder
}

func (builder *Builder) Build(panel grafanaspec.Panel) (map[string]any, error) {
	return map[string]any{
		"type": builder.Meta().GuanceType,
		"name": types.StringValue(panel.Title),
		"group": map[string]any{
			"name": types.String(builder.Group),
		},
		"extend": map[string]any{
			"settings": map[string]any{
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
				"showTitle":         true,
				"titleDesc":         types.StringValue(panel.Description),
				"showFieldMapping":  false,
				"isTimeInterval":    false,
				"fixedTime":         "",
				"timeInterval":      "default",
			},
		},
	}, nil
}

const (
	ChartTypeBarChart  = "barchart"
	GuanceChartTypeBar = "bar"
)

func (builder *Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypeBar,
		GrafanaType: ChartTypeBarChart,
	}
}
