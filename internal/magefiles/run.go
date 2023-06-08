//go:build mage
// +build mage

package main

import (
	"github.com/hashicorp/go-multierror"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

type Run mg.Namespace

// Install run the installation in local
func (ns Run) Install() error {
	return sh.RunV("go", "install", "./cmd/...")
}

// Lint run the linter
func (ns Run) Lint() error {
	argList := [][]string{
		{"golangci-lint", "run", "./..."},
		{"markdownlint", "-i", "docs/references", "-f", "."},
		{"gofumpt", "-l", "-e", "."},
	}
	return batchRunV(argList)
}

// Fmt run the formatter
func (ns Run) Fmt() error {
	argList := [][]string{
		{"golangci-lint", "run", "--fix", "./..."},
		{"gofumpt", "-l", "-w", "."},
		{"goimports", "-w", "."},
		{"prettier", "-w", "**/*.md"},
	}
	return batchRunV(argList)
}

func batchRunV(argList [][]string) error {
	var mErr error
	for _, args := range argList {
		if err := sh.RunV(args[0], args[1:]...); err != nil {
			mErr = multierror.Append(mErr, err)
			continue
		}
	}
	return mErr
}
