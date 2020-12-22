package main

import (
	"bufio"
	"fmt"
)

type R struct {
}

func (r *R) Read(p []byte) (n int, err error) {
	fmt.Println("Read")
	copy(p, "abcdefghijklmnop")
	return 16, nil
}

func main() {
	r := new(R)
	br := bufio.NewReaderSize(r, 16)
	buf := make([]byte, 4)
	br.Read(buf)
	fmt.Println(string(buf))
	br.Discard(4)
	br.Read(buf)
	fmt.Println(string(buf))
}
