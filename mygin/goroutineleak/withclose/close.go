package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for true {
		requestWithClose()
		time.Sleep(time.Microsecond * 10)
	}
}

func requestWithClose() {
	resp, err := http.Get("http://127.0.0.1:8080/user/name")
	if err != nil {
		fmt.Printf("error occurred while fetching page, error: %s", err.Error())
		return
	}
	defer resp.Body.Close()
	//fmt.Println("ok")
}
