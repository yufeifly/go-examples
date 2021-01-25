package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	fmt.Fprint(&b, "hello,world")
	fmt.Println(b.String())

	b.WriteByte('\n')
	b.WriteByte('x')
	b.WriteByte('y')
	b.WriteByte('f')
	fmt.Println(b.String())
}
