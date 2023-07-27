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

type Test mg.Namespace

const (
	GaugeSpecDir    = "specs"
	GaugeProjectDir = "internal/testing"
)

// Unit run all unit test and generate the coverage report
func (ns Test) Unit() error {
	if err := sh.RunV("go", "test", "-v", "-cover", "-coverprofile=cover.out", "./..."); err != nil {
		return fmt.Errorf("run unit test failed: %w", err)
	}
	if err := sh.RunV("go", "tool", "cover", "-func=cover.out"); err != nil {
		return fmt.Errorf("generate coverage report failed: %w", err)
	}
	return nil
}

// Acc run all test specifications for acceptance testing
func (ns Test) Acc() error {
	// Copy the specs to the project directory
	err := filepath.Walk(GaugeSpecDir, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.Name() == "README.md" {
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
		if strings.HasSuffix(path, ".spec.md") {
			relPath = strings.ReplaceAll(relPath, ".spec.md", ".spec")
		}
		dstPath := filepath.Join(GaugeProjectDir, "specs", relPath)
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

	// Get the root of Gauge files
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("cannot get current directory: %w", err)
	}
	gaugeRoot, err := filepath.Abs(filepath.Join(cwd, "specs"))
	if err != nil {
		return fmt.Errorf("cannot get gauge root: %w", err)
	}

	// Run gauge
	if err := sh.RunWithV(map[string]string{
		"GAUGE_ROOT": gaugeRoot,
	}, "gauge", "run", filepath.Join(GaugeProjectDir, "specs")); err != nil {
		return fmt.Errorf("run gauge failed: %w", err)
	}
	return nil
}
