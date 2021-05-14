package main

import (
	"fmt"
	"time"
)

func main() {
	a := time.After(2 * time.Second)
	<-a
	fmt.Println("timer receive")

	ch := make(chan int)
	time.AfterFunc(4*time.Second, func() {
		fmt.Println("timer receive")
		ch <- 100
	})
	<-ch
}
