package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("AfterFunc")

		t := time.NewTicker(time.Second * 2)
		defer t.Stop()
		ch := make(chan int)
		go func() {
			for now := range t.C {
				fmt.Println(now, " Ticker")
				ch <- 100
			}
		}()
		for {
			<-ch
		}
	})
	// Block until termination signal is received
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-osSignals

	fmt.Println("exiting gracefully")
}
