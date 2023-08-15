package grafana

import (
	_ "embed"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		name    string
		path    string
		wantErr bool
	}{
		{
			name: "example",
			path: "testdata/example.json",
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			want, err := os.ReadFile(tt.path)
			if !assert.NoError(t, err) {
				t.FailNow()
			}

			_, err = ParseGrafana(want)
			if !assert.NoError(t, err) {
				t.FailNow()
			}
		})
	}
}
