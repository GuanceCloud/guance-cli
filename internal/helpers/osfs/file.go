package osfs

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/hashicorp/go-multierror"
)

type Files []File

func (files Files) Save(output string) error {
	var mErr error
	for _, file := range files {
		file = File{
			Path:    path.Join(output, file.Path),
			Content: file.Content,
		}
		if err := file.Save(); err != nil {
			mErr = multierror.Append(mErr, fmt.Errorf("save file %s error: %w", file.Path, err))
			continue
		}
	}
	return mErr
}

type File struct {
	Path    string
	Content []byte
}

func (f *File) Save() error {
	if err := os.MkdirAll(filepath.Dir(f.Path), 0o755); err != nil {
		return fmt.Errorf("mkdir %s error: %w", f.Path, err)
	}
	if err := os.WriteFile(f.Path, f.Content, 0o600); err != nil {
		return fmt.Errorf("write file %s error: %w", f.Path, err)
	}
	return nil
}
