package pie_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/pie"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &pie.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/pie.grafana.json",
			GuanceFile:  "testdata/pie.guance.json",
			Chart:       chart,
		},
	})
}
