package testutil

import "github.com/GuanceCloud/guance-cli/internal/grafana/charts/base"

// BaseChartBuilder is a base chart builder for testing.
var BaseChartBuilder = base.Builder{
	Measurement: "prom",
	Group:       "",
}
