package variables

import (
	"fmt"

	grafanaspec "github.com/GuanceCloud/guance-cli/internal/grafana/spec"
)

type Addon struct {
	Measurement string
}

func (addon *Addon) PatchDashboard(spec *grafanaspec.Spec, dashboard map[string]any) (map[string]any, error) {
	variables, err := addon.BuildVariables(spec.Templating.List)
	if err != nil {
		return nil, fmt.Errorf("failed to build targets: %w", err)
	}
	mainValue, ok := dashboard["main"]
	if !ok {
		return nil, fmt.Errorf("failed to get main")
	}
	main, ok := mainValue.(map[string]any)
	if !ok {
		return nil, fmt.Errorf("failed to cast main")
	}
	main["vars"] = variables
	dashboard["main"] = main
	return dashboard, nil
}
