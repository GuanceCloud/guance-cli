//go:build mage
// +build mage

package main

import (
	"os"
	"path"

	"github.com/GuanceCloud/guance-cli/internal/cmd"

	"github.com/magefile/mage/mg"
	"github.com/spf13/cobra/doc"
)

const (
	DocsRootDir = "docs/references"
)

// Gen generate code from RMS definitions
type Gen mg.Namespace

// Doc generate documentation
func (g Gen) Doc() error {
	rootCmd := cmd.NewRootCmd()

	// generate man page
	manPath := path.Join(DocsRootDir, "man")
	if err := os.MkdirAll(manPath, 0o755); err != nil {
		return err
	}
	header := &doc.GenManHeader{}
	err := doc.GenManTree(rootCmd, header, manPath)
	if err != nil {
		return err
	}

	// generate markdown
	err = doc.GenMarkdownTree(rootCmd, DocsRootDir)
	if err != nil {
		return err
	}
	return nil
}
