package gridpos

import (
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

// Addon is a builder to convert unit to the chart
type Addon struct {
	Measurement string
}

func (addon *Addon) PatchChart(panel *grafanaspec.Panel, chart map[string]any) (map[string]any, error) {
	if panel.GridPos == nil {
		return chart, nil
	}
	patch := map[string]any{
		"pos": map[string]any{
			"h": panel.GridPos.H,
			"w": panel.GridPos.W,
			"x": panel.GridPos.X,
			"y": panel.GridPos.Y,
		},
	}
	result := types.PatchMap(chart, patch, true)
	return result, nil
}
