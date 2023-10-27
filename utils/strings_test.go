package utils

import "testing"

func TestRemoveNonAlphanumeric(t *testing.T) {
	cases := map[string]string{
		"foobar":    "foobar",
		"foo123":    "foo123",
		"foo bar":   "foobar",
		"F1zz1e":    "F1zz1e",
		"Foo's Bar": "FoosBar",
	}
	for k, v := range cases {
		if RemoveNonAlphanumeric(k) != v {
			t.Errorf("expected %s to equal %s", k, v)
		}
	}
}

func TestToKebabCase(t *testing.T) {
	cases := map[string]string{
		"Foo's Bar":   "foos-bar",
		"Foo 123":     "foo-123",
		"Bar Baz 12e": "bar-baz-12e",
	}
	for k, v := range cases {
		if ToKebabCase(k) != v {
			t.Errorf("expected %s to equal %s", k, v)
		}
	}
}
