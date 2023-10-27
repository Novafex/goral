package decl

import (
	"math"
	"testing"
)

func TestPropertyTypeString(t *testing.T) {
	if PropertyTypeInvalid.String() != "invalid" || PropertyTypeInvalid.String() != PropertyTypes[PropertyTypeInvalid] {
		t.Error("PropertyTypeInvalid did not string correctly")
	}

	if PropertyTypeBoolean.String() != "boolean" || PropertyTypeBoolean.String() != PropertyTypes[PropertyTypeBoolean] {
		t.Error("PropertyTypeBoolean did not string correctly")
	}

	if PropertyTypeByte.String() != "byte" || PropertyTypeByte.String() != PropertyTypes[PropertyTypeByte] {
		t.Error("PropertyTypeByte did not string correctly")
	}

	if PropertyTypeInteger.String() != "integer" || PropertyTypeInteger.String() != PropertyTypes[PropertyTypeInteger] {
		t.Error("PropertyTypeInteger did not string correctly")
	}

	if PropertyTypeFloat.String() != "float" || PropertyTypeFloat.String() != PropertyTypes[PropertyTypeFloat] {
		t.Error("PropertyTypeFloat did not string correctly")
	}

	if PropertyTypeString.String() != "string" || PropertyTypeString.String() != PropertyTypes[PropertyTypeString] {
		t.Error("PropertyTypeString did not string correctly")
	}

	if PropertyTypeObject.String() != "object" || PropertyTypeObject.String() != PropertyTypes[PropertyTypeObject] {
		t.Error("PropertyTypeObject did not string correctly")
	}

	if PropertyTypeReference.String() != "reference" || PropertyTypeReference.String() != PropertyTypes[PropertyTypeReference] {
		t.Error("PropertyTypeReference did not string correctly")
	}

	if PropertyType(math.MaxInt).String() != "invalid" {
		t.Error("PropertyType.String() allowed invalid value")
	}
}

func TestPropertyTypeValid(t *testing.T) {
	if PropertyTypeInvalid.Valid() {
		t.Error("expected invalid to not be valid")
	}

	for k, v := range PropertyTypes {
		if k != PropertyTypeInvalid && !k.Valid() {
			t.Errorf("expected %s to be valid", v)
		}
	}

	if PropertyType(math.MaxInt).Valid() {
		t.Error("expected max int to be invalid")
	}
}

func TestParsePropertyType(t *testing.T) {
	for k, v := range PropertyTypes {
		if ParsePropertyType(v) != k {
			t.Errorf("expected %s to parse correctly", v)
		}
	}

	if ParsePropertyType("sTrIng") != PropertyTypeString {
		t.Error("expected case-insensitive matching for string")
	}

	if ParsePropertyType("foo") != PropertyTypeInvalid {
		t.Error("allowed invalid string")
	}
}

func TestUnmarshalText(t *testing.T) {
	var pt PropertyType

	for k, v := range PropertyTypes {
		if k == PropertyTypeInvalid {
			continue
		}

		if err := pt.UnmarshalText([]byte(v)); err != nil {
			t.Error(err)
		} else if pt != k {
			t.Errorf("expected unmarshaling %s to match", v)
		}
	}

	if err := pt.UnmarshalText([]byte(PropertyTypes[PropertyTypeInvalid])); err == nil {
		t.Error("expected invalid to produce error")
	}

	if err := pt.UnmarshalText([]byte("foo")); err == nil {
		t.Error("expected garbage to produce error")
	}
}
