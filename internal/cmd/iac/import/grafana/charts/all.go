package charts

import (
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/chart"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/bar"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/bargauge"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/gauge"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/pie"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/stat"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/table"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/text"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/timeseries"
)

const (
	ChartTypeTimeSeries = "timeseries"
	ChartTypeBarGauge   = "bargauge"
	ChartTypeGauge      = "gauge"
	ChartTypeStat       = "stat"
	ChartTypeTable      = "table"
	ChartTypePie        = "piechart"
	ChartTypeBarChart   = "barchart"
	ChartTypeText       = "text"
)

const (
	GuanceChartTypeSequence   = "sequence"
	GuanceChartTypeSingleStat = "singlestat"
	GuanceChartTypeTable      = "table"
	GuanceChartTypeGauge      = "gauge"
	GuanceChartTypeBar        = "bar"
	GuanceChartTypePie        = "pie"
	GuanceChartTypeText       = "text"
)

var charts map[string]chart.Builder

func init() {
	charts = make(map[string]chart.Builder)
	charts[ChartTypeTimeSeries] = &timeseries.ChartBuilder{Type: GuanceChartTypeSequence}
	charts[ChartTypeBarGauge] = &bargauge.ChartBuilder{Type: GuanceChartTypeBar}
	charts[ChartTypeGauge] = &gauge.ChartBuilder{Type: GuanceChartTypeGauge}
	charts[ChartTypeStat] = &stat.ChartBuilder{Type: GuanceChartTypeSingleStat}
	charts[ChartTypeTable] = &table.ChartBuilder{Type: GuanceChartTypeTable}
	charts[ChartTypePie] = &pie.ChartBuilder{Type: GuanceChartTypePie}
	charts[ChartTypeBarChart] = &bar.ChartBuilder{Type: GuanceChartTypeBar}
	charts[ChartTypeText] = &text.ChartBuilder{Type: GuanceChartTypeText}
}

func NewChartBuilder(chartType string) chart.Builder {
	if b, ok := charts[chartType]; ok {
		return b
	}
	return &dummyChartBuilder{Type: chartType}
}
