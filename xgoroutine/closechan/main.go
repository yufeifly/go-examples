package main

import "fmt"

func main() {
	fmt.Println("Begin doing something!")
	c := make(chan bool)
	go func() {
		fmt.Println("Doing something…")
		close(c)
	}()
	<-c
	fmt.Println("Done!")
}
