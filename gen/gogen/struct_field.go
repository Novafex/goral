package gogen

import (
	"bytes"
	"strings"
)

type Field struct {
	Comment  string
	Name     string
	Type     string
	Optional bool
	Tags     map[string][]string
}

func (f Field) Write(buf *bytes.Buffer, indent int) error {
	if len(f.Comment) > 0 {
		WriteIndent(buf, indent)
		WriteComment(buf, []byte(f.Comment), indent)
	}

	// Write the name
	WriteIndent(buf, indent)
	buf.WriteString(f.Name)
	buf.WriteByte(' ')

	// Write the type
	if f.Optional {
		buf.WriteByte('*')
	}
	buf.WriteString(f.Type)

	// Add the tags
	if len(f.Tags) > 0 {
		buf.WriteByte(' ')
		buf.WriteByte('`')

		first := true
		for g, t := range f.Tags {
			if !first {
				buf.WriteByte(' ')
			}
			first = false

			// Write the group
			buf.WriteString(g)
			buf.WriteByte(':')
			buf.WriteByte('"')

			// Write tags
			buf.WriteString(strings.Join(t, ","))

			// Close group
			buf.WriteByte('"')
		}

		buf.WriteByte('`')
	}

	WriteNL(buf)

	return nil
}

func (f *Field) AddTag(group, content string) {
	if f.Tags == nil {
		f.Tags = make(map[string][]string)
	}

	f.Tags[group] = append(f.Tags[group], content)
}
