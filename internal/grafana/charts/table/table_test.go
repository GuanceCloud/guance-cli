package table_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/table"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &table.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/table.grafana.json",
			GuanceFile:  "testdata/table.guance.json",
			Chart:       chart,
		},
	})
}
