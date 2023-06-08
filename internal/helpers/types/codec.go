package types

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func Decode(in map[string]any, out any) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           out,
		TagName:          "json",
		WeaklyTypedInput: true,
	})
	if err != nil {
		return fmt.Errorf("new decoder failed: %w", err)
	}
	if err := decoder.Decode(in); err != nil {
		return fmt.Errorf("decode map failed: %w", err)
	}
	return decoder.Decode(in)
}
