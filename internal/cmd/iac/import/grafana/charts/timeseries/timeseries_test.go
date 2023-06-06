package timeseries_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/chart"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/graph"
)

func TestChartBuilder_Build(t *testing.T) {
	for _, tt := range []struct {
		name    string
		grafana string
		guance  string
	}{
		{
			name:    "gauge",
			grafana: "testdata/timeseries.grafana.json",
			guance:  "testdata/timeseries.guance.json",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			input := make(map[string]any)
			inputJSON, err := os.ReadFile(tt.grafana)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if err := json.Unmarshal(inputJSON, &input); !assert.NoError(t, err) {
				t.FailNow()
			}

			builder := graph.ChartBuilder{}
			actual, err := builder.Build(input, chart.BuildOptions{})
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			actualJSON, err := json.Marshal(actual)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			fmt.Println(string(actualJSON))

			expectedJSON, err := os.ReadFile(tt.guance)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			assert.JSONEq(t, string(expectedJSON), string(actualJSON))
		})
	}
}
