package variables

import (
	"fmt"
	"regexp"
)

var (
	// PromQLLabelNamesFuncName is the label_names function in Grafana
	PrometheusLabelNamesRegex = regexp.MustCompile(`^label_names\(\)\s*$`)

	// PromQLLabelValuesFuncName is the label_values function in Grafana
	PrometheusLabelValuesRegex = regexp.MustCompile(`^label_values\((?:(.+),\s*)?([a-zA-Z_$][a-zA-Z0-9_]*)\)\s*$`)

	// PromQLMetricsFuncName is the metrics function in Grafana
	PrometheusMetricNamesRegex = regexp.MustCompile(`^metrics\((.+)\)\s*$`)

	// PromQLQueryResultFuncName is the query_result function in Grafana
	PrometheusQueryResultRegex = regexp.MustCompile(`^query_result\((.+)\)\s*$`)

	// PromQLLabelNamesFuncNameWithMatch is the label_names function in Grafana
	PrometheusLabelNamesRegexWithMatch = regexp.MustCompile(`^label_names\((.+)\)\s*$`)
)

const (
	// PrometheusLabelNamesFuncName is the label_names function name in Grafana
	PrometheusLabelNamesFuncName = "label_names"

	// PrometheusLabelValuesFuncName is the label_values function name in Grafana
	PrometheusLabelValuesFuncName = "label_values"

	// PrometheusMetricsFuncName is the metrics function name in Grafana
	PrometheusMetricNamesFuncName = "metrics"

	// PrometheusQueryResultFuncName is the query_result function name in Grafana
	PrometheusQueryResultFuncName = "query_result"
)

type GrafanaVariable struct {
	Expr     string
	FuncName string
	Func     interface{}
}

func ParseGrafanaVariable(expr string) (*GrafanaVariable, error) {
	switch {
	case PrometheusLabelValuesRegex.MatchString(expr):
		match := PrometheusLabelValuesRegex.FindStringSubmatch(expr)
		if len(match) != 3 {
			return nil, fmt.Errorf("failed to get label from variable: %s", expr)
		}
		return &GrafanaVariable{
			Expr:     expr,
			FuncName: PrometheusLabelValuesFuncName,
			Func: &GrafanaFuncLabelValues{
				Metric: match[1],
				Label:  match[2],
			},
		}, nil
	case PrometheusQueryResultRegex.MatchString(expr):
		match := PrometheusQueryResultRegex.FindStringSubmatch(expr)
		if len(match) != 2 {
			return nil, fmt.Errorf("failed to get query: %s", expr)
		}
		return &GrafanaVariable{
			Expr:     expr,
			FuncName: PrometheusQueryResultFuncName,
			Func: &GrafanaFuncQueryResult{
				Query: match[1],
			},
		}, nil
	case PrometheusMetricNamesRegex.MatchString(expr):
		match := PrometheusMetricNamesRegex.FindStringSubmatch(expr)
		if len(match) != 2 {
			return nil, fmt.Errorf("failed to get metric: %s", expr)
		}
		return &GrafanaVariable{
			Expr:     expr,
			FuncName: PrometheusMetricNamesFuncName,
			Func: &GrafanaFuncMetrics{
				MetricRegexp: match[1],
			},
		}, nil
	case PrometheusLabelNamesRegexWithMatch.MatchString(expr):
		match := PrometheusLabelNamesRegexWithMatch.FindStringSubmatch(expr)
		if len(match) != 2 {
			return nil, fmt.Errorf("failed to get metric: %s", expr)
		}
		return &GrafanaVariable{
			Expr:     expr,
			FuncName: PrometheusLabelNamesFuncName,
			Func: &GrafanaFuncLabelNames{
				MetricRegexp: match[1],
			},
		}, nil
	case PrometheusLabelNamesRegex.MatchString(expr):
		return &GrafanaVariable{
			Expr:     expr,
			FuncName: PrometheusLabelNamesFuncName,
			Func: &GrafanaFuncLabelNames{
				MetricRegexp: "",
			},
		}, nil
	default:
		return &GrafanaVariable{
			Expr:     expr,
			FuncName: "",
			Func:     nil,
		}, nil
	}
}

func toDQL(promExpr string) (string, error) {
	grafanaVariable, err := ParseGrafanaVariable(promExpr)
	if err != nil {
		return "", fmt.Errorf("failed to parse grafana variable: %w", err)
	}

	switch grafanaVariable.FuncName {
	case PrometheusLabelNamesFuncName:
		return "", nil
	case PrometheusLabelValuesFuncName:
		return "", nil
	case PrometheusQueryResultFuncName:
		return "", nil
	case PrometheusMetricNamesFuncName:
		return "", nil
	default:
		return "", nil
	}
}
