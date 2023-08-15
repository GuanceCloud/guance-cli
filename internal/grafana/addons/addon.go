package addons

import (
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
)

type ChartAddon interface {
	PatchChart(panel *grafanaspec.Panel, chart map[string]any) (map[string]any, error)
}

type DashboardAddon interface {
	PatchDashboard(spec *grafanaspec.Spec, dashboard map[string]any) (map[string]any, error)
}
