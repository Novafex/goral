package decl

type Property interface {
	Base() *BaseProperty
}

// Property describes any field that is a part of a resource declaration. It is
// called "Property" because Field is already common in Go reflect, so I didn't
// want to confuse code. Also, JS generally calls them properties.
//
// Properties only define data related directly to them. Within a declaration,
// the Properties are kept as a list instead of a dictionary so that the ID, or
// name is not separated from the data structure.
//
// Note that this is sort of a base type or interface because depending on the
// type, it can declare other fields specific to that type.
type BaseProperty struct {
	// Name is the title-case displayable name for the property.
	//
	// These are title case (with spaces, properly named) so that generators can
	// use the proper identifier format depending on the target language.
	//
	// Ex: "First Customer ID" becomes "firstCustomerID" in JSON.
	Name string `yaml:"name" toml:"name" json:"name"`

	// Description holds a short text explaining what the property represents.
	// This is used in documentation snippets to help describe the property.
	Description string `yaml:"description" toml:"description" json:"description"`

	// Type is a masked enumeration for valid property types. It is serialized
	// as a string. See []
	Type PropertyType `yaml:"type" toml:"type" json:"type"`

	// Optional declares the property optional, as in "not required". In Go terms
	// this makes the Field a pointer that can and defaults to nil.
	Optional bool `yaml:"optional" toml:"optional" json:"optional"`
}

func (pt *BaseProperty) Base() *BaseProperty {
	return pt
}

type BooleanProperty struct {
	*BaseProperty

	Default bool `yaml:"default" toml:"default" json:"default"`
}

type ByteProperty struct {
	*BaseProperty

	Default byte `yaml:"default" toml:"default" json:"default"`
}

type IntegerProperty struct {
	*BaseProperty

	Default int `yaml:"default" toml:"default" json:"default"`
}

type FloatProperty struct {
	*BaseProperty

	Default float32 `yaml:"default" toml:"default" json:"default"`
}

type StringProperty struct {
	*BaseProperty

	Default string `yaml:"default" toml:"default" json:"default"`
}

type ObjectProperty struct {
	*BaseProperty

	Default any `yaml:"default" toml:"default" json:"default"`
}

type ReferenceProperty struct {
	*BaseProperty
}
