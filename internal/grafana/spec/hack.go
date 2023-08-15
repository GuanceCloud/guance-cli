package spec

import (
	"encoding/json"
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/helpers/types"
	"github.com/hashicorp/go-multierror"

	"github.com/tidwall/gjson"
)

func (ref *DataSourceRef) UnmarshalJSON(data []byte) error {
	if gjson.GetBytes(data, ".").IsObject() {
		return json.Unmarshal(data, ref)
	}
	typeName := "template"
	ref.Type = &typeName
	return nil
}

// TypedTarget is the query target of Grafana
type TypedTarget struct {
	Datasource   any    `json:"datasource"`
	Expr         string `json:"expr"`
	Hide         bool   `json:"hide"`
	Interval     string `json:"interval"`
	LegendFormat string `json:"legendFormat"`
	RefID        string `json:"refId"`
}

// ParseTargets parses targets for Guance Cloud
func ParseTargets(values []Target) ([]TypedTarget, error) {
	var mErr error
	result := make([]TypedTarget, 0, len(values))
	for _, item := range values {
		target := TypedTarget{}
		if err := types.Decode(item, &target); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to decode target: %w", err))
			continue
		}
		result = append(result, target)
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to decode targets: %w", mErr)
	}
	return result, nil
}
