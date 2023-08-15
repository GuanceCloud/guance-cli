package promql

import "testing"

func TestGenName(t *testing.T) {
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				query: "sum by (instance) (rate(node_cpu_seconds_total{mode=\"idle\"}[5m]))",
			},
			want:    "sum by (instance) (rate(node_cpu_seconds_total{mode=\"idle\"}[5m]))",
			wantErr: false,
		},
		{
			name: "single",
			args: args{
				query: "prom:node_cpu_seconds_total",
			},
			want:    "node_cpu_seconds_total",
			wantErr: false,
		},
		{
			name: "single",
			args: args{
				query: "prom:node_cpu_seconds_total{mode=\"idle\"}}",
			},
			want:    "node_cpu_seconds_total",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenName(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenName() = %v, want %v", got, tt.want)
			}
		})
	}
}
