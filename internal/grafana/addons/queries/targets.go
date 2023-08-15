package queries

import (
	"fmt"

	"github.com/hashicorp/go-multierror"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/promql"
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
func (addon *Addon) BuildTargets(targetValues []grafanaspec.Target) ([]map[string]any, error) {
	var mErr error
	queries := make([]map[string]any, 0, len(targetValues))

	targets, err := grafanaspec.ParseTargets(targetValues)
	if err != nil {
		return nil, fmt.Errorf("failed to parse targets: %w", err)
	}

	for _, target := range targets {
		q, err := (&promql.Rewriter{Measurement: addon.Measurement}).Rewrite(target.Expr)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to rewrite promql: %w", err))
			continue
		}
		queries = append(queries, map[string]any{
			"qtype": "promql",
			"query": map[string]any{
				"q":        q,
				"type":     "promql",
				"funcList": []string{},
			},
			"datasource": "dataflux",
			"color":      "",
			"name":       "",
			"disabled":   target.Hide,
			"type":       "",
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}
	return queries, nil
}
