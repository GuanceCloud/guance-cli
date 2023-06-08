package prettier

import (
	"reflect"
	"testing"
)

func TestFormatJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    []byte
		expected []byte
		wantErr  bool
	}{
		{
			name:     "valid json",
			input:    []byte(`{"name":"John","age":30,"city":"New York"}`),
			expected: []byte("{\n  \"age\": 30,\n  \"city\": \"New York\",\n  \"name\": \"John\"\n}"),
			wantErr:  false,
		},
		{
			name:     "invalid json",
			input:    []byte(`{"name":"John","age":30,"city":"New York"`),
			expected: nil,
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := FormatJSON(tc.input)
			if (err != nil) != tc.wantErr {
				t.Errorf("FormatJSON() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("FormatJSON() = %v, want %v", string(actual), string(tc.expected))
			}
		})
	}
}
