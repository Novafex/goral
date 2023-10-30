package gogen

import "bytes"

type BufferWritable interface {
	Write(buf *bytes.Buffer, indent int) error
}
