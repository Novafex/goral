package decl

import (
	"fmt"
	"strings"
)

// PropertyType masks the string property type in an integer so that Go can
// treat it like a enumeration.
type PropertyType uint

const (
	// PropertyTypeInvalid is the default value and marks an invalid type
	PropertyTypeInvalid PropertyType = iota

	// PropertyTypeBoolean is a boolean value
	PropertyTypeBoolean

	// PropertyTypeByte is a byte value 0..255
	PropertyTypeByte

	// PropertyTypeInteger is a whole-number integer
	PropertyTypeInteger

	// PropertyTypeFloat is a fractional number with decimals
	PropertyTypeFloat

	// PropertyTypeString is a string of characters
	PropertyTypeString

	// PropertyTypeObject is reserved for later usage but will most likely be
	// sub-resources
	PropertyTypeObject

	// PropertyTypeReference is reserved for later usage but will most likely be
	// sub-resources
	PropertyTypeReference
)

var PropertyTypes = map[PropertyType]string{
	PropertyTypeInvalid:   "invalid",
	PropertyTypeBoolean:   "boolean",
	PropertyTypeByte:      "byte",
	PropertyTypeInteger:   "integer",
	PropertyTypeFloat:     "float",
	PropertyTypeString:    "string",
	PropertyTypeObject:    "object",
	PropertyTypeReference: "reference",
}

// String returns a string form of the enumeration
func (pt PropertyType) String() string {
	str, ok := PropertyTypes[pt]
	if !ok {
		return PropertyTypes[PropertyTypeInvalid]
	}
	return str
}

// Valid returns true if the PropertyType enum is considered valid
func (pt PropertyType) Valid() bool {
	return pt > PropertyTypeInvalid && pt <= PropertyTypeReference
}

// UnmarshalText implements the text unmarshaler, which is used by most other
// formats such as JSON to unmarshal string fields.
//
// Will return an error if the type is invalid
func (pt *PropertyType) UnmarshalText(src []byte) error {
	*pt = ParsePropertyType(string(src))
	if *pt == PropertyTypeInvalid {
		return fmt.Errorf("invalid property type %s", src)
	}
	return nil
}

// ParsePropertyType reads an incoming string and finds a matching property type
// enumeration. This is case-insensitive brute-force so it may not be the fastest.
//
// Instead of an error, this returns 0 or [PropertyTypeInvalid] if it cannot be
// found.
func ParsePropertyType(str string) PropertyType {
	clean := strings.ToLower(str)
	for key, val := range PropertyTypes {
		if clean == val {
			return key
		}
	}

	return PropertyTypeInvalid
}
