package testutil

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
)

type ChartSnapshotTest struct {
	Name        string
	GrafanaFile string
	GuanceFile  string
	Chart       charts.Builder
}

func TestChartSnapshots(t *testing.T, tests []ChartSnapshotTest) {
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			input := make(map[string]any)
			inputJSON, err := os.ReadFile(tt.GrafanaFile)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if err := json.Unmarshal(inputJSON, &input); !assert.NoError(t, err) {
				t.FailNow()
			}

			delete(input, "datasource")
			panel := grafanaspec.Panel{}
			if err := types.Decode(input, &panel); !assert.NoError(t, err) {
				t.FailNow()
			}

			actual, err := tt.Chart.Build(panel)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			for _, addon := range tt.Chart.Addons() {
				actual, err = addon.PatchChart(&panel, actual)
				if !assert.NoError(t, err) {
					t.FailNow()
				}
			}

			actualJSON, err := json.Marshal(actual)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			expectedJSON, err := os.ReadFile(tt.GuanceFile)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if !assert.JSONEq(t, string(expectedJSON), string(actualJSON)) {
				t.Log(string(actualJSON))
			}
		})
	}
}
