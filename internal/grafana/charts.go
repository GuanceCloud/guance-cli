package grafana

import (
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/bar"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/bargauge"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/base"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/gauge"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/pie"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/stat"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/table"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/text"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/timeseries"
	"github.com/hashicorp/go-multierror"
)

type chartOptions struct {
	Type        string
	Measurement string
	Group       string
}

type ChartOption func(*chartOptions) error

func WithType(t string) ChartOption {
	return func(opts *chartOptions) error {
		opts.Type = t
		return nil
	}
}

func WithMeasurement(m string) ChartOption {
	return func(opts *chartOptions) error {
		opts.Measurement = m
		return nil
	}
}

func WithGroup(g string) ChartOption {
	return func(opts *chartOptions) error {
		opts.Group = g
		return nil
	}
}

func NewChartBuilder(opts ...ChartOption) (charts.Builder, error) {
	var mErr error
	options := &chartOptions{}
	for _, opt := range opts {
		if err := opt(options); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to apply chart option: %w", err))
		}
	}
	if mErr != nil {
		return nil, mErr
	}

	baseBuilder := base.Builder{
		Measurement: options.Measurement,
		Group:       options.Group,
	}

	var chartBuilder charts.Builder
	switch options.Type {
	case timeseries.ChartTypeTimeSeries:
		chartBuilder = &timeseries.Builder{Builder: baseBuilder}
	case bargauge.ChartTypeBarGauge:
		chartBuilder = &bargauge.Builder{Builder: baseBuilder}
	case gauge.ChartTypeGauge:
		chartBuilder = &gauge.Builder{Builder: baseBuilder}
	case stat.ChartTypeStat:
		chartBuilder = &stat.Builder{Builder: baseBuilder}
	case table.ChartTypeTable:
		chartBuilder = &table.Builder{Builder: baseBuilder}
	case pie.ChartTypePie:
		chartBuilder = &pie.Builder{Builder: baseBuilder}
	case bar.ChartTypeBarChart:
		chartBuilder = &bar.Builder{Builder: baseBuilder}
	case text.ChartTypeText:
		chartBuilder = &text.Builder{Builder: baseBuilder}
	default:
		return nil, fmt.Errorf("unknown chart type: %s", options.Type)
	}
	return chartBuilder, nil
}
