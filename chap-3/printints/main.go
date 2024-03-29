package main

import (
	"bytes"
	"fmt"
)

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')

	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		_, _ = fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	fmt.Println(buf)
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3}))
}
