package charts

import (
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/bargauge"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/gauge"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/graph"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/stat"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/table"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/timeseries"
)

const (
	ChartTypeTimeSeries = "timeseries"
	ChartTypeBarGauge   = "bargauge"
	ChartTypeGauge      = "gauge"
	ChartTypeGraph      = "graph"
	ChartTypeHeatmap    = "heatmap"
	ChartTypeStat       = "stat"
	ChartTypeTable      = "table"
)

var charts map[string]chart.Builder

func init() {
	charts = make(map[string]chart.Builder)
	charts[ChartTypeTimeSeries] = &timeseries.ChartBuilder{}
	charts[ChartTypeBarGauge] = &bargauge.ChartBuilder{}
	charts[ChartTypeGauge] = &gauge.ChartBuilder{}
	charts[ChartTypeGraph] = &graph.ChartBuilder{}
	charts[ChartTypeHeatmap] = &dummyChartBuilder{Type: ChartTypeHeatmap}
	charts[ChartTypeStat] = &stat.ChartBuilder{}
	charts[ChartTypeTable] = &table.ChartBuilder{}
}

func NewChartBuilder(chartType string) chart.Builder {
	if b, ok := charts[chartType]; ok {
		return b
	}
	return &dummyChartBuilder{Type: chartType}
}
