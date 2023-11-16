package prettier

import (
	"encoding/json"
	"fmt"
)

// FormatJSON will execute formatting for json bytes with indent
func FormatJSON(src []byte) ([]byte, error) {
	var m any
	if err := json.Unmarshal(src, &m); err != nil {
		return nil, fmt.Errorf("unmarshal json error when formatting: %w", err)
	}
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("marshal json error when formatting: %w", err)
	}
	return b, nil
}
