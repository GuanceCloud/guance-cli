package types

import (
	"reflect"
	"testing"
)

func TestPatchMap(t *testing.T) {
	type args struct {
		original       map[string]interface{}
		patch          map[string]interface{}
		reserveDefault bool
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
				reserveDefault: false,
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
				reserveDefault: false,
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
		{
			name: "defaults",
			args: args{
				original: map[string]interface{}{
					"key": map[string]interface{}{
						"pos": map[string]interface{}{
							"height": 5,
						},
					},
				},
				patch: map[string]interface{}{
					"key": map[string]interface{}{
						"pos": map[string]interface{}{
							"height": 10,
							"width":  10,
						},
					},
				},
				reserveDefault: true,
			},
			want: map[string]interface{}{
				"key": map[string]interface{}{
					"pos": map[string]interface{}{
						"height": 5,
						"width":  10,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PatchMap(tt.args.original, tt.args.patch, tt.args.reserveDefault); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PatchMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
