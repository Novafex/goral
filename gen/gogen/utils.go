package gogen

import (
	"bytes"
	"strconv"
	"time"
)

// SCT (Source Code Template) for the start of a comment
var sctComment = []byte("// ")

var ThisYear = strconv.FormatInt(int64(time.Now().Year()), 10)

type Number interface {
	byte | int | int8 | int16 | int32 | int64 |
		uint | uint16 | uint32 | uint64 |
		float32 | float64
}

// Min takes the smaller of the two numbers
func Min[T Number](lhs, rhs T) T {
	if lhs <= rhs {
		return lhs
	}
	return rhs
}

// Max takes the larger of the two numbers
func Max[T Number](lhs, rhs T) T {
	if lhs >= rhs {
		return lhs
	}
	return rhs
}

// BreakLines takes a block of text as bytes and breaks them up into sub byte
// arrays in order to fit in a 80-column layout. The offset (starting padding)
// is taken into account for each line.
func BreakLines(text []byte, offset int) [][]byte {
	ln := len(text)
	width := 80 - offset
	num := Max(ln/width, 1)
	lines := make([][]byte, num)
	start := 0
	end := ln
	for i := 0; i < num; i++ {
		start = i * width
		end = Min((i+1)*width, ln)
		lines[i] = text[start:end]
	}
	return lines
}

// WriteNL adds a new-line character to the buffer
func WriteNL(buf *bytes.Buffer) {
	buf.WriteByte('\n')
}

// WriteLines takes a slice of byte-arrays and treats each one as a line of text.
func WriteLines(buf *bytes.Buffer, content [][]byte) {
	for _, line := range content {
		buf.Write(line)
		buf.WriteByte('\n')
	}
}

func WriteComment(buf *bytes.Buffer, content []byte, indent int) {
	for _, line := range BreakLines(content, indent+3) {
		buf.Write(sctComment)
		buf.Write(line)
		buf.WriteByte('\n')
	}
}
