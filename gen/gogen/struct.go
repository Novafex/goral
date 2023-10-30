package gogen

import "bytes"

type Struct struct {
	Comment string
	Name    string
}

func (s Struct) Write(buf *bytes.Buffer, indent int) error {
	if len(s.Comment) > 0 {
		WriteComment(buf, []byte(s.Comment), indent)
	}

	buf.Write([]byte("type "))
	buf.WriteString(s.Name)
	buf.Write([]byte(" struct {"))
	buf.WriteByte('\n')

	buf.WriteByte('\n')
	buf.WriteByte('}')
	return nil
}
