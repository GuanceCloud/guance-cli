package variables

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-multierror"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
)

// BuildVariables builds variables for Guance Cloud
func (addon *Addon) BuildVariables(variables []grafanaspec.VariableModel) ([]any, error) {
	var mErr error
	vars := make([]any, 0, len(variables))
	for _, variable := range variables {
		if variable.Type != "query" {
			continue
		}

		dqlQuery, err := addon.toDQL(variable)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to get label from variable: %w", err))
			continue
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
				"value":  dqlQuery,
			},
			"hide":             0,
			"isHiddenAsterisk": 0,
			"name":             variable.Label,
			"seq":              2,
			"type":             "QUERY",
			"valueSort":        "asc",
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}
	return vars, nil
}

var labelFuncPattern = regexp.MustCompile(`label_values\((.+),?\s*(.*)\)`)

func (addon *Addon) toDQL(variable grafanaspec.VariableModel) (string, error) {
	queryString, err := getPromExpr(variable)
	if err != nil {
		return "", fmt.Errorf("failed to get prometheus expression: %w", err)
	}

	switch {
	case strings.HasPrefix(queryString, "label_values("):
		match := labelFuncPattern.FindStringSubmatch(queryString)
		if len(match) != 3 {
			return "", fmt.Errorf("failed to get label from variable: %s", variable.Name)
		}
		return fmt.Sprintf("SHOW_TAG_VALUE(from=['%s'], keyin=['%s'])", addon.Measurement, match[2]), nil
	case strings.HasPrefix(queryString, "query_result("):
		return "", nil
	default:
		return "", fmt.Errorf("failed to get label from variable: %s", variable.Name)
	}
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
