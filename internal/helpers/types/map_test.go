package types

import (
	"reflect"
	"testing"
)

func TestPatchMap(t *testing.T) {
	type args struct {
		original map[string]interface{}
		patch    map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want map[string]interface{}
	}{
		{
			name: "single",
			args: args{
				original: map[string]interface{}{
					"key": "foo",
				},
				patch: map[string]interface{}{
					"key": "bar",
					"foo": "bar",
				},
			},
			want: map[string]interface{}{
				"key": "bar",
				"foo": "bar",
			},
		},
		{
			name: "nested",
			args: args{
				original: map[string]interface{}{
					"key": map[string]interface{}{
						"foo": map[string]interface{}{
							"bar": 0,
						},
					},
				},
				patch: map[string]interface{}{
					"key": map[string]interface{}{
						"foo": map[string]interface{}{
							"bar": 42,
							"key": 42,
						},
					},
				},
			},
			want: map[string]interface{}{
				"key": map[string]interface{}{
					"foo": map[string]interface{}{
						"bar": 42,
						"key": 42,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PatchMap(tt.args.original, tt.args.patch); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatchMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
