package variables

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGrafanaVariable(t *testing.T) {
	tests := []struct {
		expr    string
		wantErr bool
	}{
		{
			expr:    "label_names(cpu_seconds_total)",
			wantErr: true,
		},
		{
			expr:    "label_names(cpu_seconds_.*)",
			wantErr: true,
		},
		{
			expr:    "label_values(cpu_seconds_total)",
			wantErr: true,
		},
		{
			expr:    "label_values(cpu_seconds_total, app)",
			wantErr: true,
		},
		{
			expr:    `label_values(cpu_seconds_total{namespace="$namespace"})`,
			wantErr: true,
		},
		{
			expr:    `label_values(cpu_seconds_total{namespace="$namespace"}, app)`,
			wantErr: true,
		},
		{
			expr:    "metrics(cpu_seconds_total)",
			wantErr: true,
		},
		{
			expr:    "metrics(cpu_seconds_.*)",
			wantErr: true,
		},
		{
			expr:    `query_result(cpu_seconds_total{namespace="$namespace"})`,
			wantErr: true,
		},
		{
			expr:    `cpu_seconds_total{namespace="$namespace"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		dql, err := toDQL(tt.expr)
		if err != nil {
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			continue
		}
		fmt.Println(dql)
	}
}
