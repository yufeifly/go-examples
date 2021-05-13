package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := strings.NewReader("abcdef|ghij")
	r := bufio.NewReaderSize(s, 4)
	token, err := r.ReadSlice('z')
	if err != nil {
		panic(err)

	}
	fmt.Printf("token: %q\n", token)
}
