package gauge

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
	var currentTop float64
	var levels []map[string]any
	stepCnt := len(panel.FieldConfig.Defaults.Thresholds.Steps)
	for i := 0; i < stepCnt; i++ {
		// the next step value in Grafana is the current value (current top) in Guance Cloud
		currentStep := panel.FieldConfig.Defaults.Thresholds.Steps[i]
		if i == stepCnt-1 {
			currentTop = 100
		} else {
			nextStep := panel.FieldConfig.Defaults.Thresholds.Steps[i+1]
			currentTop = float64(types.Float32Value(nextStep.Value))
		}
		levels = append(levels, map[string]any{
			"operation": "<=",
			"value": []any{
				currentTop,
			},
			"lineColor": currentStep.Color,
		})
	}

	height := 10 // min-height
	if panel.GridPos.H > height {
		height = panel.GridPos.H
	}
	return map[string]any{
		"type": builder.Meta().GuanceType,
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
				"timeInterval":     "auto",
				"levels":           levels,
				"min":              panel.FieldConfig.Defaults.Min,
				"max":              panel.FieldConfig.Defaults.Max,
			},
		},
	}, nil
}

const (
	ChartTypeGauge       = "gauge"
	GuanceChartTypeGauge = "gauge"
)

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  GuanceChartTypeGauge,
		GrafanaType: ChartTypeGauge,
	}
}
