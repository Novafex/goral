// package fs gives utilities and other functionality for finding and loading
// declarations, configurations, and other Goral related files.
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
package fs

import "os"

var (
	extensionOrder = []string{
		"yaml",
		"toml",
		"json",
	}
)

// GetExtensionOrder returns the static array of extensions (without dot) that
// should be used to resolve files.
//
// These are decided statically so that all systems use the same order.
func GetExtensionOrder() []string {
	return extensionOrder
}

// DirExists returns ture if the given path exists and is a directory.
//
// See [FileExists] to check for files.
func DirExists(path string) bool {
	stat, err := os.Stat(path)
	return !os.IsNotExist(err) && stat.IsDir()
}

// FileExists returns true if the given path exists and is a file.
//
// See [DirExists] to check for directories.
func FileExists(path string) bool {
	stat, err := os.Stat(path)
	return !os.IsNotExist(err) && !stat.IsDir()
}

// CombineBaseExt is a convienience function to join a base path and the
// given extension with a dot in the middle.
func CombineBaseExt(base, ext string) string {
	return base + "." + ext
}

// FindPathWithExtensions takes a base path to a file (without extension) and
// searches for one that matches using the [GetExtensionOrder] array of accepted
// formats. It then returns a boolean of search success, and the extension that
// was found. If no extension was found, it defaults to the first wanted in
// case you want to create it after.
func FindPathWithExtensions(base string) (bool, string) {
	for i := range extensionOrder {
		if FileExists(CombineBaseExt(base, extensionOrder[i])) {
			return true, extensionOrder[i]
		}
	}
	return false, extensionOrder[0]
}
