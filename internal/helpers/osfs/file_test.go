package osfs

import (
	"os"
	"path"
	"testing"
)

func TestFiles_Save(t *testing.T) {
	dir, err := os.MkdirTemp("", "test")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	files := Files{
		{
			Path:    "file1.txt",
			Content: []byte("hello world"),
		},
		{
			Path:    "file2.txt",
			Content: []byte("goodbye world"),
		},
	}

	if err := files.Save(dir); err != nil {
		t.Fatalf("failed to save files: %v", err)
	}

	for _, file := range files {
		filePath := path.Join(dir, file.Path)
		content, err := os.ReadFile(filePath)
		if err != nil {
			t.Fatalf("failed to read file %s: %v", filePath, err)
		}
		if string(content) != string(file.Content) {
			t.Errorf("file content mismatch, expected %q but got %q", string(file.Content), string(content))
		}
	}
}
