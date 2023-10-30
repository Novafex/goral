package utils

import (
	"strings"
	"unicode"
)

// RemoveNonAlphanumeric removes any characters in a string that are not alpha-
// numeric. That being anything that is not considered a unicode letter or digit.
func RemoveNonAlphanumeric(str string) string {
	bld := strings.Builder{}
	for _, c := range str {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			bld.WriteRune(c)
		}
	}
	return bld.String()
}

func ToTitleCase(str string) string {
	bld := strings.Builder{}
	bld.WriteRune(unicode.ToTitle(rune(str[0])))
	bld.WriteString(strings.ToLower(str[1:]))
	return bld.String()
}

// ToKebabCase takes a string and breaks it up by spaces and rejoins as a alpha-
// numeric string joined by dashes.
//
// "Foo's Bar" => "foos-bar"
func ToKebabCase(str string) string {
	parts := strings.Split(strings.ToLower(str), " ")
	for i, p := range parts {
		parts[i] = RemoveNonAlphanumeric(p)
	}
	return strings.Join(parts, "-")
}

func ToSnakeCase(str string) string {
	parts := strings.Split(strings.ToLower(str), " ")
	for i, p := range parts {
		parts[i] = RemoveNonAlphanumeric(p)
	}
	return strings.Join(parts, "_")
}

func ToPascalCase(str string) string {
	parts := strings.Split(strings.ToLower(str), " ")
	for i, p := range parts {
		parts[i] = ToTitleCase(RemoveNonAlphanumeric(p))
	}
	return strings.Join(parts, "")
}

func ToCamelCase(str string) string {
	parts := strings.Split(strings.ToLower(str), " ")
	for i, p := range parts {
		parts[i] = RemoveNonAlphanumeric(p)
		if i > 0 {
			parts[i] = ToTitleCase(parts[i])
		}
	}
	return strings.Join(parts, "_")
}
