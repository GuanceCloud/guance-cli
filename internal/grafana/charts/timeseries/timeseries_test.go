package timeseries_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/timeseries"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &timeseries.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/timeseries.grafana.json",
			GuanceFile:  "testdata/timeseries.guance.json",
			Chart:       chart,
		},
	})
}
