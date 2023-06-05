package grafana

import (
	_ "embed"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	for _, tt := range []struct {
		name    string
		path    string
		wantErr bool
	}{
		//{
		//	name: "example",
		//	path: "testdata/example.json",
		//},
		{
			name: "nginx",
			path: "testdata/nginx.json",
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

			//got, err := json.Marshal(dashboard)
			//if !assert.NoError(t, err) {
			//	t.FailNow()
			//}
			//assert.Equal(t, string(want), string(got))
		})
	}

}
