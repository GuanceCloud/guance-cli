package dashboard

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"html/template"

	"github.com/GuanceCloud/guance-cli/internal/helpers/osfs"
	"github.com/GuanceCloud/guance-cli/internal/helpers/prettier"
)

var (
	//go:embed template/main.tf.gotpl
	moduleTemplateMain string
)

type Manifest struct {
	Name    string
	Title   string
	Content json.RawMessage
}

type Options struct {
	Manifests []Manifest
}

func Generate(opts Options) (osfs.Files, error) {
	var files osfs.Files
	for _, manifest := range opts.Manifests {
		fixedManifest, err := prettier.FormatJSON(manifest.Content)
		if err != nil {
			return nil, fmt.Errorf("failed to format manifest: %w", err)
		}
		files = append(files, osfs.File{
			Path:    fmt.Sprintf("dashboards/%s.json", manifest.Name),
			Content: fixedManifest,
		})
	}

	t, err := template.New("terraform-template").Parse(moduleTemplateMain)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, opts.Manifests); err != nil {
		return nil, fmt.Errorf("failed to execute template: %w", err)
	}
	files = append(files, osfs.File{Path: "main.tf", Content: buf.Bytes()})
	return files, nil
}
