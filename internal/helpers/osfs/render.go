package osfs

import (
	"bytes"
	"fmt"
	"html/template"
)

func RenderTemplate(templateContent []byte, data any) ([]byte, error) {
	tpl, err := template.New("guance").Parse(string(templateContent))
	if err != nil {
		return nil, fmt.Errorf("parse template error: %w", err)
	}
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, data); err != nil {
		return nil, fmt.Errorf("execute template error: %w", err)
	}
	return buf.Bytes(), nil
}
