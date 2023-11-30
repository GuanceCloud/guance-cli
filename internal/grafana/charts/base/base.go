package base

import (
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons"
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons/gridpos"
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons/queries"
	"github.com/GuanceCloud/guance-cli/internal/grafana/addons/units"
	"github.com/GuanceCloud/guance-cli/internal/grafana/charts"
	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
)

type Builder struct {
	Measurement string
	Group       string
}

func (builder Builder) Addons() []addons.ChartAddon {
	return []addons.ChartAddon{
		&queries.Addon{Measurement: builder.Measurement},
		&units.Addon{Measurement: builder.Measurement},
		&gridpos.Addon{Measurement: builder.Measurement},
	}
}

func (builder Builder) Meta() charts.Meta {
	return charts.Meta{
		GuanceType:  "dummy",
		GrafanaType: "unknown",
	}
}

func (builder Builder) Build(panel grafanaspec.Panel) (map[string]any, error) {
	panic("implement me")
}
