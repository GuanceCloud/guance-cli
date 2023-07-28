package prometheus

import (
	"fmt"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
	"github.com/hashicorp/go-multierror"
)

// Target is the query target of Grafana
type Target struct {
	Datasource   any    `json:"datasource"`
	Expr         string `json:"expr"`
	Hide         bool   `json:"hide"`
	Interval     string `json:"interval"`
	LegendFormat string `json:"legendFormat"`
	RefID        string `json:"refId"`
}

// BuildTargets builds targets for Guance Cloud
func (builder *Builder) BuildTargets(targets []grafanaspec.Target) ([]any, error) {
	var mErr error
	queries := make([]any, 0, len(targets))
	for _, item := range targets {
		target := Target{}
		if err := types.Decode(item, &target); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to decode target: %w", err))
			continue
		}
		promql, err := (&Rewriter{Measurement: builder.Measurement}).Rewrite(target.Expr)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to rewrite promql: %w", err))
			continue
		}
		queries = append(queries, map[string]any{
			"qtype": "promql",
			"query": map[string]any{
				"code":     target.RefID,
				"q":        promql,
				"type":     "promql",
				"funcList": []string{},
			},
			"type":       builder.ChartType,
			"datasource": "dataflux",
			"color":      "",
			"name":       "",
			"unit":       builder.Unit,
			"disabled":   target.Hide,
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}
	return queries, nil
}
