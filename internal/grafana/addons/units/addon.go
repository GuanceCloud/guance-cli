package units

import (
	_ "embed"
	"fmt"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/promql"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
	"github.com/hashicorp/go-multierror"
)

// Addon is a builder to convert unit to the chart
type Addon struct {
	Measurement string
}

func (addon *Addon) PatchChart(panel *grafanaspec.Panel, chart map[string]any) (map[string]any, error) {
	grafanaUnit := types.StringValue(panel.FieldConfig.Defaults.Unit)
	if grafanaUnit == "" {
		grafanaUnit = types.StringValue(panel.Format)
	}

	units := convertUnit(grafanaUnit)
	if units == nil {
		return chart, nil
	}

	targets, err := grafanaspec.ParseTargets(panel.Targets)
	if err != nil {
		return nil, fmt.Errorf("failed to parse targets: %w", err)
	}

	var mErr error
	result := make([]map[string]any, 0, len(targets))
	for _, target := range targets {
		q, err := promql.NewRewriter(addon.Measurement).Rewrite(target.Expr)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to rewrite promql: %w", err))
			continue
		}
		queryID, err := promql.GenName(q)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to generate query id: %w", err))
			continue
		}
		result = append(result, map[string]any{
			"key":   queryID,
			"name":  queryID,
			"units": units,
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}

	return types.PatchMap(chart, map[string]any{
		"extend": map[string]any{
			"settings": map[string]any{
				"units": result,
			},
		},
	}, false), nil
}
