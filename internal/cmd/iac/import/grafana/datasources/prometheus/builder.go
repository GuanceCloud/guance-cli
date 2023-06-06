package prometheus

import (
	"fmt"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
	"github.com/hashicorp/go-multierror"
)

// Target is the query target of prometheus
type Target struct {
	Datasource   any    `json:"datasource"`
	Expr         string `json:"expr"`
	Hide         bool   `json:"hide"`
	Interval     string `json:"interval"`
	LegendFormat string `json:"legendFormat"`
	RefID        string `json:"refId"`
}

type Builder struct {
	Measurement string
	ChartType   string
}

func (builder *Builder) BuildTargets(targets []grafanaspec.Target) ([]interface{}, error) {
	var mErr error
	queries := make([]interface{}, 0, len(targets))
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
		queries = append(queries, map[string]interface{}{
			"qtype": "promql",
			"query": map[string]interface{}{
				"code":     target.RefID,
				"q":        promql,
				"type":     "promql",
				"funcList": []string{},
			},
			"type":       builder.ChartType,
			"datasource": "dataflux",
			"color":      "",
			"name":       "",
			"disabled":   target.Hide,
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}
	return queries, nil
}
