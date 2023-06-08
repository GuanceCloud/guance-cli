package charts

import (
	"fmt"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/chart"
)

type dummyChartBuilder struct {
	Type string
}

func (builder *dummyChartBuilder) Build(m map[string]any, opts chart.BuildOptions) (chart map[string]any, err error) {
	return chart, fmt.Errorf("chart type %s not implemented", builder.Type)
}
