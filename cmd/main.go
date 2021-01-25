package main

import (
	"bufio"
	"os"
)

// write object pointer
var writer *bufio.Writer

func main() {
	writer = bufio.NewWriter(os.Stdout)
	s := "How many stars does Orion have?\n"
	var b byte = 'H'

	writer.WriteString(s)
	writer.WriteByte(b)
	writer.WriteString("\n")

	defer writer.Flush()
}
