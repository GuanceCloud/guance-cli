package singlestat_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/singlestat"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &singlestat.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/singlestat.grafana.json",
			GuanceFile:  "testdata/singlestat.guance.json",
			Chart:       chart,
		},
	})
}
