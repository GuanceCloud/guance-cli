package charts

import (
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
)

type Meta struct {
	GuanceType  string
	GrafanaType string
}

type Builder interface {
	Build(panel grafanaspec.Panel) (map[string]any, error)
	Addons() []addons.ChartAddon
	Meta() Meta
}
