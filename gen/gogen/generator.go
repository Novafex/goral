// package gogen is for generating Go code. It is specialized at Goral's way of
// making code.
package gogen

import (
	"bytes"
	"errors"
)

// GoGenerator is a struct holding the buffer for a Go source code file.
type GoGenerator struct {
	buf *bytes.Buffer

	// sections holds independent parts of the source code that can write
	// themselves to the buffer.
	sections []BufferWritable

	// Package is the required package name at the top of the file
	Package string

	// PackageDescription is the overall package description that is placed at
	// the top of the file. It DOES NOT need the package prefix.
	PackageDescription []byte

	// Author holds the name of the author for the copyright. If not provided,
	// it will be excluded.
	Author []byte

	// License holds the content of the license text. It will automatically be
	// broken up into the related comments by recommended 80-char length.
	License []byte
}

// Bytes returns the contents of the byte buffer. If one has not been created
// yet it will be made using [Generate].
func (gen GoGenerator) Bytes() []byte {
	if gen.buf == nil {
		panic("GoGenerator has not been generated")
	}
	return gen.buf.Bytes()
}

// String returns the contents of the byte buffer. If one has not been created
// yet it will be made using [Generate].
func (gen GoGenerator) String() string {
	if gen.buf == nil {
		panic("GoGenerator has not been generated")
	}
	return gen.buf.String()
}

// MarshalText implements the text marshaler for writing out the Go source code.
// It uses the internal buffer, so if one has not been created yet it will be
// generated when calling this method using [Generate]
func (gen GoGenerator) MarshalText() ([]byte, error) {
	if gen.buf == nil {
		return nil, errors.New("GoGenerator has not been generated")
	}
	return gen.Bytes(), nil
}

// Generate is the magic function which builds the internal buffer of the source
// code. It will clear the current buffer if there is one.
func (gen GoGenerator) Generate() error {
	if gen.buf != nil {
		gen.buf.Reset()
	}

	// Write package description if there is one
	if len(gen.PackageDescription) > 0 {
		WriteComment(gen.buf, gen.PackageDescription, 0)

		// If there will be license, pad it by two lines
		if len(gen.License) > 0 {
			gen.buf.Write(sctComment)
			gen.buf.WriteByte('\n')
			gen.buf.Write(sctComment)
			gen.buf.WriteByte('\n')
		}
	}

	// Write author
	if len(gen.Author) > 0 {
		gen.buf.WriteString("// © ")
		gen.buf.WriteString(ThisYear)
		gen.buf.WriteByte(' ')
		gen.buf.Write(gen.Author)
		gen.buf.WriteByte('\n')
		gen.buf.Write(sctComment)
		gen.buf.WriteByte('\n')
	}

	// Write license
	if len(gen.License) > 0 {
		WriteComment(gen.buf, gen.License, 0)
	}

	// Write package
	gen.buf.WriteString("package ")
	gen.buf.WriteString(gen.Package)
	gen.buf.WriteByte('\n')
	gen.buf.WriteByte('\n')

	// Throw in all the sections
	for _, sect := range gen.sections {
		if sect != nil {
			if err := sect.Write(gen.buf, 0); err != nil {
				return err
			}
			gen.buf.WriteByte('\n')
		}
	}

	return nil
}

func (gen *GoGenerator) AddSection(sect BufferWritable) {
	gen.sections = append(gen.sections, sect)
}

func NewGoGenerator(pkg string) *GoGenerator {
	return &GoGenerator{
		buf:      new(bytes.Buffer),
		sections: make([]BufferWritable, 0),
		Package:  pkg,
	}
}
