package stat_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/stat"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &stat.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/stat.grafana.json",
			GuanceFile:  "testdata/stat.guance.json",
			Chart:       chart,
		},
	})
}
