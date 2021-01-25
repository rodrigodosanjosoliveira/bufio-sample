package main

import (
	"bufio"
	"os"
)

// writer object pointer
var writer *bufio.Writer

func main() {
	// All output will be redirected to
	// bufio.NewWriter
	writer = bufio.NewWriter(os.Stdout)
	s := "How many stars does Orion have?\n"
	var b byte = 'H'

	writer.WriteString(s)
	writer.WriteByte(b)
	writer.WriteString("\n")

	// close buffer
	defer writer.Flush()
}
