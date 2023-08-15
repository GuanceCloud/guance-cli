package units

import (
	"reflect"
	"testing"
)

func Test_convertUnit(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "miss",
			args: args{
				id: "unknown",
			},
			want: nil,
		},
		{
			name: "kWh",
			args: args{
				id: "kwatth",
			},
			want: []string{"energy", "kWh"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertUnit(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertUnit() = %v, want %v", got, tt.want)
			}
		})
	}
}
