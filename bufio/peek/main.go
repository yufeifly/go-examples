package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s1 := strings.NewReader(strings.Repeat("a", 20))
	r := bufio.NewReaderSize(s1, 16)
	b, err := r.Peek(3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}
