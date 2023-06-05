package prometheus

import (
	"fmt"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
	"github.com/hashicorp/go-multierror"
)

type Target struct {
	Datasource   any    `json:"datasource"`
	Expr         string `json:"expr"`
	Hide         bool   `json:"hide"`
	Interval     string `json:"interval"`
	LegendFormat string `json:"legendFormat"`
	RefId        string `json:"refId"`
}

func BuildTargets(targets []grafanaspec.Target, chartType string) ([]interface{}, error) {
	var mErr error
	var queries []interface{}
	for _, item := range targets {
		target := Target{}
		if err := types.Decode(item, &target); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to decode target: %w", err))
			continue
		}
		promql, err := (&Rewriter{Measurement: "prom"}).Rewrite(target.Expr)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to rewrite promql: %w", err))
			continue
		}
		queries = append(queries, map[string]interface{}{
			"qtype": "promql",
			"query": map[string]interface{}{
				"code":     target.RefId,
				"q":        promql,
				"type":     "promql",
				"funcList": []string{},
			},
			"type":       chartType,
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
