package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b bytes.Buffer
	fmt.Fprint(&b, "hello,world")
	fmt.Println(b.String())
}
