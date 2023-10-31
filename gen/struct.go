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

	// Prepare struct
	s := &gogen.Struct{
		Comment: dcl.Description,
		Name:    utils.ToPascalCase(dcl.Name),
	}

	// Prepare properties
	var newField *gogen.Field
	for _, prop := range dcl.Properties {
		newField = &gogen.Field{
			Comment:  prop.Description,
			Name:     utils.ToPascalCase(prop.Name),
			Type:     prop.Type.ToGo(),
			Optional: prop.Optional,
		}

		newField.AddTag("json", utils.ToCamelCase(prop.Name))
		newField.AddTag("sql", utils.ToSnakeCase(prop.Name))

		s.AddField(newField)
	}

	file.AddSection(s)

	if err := file.Generate(); err != nil {
		return nil, err
	}
	return file.Bytes(), nil
}
