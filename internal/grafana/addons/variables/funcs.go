package variables

// PromQLFuncLabelNames is the label_names function in Grafana
type GrafanaFuncLabelNames struct {
	MetricRegexp string
}

// PromQLFuncLabelValues is the label_values function in Grafana
type GrafanaFuncLabelValues struct {
	Metric string
	Label  string
}

// PromQLFuncQueryResult is the query_result function in Grafana
type GrafanaFuncQueryResult struct {
	Query string
}

// PromQLFuncMetrics is the metrics function in Grafana
type GrafanaFuncMetrics struct {
	MetricRegexp string
}
