package grafana

import (
	"encoding/json"
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/dashboard"
)

// ParseGrafana will parse grafana dashboard json bytes as a json model struct
func ParseGrafana(content []byte) (*dashboard.Spec, error) {
	dashboard := dashboard.Spec{}
	if err := json.Unmarshal(content, &dashboard); err != nil {
		return nil, fmt.Errorf("parsing grafana dashboard, unmarshal json error: %w", err)
	}
	return &dashboard, nil
}
