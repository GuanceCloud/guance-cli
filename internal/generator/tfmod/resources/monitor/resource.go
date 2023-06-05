package monitor

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/GuanceCloud/guance-cli/internal/helpers/prettier"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"

	"github.com/hashicorp/go-multierror"

	"github.com/GuanceCloud/guance-cli/internal/helpers/osfs"
)

var (
	//go:embed template/main.tf.gotpl
	moduleTemplateMain []byte

	//go:embed template/outputs.tf.gotpl
	moduleTemplateOutput []byte

	//go:embed template/variables.tf
	moduleTemplateVar []byte

	//go:embed template/versions.tf
	moduleTemplateVersion []byte
)

type Options struct {
	Manifests []json.RawMessage
}

func Generate(opts Options) (osfs.Files, error) {
	var files osfs.Files

	// copy template
	for name, content := range map[string]json.RawMessage{
		"versions.tf":  moduleTemplateVersion,
		"variables.tf": moduleTemplateVar,
	} {
		files = append(files, osfs.File{
			Path:    name,
			Content: content,
		})
	}

	// render template
	var mErr error
	for name, content := range map[string]json.RawMessage{
		"main.tf":    moduleTemplateMain,
		"outputs.tf": moduleTemplateOutput,
	} {
		got, err := osfs.RenderTemplate(content, opts)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to render template %s: %w", name, err))
		}
		files = append(files, osfs.File{
			Path:    name,
			Content: got,
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to render template: %w", mErr)
	}

	// add manifest
	for i, manifest := range opts.Manifests {
		fixed, err := fixNoData(manifest)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to fix manifest: %w", err))
		}
		fixed, err = prettier.FormatJSON(fixed)
		if err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("failed to format manifest: %w", err))
		}
		files = append(files, osfs.File{
			Path:    fmt.Sprintf("manifest-%03d.json", i+1),
			Content: fixed,
		})
	}
	if mErr != nil {
		return nil, fmt.Errorf("failed to fix data zero-value: %w", mErr)
	}
	return nil, nil
}

func fixNoData(src []byte) ([]byte, error) {
	var err error
	if gjson.GetBytes(src, "jsonScript.noDataPeriodCount").Int() == 0 {
		src, err = sjson.DeleteBytes(src, "jsonScript.noDataPeriodCount")
		if err != nil {
			return nil, fmt.Errorf("delete noDataPeriodCount error: %w", err)
		}
	}
	if gjson.GetBytes(src, "extend.noDataPeriodCount").Int() == 0 {
		src, err = sjson.DeleteBytes(src, "extend.noDataPeriodCount")
		if err != nil {
			return nil, fmt.Errorf("delete noDataPeriodCount error: %w", err)
		}
	}
	if gjson.GetBytes(src, "jsonScript.noDataMessage").String() == "" {
		src, err = sjson.DeleteBytes(src, "jsonScript.noDataMessage")
		if err != nil {
			return nil, fmt.Errorf("delete noDataMessage error: %w", err)
		}
	}
	if gjson.GetBytes(src, "jsonScript.noDataTitle").String() == "" {
		src, err = sjson.DeleteBytes(src, "jsonScript.noDataTitle")
		if err != nil {
			return nil, fmt.Errorf("delete noDataTitle error: %w", err)
		}
	}
	if !gjson.GetBytes(src, "jsonScript.checkerOpt.infoEvent").Bool() {
		src, err = sjson.DeleteBytes(src, "jsonScript.checkerOpt.infoEvent")
		if err != nil {
			return nil, fmt.Errorf("delete infoEvent error: %w", err)
		}
	}
	return src, nil
}
