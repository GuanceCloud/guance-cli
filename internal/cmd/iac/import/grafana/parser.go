package grafana

import (
	"encoding/json"
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/cmd/iac/import/grafana/spec"
)

// ParseGrafana will parse grafana dashboard json bytes as a json model struct
func ParseGrafana(content []byte) (*spec.Spec, error) {
	dashboard := spec.Spec{}
	if err := json.Unmarshal(content, &dashboard); err != nil {
		return nil, fmt.Errorf("parsing grafana dashboard, unmarshal json error: %w", err)
	}
	return &dashboard, nil
}
