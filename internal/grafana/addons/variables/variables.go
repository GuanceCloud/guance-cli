package variables

import (
	"fmt"
	"github.com/GuanceCloud/guance-cli/internal/helpers/promql"

	"github.com/hashicorp/go-multierror"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

// BuildVariables builds variables for Guance Cloud
func (addon *Addon) BuildVariables(variables []grafanaspec.VariableModel) ([]any, error) {
	var mErr error

	vars := make([]any, 0, len(variables))
	for _, variable := range variables {
		if variable.Type != "query" {
			continue
		}

		promExpr, err := getPromExpr(variable)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to get prometheus expression from variable: %w", err))
			continue
		}

		q, err := (&promql.Rewriter{Measurement: addon.Measurement}).Rewrite(promExpr)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to rewrite promql: %w", err))
			continue
		}

		name := types.StringValue(variable.Label)
		if name == "" {
			name = variable.Name
		}

		vars = append(vars, map[string]any{
			"code":       variable.Name,
			"datasource": "dataflux",
			"definition": map[string]any{
				"defaultVal": map[string]any{
					"label": "",
					"value": "",
				},
				"field":  "",
				"metric": "",
				"object": "",
				"tag":    "",
				"value":  q,
			},
			"hide":             0,
			"isHiddenAsterisk": 0,
			"name":             name,
			"seq":              2,
			"type":             "PROMQL_QUERY",
			"valueSort":        "asc",
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}
	return vars, nil
}

func getPromExpr(variable grafanaspec.VariableModel) (string, error) {
	if variable.Query == nil {
		return "", fmt.Errorf("query %s is empty", variable.Name)
	}

	switch t := (*variable.Query).(type) {
	case string:
		return t, nil
	case map[string]any:
		queryString, ok := t["query"].(string)
		if !ok {
			return "", fmt.Errorf("failed to get query string from variable: %s", variable.Name)
		}
		return queryString, nil
	default:
		return "", fmt.Errorf("failed to get query string from variable: %s", variable.Name)
	}
}
