package osfs

import (
	"bytes"
	"testing"
)

func TestRenderTemplate(t *testing.T) {
	templateContent := []byte("Hello, {{.Name}}!")
	data := struct{ Name string }{"World"}
	expectedOutput := []byte("Hello, World!")

	output, err := RenderTemplate(templateContent, data)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !bytes.Equal(output, expectedOutput) {
		t.Errorf("unexpected output: %s", output)
	}
}
