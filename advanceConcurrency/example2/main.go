// xiaorui.cc

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctimer()
	fmt.Printf("exit")
	time.Sleep(10 * time.Second)
}

func ctimer() {
	timer := time.NewTimer(5 * time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(2 * time.Second)
		timer.Stop()
		cancel()
		fmt.Println("stop")
	}()

	for {
		select {
		case <-timer.C:
			fmt.Println("ticker.C call")
			return

		case <-ctx.Done():
			return
		}
	}
}
