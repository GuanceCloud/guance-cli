package dashboard

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/GuanceCloud/guance-cli/internal/helpers/osfs"
	"github.com/GuanceCloud/guance-cli/internal/helpers/prettier"
)

var (
	//go:embed template/main.tf
	moduleTemplateMain []byte

	//go:embed template/outputs.tf
	moduleTemplateOutput []byte

	//go:embed template/variables.tf
	moduleTemplateVar []byte

	//go:embed template/versions.tf
	moduleTemplateVersion []byte
)

type Options struct {
	Manifest json.RawMessage
}

func Generate(opts Options) (osfs.Files, error) {
	fixedManifest, err := prettier.FormatJSON(opts.Manifest)
	if err != nil {
		return nil, fmt.Errorf("failed to format manifest: %w", err)
	}
	var files osfs.Files
	for name, content := range map[string][]byte{
		"versions.tf":   moduleTemplateVersion,
		"manifest.json": fixedManifest,
		"main.tf":       moduleTemplateMain,
		"variables.tf":  moduleTemplateVar,
		"outputs.tf":    moduleTemplateOutput,
	} {
		files = append(files, osfs.File{
			Path:    name,
			Content: content,
		})
	}
	return files, nil
}
