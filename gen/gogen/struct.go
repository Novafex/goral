package gogen

import "bytes"

type Struct struct {
	Comment string
	Name    string

	fields []*Field
}

func (s Struct) Write(buf *bytes.Buffer, indent int) error {
	if len(s.Comment) > 0 {
		WriteComment(buf, []byte(s.Comment), indent)
	}

	buf.Write([]byte("type "))
	buf.WriteString(s.Name)
	buf.Write([]byte(" struct {"))
	buf.WriteByte('\n')

	flen := len(s.fields)
	for i, f := range s.fields {
		f.Write(buf, 4)

		if i < flen-1 {
			buf.WriteByte('\n')
		}
	}
	buf.WriteByte('}')

	return nil
}

func (s *Struct) AddField(f *Field) {
	if s.fields == nil {
		s.fields = make([]*Field, 0)
	}
	s.fields = append(s.fields, f)
}
