package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/yufeifly/go-examples/goprof/data"
)

func main() {
	go func() {
		for {
			log.Println(data.Add("https://github.com/yufeifly"))
			time.Sleep(10 * time.Millisecond)
		}
	}()
	http.ListenAndServe("0.0.0.0:6060", nil)
}
