package mksb

import (
	"bytes"
	"fmt"
	"strings"
)

func WrapSB(sbName string, lines [][]byte, delimiter int) *strings.Builder {
	sb := &strings.Builder{}
	numLines := len(lines)
	for i, line := range lines {
		sb.WriteString(fmt.Sprintf("%s.WriteString(\"", sbName))
		sb.Write(bytes.Replace(line, []byte{34}, []byte{92, 34}, -1))
		if i != numLines-1 {
			sb.WriteByte(byte(delimiter))
		}
		sb.WriteString("\")\n")
	}

	return sb
}
