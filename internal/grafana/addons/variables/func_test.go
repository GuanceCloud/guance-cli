package variables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseGrafanaVariable(t *testing.T) {
	tests := []struct {
		expr     string
		expected string
		wantErr  bool
	}{
		{
			expr:     "label_names(cpu_seconds_total)",
			expected: "label_names(cpu_seconds_total)",
			wantErr:  false,
		},
		{
			expr:     "label_names(cpu_seconds_.*)",
			expected: "label_names(cpu_seconds_.*)",
			wantErr:  false,
		},
		{
			expr:     `label_values(app)`,
			expected: `label_values("app")`,
			wantErr:  false,
		},
		{
			expr:     `label_values(cpu_seconds_total, app)`,
			expected: `label_values(prom:cpu_seconds_total, "app")`,
			wantErr:  false,
		},
		{
			expr:     `label_values(cpu_seconds_total{namespace="$namespace"}, app)`,
			expected: `label_values(prom:cpu_seconds_total{namespace="#{namespace}"}, "app")`,
			wantErr:  false,
		},
		{
			expr:     "metrics(cpu_seconds_total)",
			expected: "metrics(cpu_seconds_total)",
			wantErr:  false,
		},
		{
			expr:     "metrics(cpu_seconds_.*)",
			expected: "metrics(cpu_seconds_.*)",
			wantErr:  false,
		},
		{
			expr:     `query_result(cpu_seconds_total{namespace="$namespace"})`,
			expected: `query_result(prom:cpu_seconds_total{namespace="#{namespace}"})`,
			wantErr:  false,
		},
		{
			expr:    `cpu_seconds_total{namespace="$namespace"}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.expr, func(t *testing.T) {
			v := &GrafanaVariable{
				Expr:        tt.expr,
				Measurement: "prom",
			}

			f, err := v.Func()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if f != nil && tt.expected != "" {
				assert.Equal(t, tt.expected, f.ToGuance())
			}
		})
	}
}
