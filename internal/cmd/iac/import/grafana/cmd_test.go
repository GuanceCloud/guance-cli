package grafana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrafana(t *testing.T) {
	cmd := NewCmd()
	cmd.SetArgs([]string{"-f", "testdata/node.json", "-t", "terraform-module", "-o", "testdata/output"})
	if err := cmd.Execute(); err != nil {
		assert.NoError(t, err)
	}
}
