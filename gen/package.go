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

import "github.com/novafex/goral/utils"

func ProcessDeclarationFile(path string) error {
	return nil
}

func ProcessDeclarationFiles(paths []string) []error {
	return nil
}

// ParallelProcessDeclarationFiles takes an incoming list of absolute paths and
// runs the [ProcessDeclarationFile] on each one using a go-routine to make it
// parallel. Any errors are collected and returned and a slice with indices
// matching the arguments.
func ParallelProcessDeclarationFiles(paths []string) []error {
	return utils.ParallelizeArgsWithError(ProcessDeclarationFile, paths)
}
