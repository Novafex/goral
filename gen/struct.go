package gen

import (
	"github.com/novafex/goral/decl"
	"github.com/novafex/goral/gen/gogen"
	"github.com/novafex/goral/utils"
)

// GenerateStructFile is responsible for building the text that represents a struct
// in source code form. It takes an incoming [decl.Declaration] and returns the
// source code file for the struct type and any needed support methods.
//
// Note this is the whole file with the package header and imports.
func GenerateStructFile(dcl *decl.Declaration) ([]byte, error) {
	file := gogen.NewGoGenerator("goral")

	file.PackageDescription = []byte(dcl.Description)
	file.License = []byte("Some dummy license")
	file.Author = []byte("My Name")

	file.AddSection(&gogen.Struct{
		Comment: dcl.Description,
		Name:    utils.ToPascalCase(dcl.Name),
	})

	if err := file.Generate(); err != nil {
		return nil, err
	}
	return file.Bytes(), nil
}
