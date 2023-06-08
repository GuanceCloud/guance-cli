package spec

import (
	"encoding/json"

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
