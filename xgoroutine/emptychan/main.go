package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println("hello, yufei")
	_, ok := <-ch
	if !ok {
		fmt.Println("not ok")
	}
}
