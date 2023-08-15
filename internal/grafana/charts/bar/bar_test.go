package bar_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/bar"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &bar.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/bar.grafana.json",
			GuanceFile:  "testdata/bar.guance.json",
			Chart:       chart,
		},
	})
}
