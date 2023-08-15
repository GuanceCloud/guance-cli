package queries

import (
	"fmt"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
)

type Addon struct {
	Measurement string
}

func (addon *Addon) PatchChart(panel *grafanaspec.Panel, chart map[string]any) (map[string]any, error) {
	chartType, ok := chart["type"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to get chart type")
	}

	queries, err := addon.BuildTargets(panel.Targets)
	if err != nil {
		return nil, fmt.Errorf("failed to build targets: %w", err)
	}

	for _, query := range queries {
		query["type"] = chartType
	}

	chart["queries"] = queries
	return chart, nil
}
