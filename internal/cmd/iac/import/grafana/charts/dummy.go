package charts

import (
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
)

type dummyChartBuilder struct {
	Type string
}

func (builder *dummyChartBuilder) Build(m map[string]interface{}, opts chart.BuildOptions) (chart map[string]interface{}, err error) {
	return chart, fmt.Errorf("chart type %s not implemented", builder.Type)
}
