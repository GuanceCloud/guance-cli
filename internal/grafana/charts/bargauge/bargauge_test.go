package bargauge_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/bargauge"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &bargauge.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/bargauge.grafana.json",
			GuanceFile:  "testdata/bargauge.guance.json",
			Chart:       chart,
		},
	})
}
