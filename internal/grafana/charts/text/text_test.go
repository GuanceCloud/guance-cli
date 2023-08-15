package text_test

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts/text"
	"github.com/GuanceCloud/guance-cli/internal/grafana/testutil"
)

func TestChartBuilder_Build(t *testing.T) {
	chart := &text.Builder{Builder: testutil.BaseChartBuilder}
	testutil.TestChartSnapshots(t, []testutil.ChartSnapshotTest{
		{
			Name:        "ok",
			GrafanaFile: "testdata/text.grafana.json",
			GuanceFile:  "testdata/text.guance.json",
			Chart:       chart,
		},
	})
}
