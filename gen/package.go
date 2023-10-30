// package gen handles generating the Go code based on declarations
//
// ============================================================================
// Copyright © 2023 Novafex
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package gen

import (
	"os"
	"path/filepath"

	"github.com/novafex/goral/decl"
	"github.com/novafex/goral/fs"
	"github.com/novafex/goral/utils"
)

func ProcessDeclarationFile(path string) error {
	dcl := &decl.Declaration{}
	if err := fs.Read(path, dcl); err != nil {
		return err
	}

	data, err := GenerateStructFile(dcl)
	if err != nil {
		return err
	}

	dir := filepath.Dir(path)
	filename := utils.ToKebabCase(dcl.Name)
	finalPath := filepath.Join(dir, filename+".go")
	if err := os.WriteFile(finalPath, data, 0644); err != nil {
		return err
	}

	return nil
}

func ProcessDeclarationFiles(paths []string) []error {
	errs := make([]error, 0)

	for _, path := range paths {
		if err := ProcessDeclarationFile(path); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}

// ParallelProcessDeclarationFiles takes an incoming list of absolute paths and
// runs the [ProcessDeclarationFile] on each one using a go-routine to make it
// parallel. Any errors are collected and returned and a slice with indices
// matching the arguments.
func ParallelProcessDeclarationFiles(paths []string) []error {
	return utils.ParallelizeArgsWithError(ProcessDeclarationFile, paths)
}
