package grafana

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// All test code is migrated to User Specification Test
// You can see the `specs/iac/import/grafana` folder for more details.

func TestGrafana(t *testing.T) {
	t.Skip()

	cmd := NewCmd()
	cmd.SetArgs([]string{"-f", "testdata/node.json", "-t", "terraform-module", "-o", "testdata/output"})
	if err := cmd.Execute(); err != nil {
		assert.NoError(t, err)
	}
}
