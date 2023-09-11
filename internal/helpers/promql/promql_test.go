package promql

import (
	"testing"

	"github.com/GuanceCloud/guance-cli/internal/helpers/prettier"
	"github.com/stretchr/testify/assert"
)

func TestRewriter_Rewrite(t *testing.T) {
	type fields struct {
		Measurement string
	}
	type args struct {
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "basic query",
			fields: fields{
				Measurement: "prom",
			},
			args: args{
				query: "sum(rate(container_cpu_usage_seconds_total{image!=\"\",name=~\"^k8s_.*\"}[5m])) by (pod_name)",
			},
			want:    "sum by (pod_name) (rate(prom:container_cpu_usage_seconds_total{image!=\"\", name=~\"^k8s_.*\"}[5m]))",
			wantErr: false,
		},
		{
			name: "no measurement",
			fields: fields{
				Measurement: "",
			},
			args: args{
				query: "sum(rate(container_cpu_usage_seconds_total{image!=\"\",name=~\"^k8s_.*\"}[5m])) by (pod_name)",
			},
			want:    "sum by (pod_name) (rate(container:cpu_usage_seconds_total{image!=\"\", name=~\"^k8s_.*\"}[5m]))",
			wantErr: false,
		},
		{
			name: "invalid query",
			fields: fields{
				Measurement: "prom",
			},
			args: args{
				query: "invalid query",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "drop grafana interval",
			fields: fields{
				Measurement: "prom",
			},
			args: args{
				query: "irate(node_forks_total{instance=\"#{node}\",job=\"#{job}\"}[$__rate_interval])",
			},
			want:    "irate(prom:node_forks_total{instance=\"#{node}\",job=\"#{job}\"})",
			wantErr: false,
		},
		{
			name: "escape measurement name",
			fields: fields{
				Measurement: "node-xxx",
			},
			args: args{
				query: "irate(node_forks_total{instance=\"#{node}\",job=\"#{job}\"}[$__rate_interval])",
			},
			want:    "irate(node\\-xxx:node_forks_total{instance=\"#{node}\",job=\"#{job}\"})",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Rewriter{
				Measurement: tt.fields.Measurement,
			}
			got, err := w.Rewrite(tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Rewriter.Rewrite() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !assert.Equal(t, prettier.RemoveSpaces(tt.want), prettier.RemoveSpaces(got), "expected: %s, got: %s", tt.want, got) {
				t.FailNow()
			}
		})
	}
}
