package gauge

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
		return nil, fmt.Errorf("failed to decode panel: %w", err)
	}

	queries, err := (&prometheus.Builder{
		Measurement: opts.Measurement,
		ChartType:   builder.Type,
	}).BuildTargets(panel.Targets)
	if err != nil {
		return nil, fmt.Errorf("failed to build targets: %w", err)
	}

	if panel.FieldConfig.Defaults.Thresholds.Mode != "absolute" {
		return nil, fmt.Errorf("threshold mode only supported absolute")
	}

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
		"type": builder.Type,
		"name": types.StringValue(panel.Title),
		"pos": map[string]any{
			"h": height,
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
				"showTitle":        true,
				"titleDesc":        types.StringValue(panel.Description),
				"showFieldMapping": false,
				"timeInterval":     "auto",
				"levels":           levels,
			},
		},
	}, nil
}
