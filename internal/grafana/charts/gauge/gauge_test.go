package gauge_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/gauge"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &gauge.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/gauge.grafana.json",
			GuanceFile:  "testdata/gauge.guance.json",
			Chart:       chart,
		},
	})
}
