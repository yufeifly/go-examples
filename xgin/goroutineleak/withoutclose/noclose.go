package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for true {
		requestWithoutClose()
		time.Sleep(time.Microsecond * 10)
	}
}

func requestWithoutClose() {
	_, err := http.Get("http://127.0.0.1:8080/user/name")
	if err != nil {
		fmt.Printf("error occurred while fetching page, error: %s", err.Error())
		return
	}
	//fmt.Println("ok")
}
