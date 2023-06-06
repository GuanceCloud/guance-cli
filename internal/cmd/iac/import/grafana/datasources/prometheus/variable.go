package prometheus

import (
	"fmt"
	"regexp"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
	"github.com/hashicorp/go-multierror"
)

// BuildVariables builds variables for Guance Cloud
func (builder *Builder) BuildVariables(variables []grafanaspec.VariableModel) ([]any, error) {
	var mErr error
	vars := make([]any, 0, len(variables))
	for _, variable := range variables {
		if variable.Type != "query" {
			continue
		}

		label, err := getLabelFromVariable(variable)
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
				"value":  fmt.Sprintf("SHOW_TAG_VALUE(from=['%s'], keyin=['%s'])", builder.Measurement, label),
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

var labelFuncPattern = regexp.MustCompile(`label_values\((.+),\s*(.+)\)`)

func getLabelFromVariable(variable grafanaspec.VariableModel) (string, error) {
	if variable.Query == nil {
		return "", fmt.Errorf("query %s is empty", variable.Name)
	}
	m, ok := (*variable.Query).(map[string]any)
	if !ok {
		return "", fmt.Errorf("failed to decode query %s map", variable.Name)
	}
	queryString, ok := m["query"].(string)
	if !ok {
		return "", fmt.Errorf("failed to get query string from variable: %s", variable.Name)
	}
	match := labelFuncPattern.FindStringSubmatch(queryString)
	if len(match) != 3 {
		return "", fmt.Errorf("failed to get label from variable: %s", variable.Name)
	}
	return match[2], nil
}
