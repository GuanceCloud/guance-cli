//go:build mage
// +build mage

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Run mg.Namespace

const (
	GaugeSpecDir    = "specs"
	GaugeProjectDir = "internal/testing"
)

// Install run the installation in local
func (ns Run) Install() error {
	return sh.RunV("go", "install", "./cmd/...")
}

// Test run all test specifications
func (ns Run) Test() error {
	mg.Deps(ns.Install)

	// Copy the specs to the project directory
	err := filepath.Walk(GaugeSpecDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".spec.md") {
			return nil
		}

		// Copy the file to the project directory
		srcBytes, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read file %s failed: %w", path, err)
		}
		relPath, err := filepath.Rel(GaugeSpecDir, path)
		if err != nil {
			return fmt.Errorf("get relative path failed: %w", err)
		}
		dstPath := filepath.Join(GaugeProjectDir, "specs", relPath[:len(relPath)-3])
		if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
			return fmt.Errorf("make dir %s failed: %w", filepath.Dir(dstPath), err)
		}
		if err := os.WriteFile(dstPath, srcBytes, 0o644); err != nil {
			return fmt.Errorf("write file %s failed: %w", dstPath, err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("walk specs failed: %w", err)
	}

	// Run gauge
	if err := sh.RunV("gauge", "run", filepath.Join(GaugeProjectDir, "specs")); err != nil {
		return fmt.Errorf("run gauge failed: %w", err)
	}
	return nil
}
