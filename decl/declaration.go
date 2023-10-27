// package decl handles "declarations" which are the key component of Goral.
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
package decl

// Declaration holds the root node and subsequent definitions for a single
// resource to be generated.
type Declaration struct {
	// Name is the proper written name of the resource. It can include spaces
	// and should be title-cased. Generators will handle formatting it for
	// generation.
	Name string `yaml:"name" toml:"name" json:"name"`

	// Description holds a short text explaining what the resource represents.
	// This is used in documentation snippets to help describe the object.
	Description string `yaml:"description" toml:"description" json:"description"`

	// Properties holds the sub-data structure for this resource and all the
	// data that is a child of this resource.
	Properties []Property `yaml:"properties" toml:"properties" json:"properties"`
}
