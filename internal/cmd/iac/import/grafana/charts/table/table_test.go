package table_test

import (
	"encoding/json"
	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/chart"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/charts/table"
)

func TestChartBuilder_Build(t *testing.T) {
	for _, tt := range []struct {
		name    string
		grafana string
		guance  string
	}{
		{
			name:    "ok",
			grafana: "testdata/table.grafana.json",
			guance:  "testdata/table.guance.json",
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

			builder := table.ChartBuilder{Type: "table"}
			actual, err := builder.Build(input, chart.BuildOptions{Measurement: "prom"})
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			actualJSON, err := json.Marshal(actual)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			expectedJSON, err := os.ReadFile(tt.guance)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
			if !assert.JSONEq(t, string(expectedJSON), string(actualJSON)) {
				t.Log(string(actualJSON))
			}
		})
	}
}
