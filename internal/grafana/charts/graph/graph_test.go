package graph_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/graph"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &graph.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/graph.grafana.json",
			GuanceFile:  "testdata/graph.guance.json",
			Chart:       chart,
		},
	})
}
