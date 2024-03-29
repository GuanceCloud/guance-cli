package text

import (
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/base"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type Builder struct {
	base.Builder
}

func (builder Builder) Build(panel grafanaspec.Panel) (map[string]any, error) {
	content, ok := panel.Options["content"]
	if !ok {
		content = panel.Content
	}

	if content == "" {
		return nil, fmt.Errorf("failed to get text panel(%d) content, %+v", types.IntValue(panel.Id), panel.Options)
	}

	return map[string]any{
		"type": builder.Meta().GuanceType,
		"name": types.StringValue(panel.Title),
		"queries": []map[string]any{
			{
				"name": "",
				"query": map[string]any{
					"content": content,
				},
			},
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
			},
		},
	}, nil
}

const (
	ChartTypeText       = "text"
	GuanceChartTypeText = "text"
)

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypeText,
		GrafanaType: ChartTypeText,
	}
}
