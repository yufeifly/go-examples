package main

import (
	"io"
	"os"
	"strings"

	"github.com/pierrec/lz4"
)

func main() {

	// Compress and uncompress an input string.
	s := "hello world"
	r := strings.NewReader(s)

	// The pipe will uncompress the data from the writer.
	pr, pw := io.Pipe()
	zw := lz4.NewWriter(pw)
	zr := lz4.NewReader(pr)

	go func() {
		// Compress the input string.
		_, _ = io.Copy(zw, r)
		_ = zw.Close() // Make sure the writer is closed
		_ = pw.Close() // Terminate the pipe
	}()

	_, _ = io.Copy(os.Stdout, zr)

	// Output:
	// hello world
}
