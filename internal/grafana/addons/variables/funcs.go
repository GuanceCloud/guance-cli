package variables

import (
	"fmt"
	"regexp"

	"github.com/GuanceCloud/guance-cli/internal/helpers/promql"
)

var (
	// GrafanaFuncLabelNamesRegex is the label_names function in Grafana
	GrafanaFuncLabelNamesRegex = regexp.MustCompile(`^label_names\(\)\s*$`)

	// GrafanaFuncLabelNamesRegexWithMatch is the label_names function in Grafana
	GrafanaFuncLabelNamesRegexWithMatch = regexp.MustCompile(`^label_names\((.+)\)\s*$`)

	// GrafanaFuncLabelValuesRegex is the label_values function in Grafana
	GrafanaFuncLabelValuesRegex = regexp.MustCompile(`^label_values\((?:(.+),\s*)?([a-zA-Z_$][a-zA-Z0-9_]*)\)\s*$`)

	// GrafanaFuncMetricNamesRegex is the metrics function in Grafana
	GrafanaFuncMetricNamesRegex = regexp.MustCompile(`^metrics\((.+)\)\s*$`)

	// GrafanaFuncQueryResultRegex is the query_result function in Grafana
	GrafanaFuncQueryResultRegex = regexp.MustCompile(`^query_result\((.+)\)\s*$`)
)

const (
	// GrafanaFuncLabelNamesFuncName is the label_names function name in Grafana
	GrafanaFuncLabelNamesFuncName = "label_names"

	// GrafanaFuncLabelValuesFuncName is the label_values function name in Grafana
	GrafanaFuncLabelValuesFuncName = "label_values"

	// GrafanaFuncMetricNamesFuncName is the metrics function name in Grafana
	GrafanaFuncMetricNamesFuncName = "metrics"

	// GrafanaFuncQueryResultFuncName is the query_result function name in Grafana
	GrafanaFuncQueryResultFuncName = "query_result"
)

// GrafanaVariable is a struct to represent a variable load from Grafana
type GrafanaVariable struct {
	Expr        string
	Measurement string
}

// GrafanaFuncLabelNames is the label_names function in Grafana
type GrafanaFuncLabelNames struct {
	MetricRegexp string
}

func (v *GrafanaFuncLabelNames) ToGuance() string {
	if v.MetricRegexp == "" {
		return "label_names()"
	}
	return fmt.Sprintf("label_names(%s)", v.MetricRegexp)
}

// GrafanaFuncLabelValues is the label_values function in Grafana
type GrafanaFuncLabelValues struct {
	Measurement string
	Metric      string
	Label       string
}

func (v *GrafanaFuncLabelValues) ToGuance() string {
	if v.Metric == "" {
		return fmt.Sprintf("label_values(%q)", v.Label)
	}
	q, err := promql.NewRewriter(v.Measurement).Rewrite(v.Metric)
	if err != nil {
		q = v.Metric
	}
	return fmt.Sprintf("label_values(%s, %q)", q, v.Label)
}

// GrafanaFuncQueryResult is the query_result function in Grafana
type GrafanaFuncQueryResult struct {
	Measurement string
	Query       string
}

func (v *GrafanaFuncQueryResult) ToGuance() string {
	q, err := promql.NewRewriter(v.Measurement).Rewrite(v.Query)
	if err != nil {
		q = v.Query
	}
	return fmt.Sprintf("query_result(%s)", q)
}

// GrafanaFuncMetrics is the metrics function in Grafana
type GrafanaFuncMetrics struct {
	MetricRegexp string
}

func (v *GrafanaFuncMetrics) ToGuance() string {
	return fmt.Sprintf("metrics(%s)", v.MetricRegexp)
}

type GrafanaVariableFunc interface {
	ToGuance() string
}

func (v *GrafanaVariable) Func() (GrafanaVariableFunc, error) {
	switch {
	case GrafanaFuncLabelValuesRegex.MatchString(v.Expr):
		match := GrafanaFuncLabelValuesRegex.FindStringSubmatch(v.Expr)
		if len(match) != 3 {
			return nil, fmt.Errorf("failed to parse function %s: %s", GrafanaFuncLabelValuesFuncName, v.Expr)
		}
		return &GrafanaFuncLabelValues{
			Measurement: v.Measurement,
			Metric:      match[1],
			Label:       match[2],
		}, nil
	case GrafanaFuncQueryResultRegex.MatchString(v.Expr):
		match := GrafanaFuncQueryResultRegex.FindStringSubmatch(v.Expr)
		if len(match) != 2 {
			return nil, fmt.Errorf("failed to parse function %s: %s", GrafanaFuncQueryResultFuncName, v.Expr)
		}
		return &GrafanaFuncQueryResult{
			Measurement: v.Measurement,
			Query:       match[1],
		}, nil
	case GrafanaFuncMetricNamesRegex.MatchString(v.Expr):
		match := GrafanaFuncMetricNamesRegex.FindStringSubmatch(v.Expr)
		if len(match) != 2 {
			return nil, fmt.Errorf("failed to parse function %s: %s", GrafanaFuncMetricNamesFuncName, v.Expr)
		}
		return &GrafanaFuncMetrics{
			MetricRegexp: match[1],
		}, nil
	case GrafanaFuncLabelNamesRegexWithMatch.MatchString(v.Expr):
		match := GrafanaFuncLabelNamesRegexWithMatch.FindStringSubmatch(v.Expr)
		if len(match) != 2 {
			return nil, fmt.Errorf("failed to parse function %s: %s", GrafanaFuncLabelNamesFuncName, v.Expr)
		}
		return &GrafanaFuncLabelNames{
			MetricRegexp: match[1],
		}, nil
	case GrafanaFuncLabelNamesRegex.MatchString(v.Expr):
		return &GrafanaFuncLabelNames{
			MetricRegexp: "",
		}, nil
	default:
		return nil, fmt.Errorf("grafana variable function must be set")
	}
}
