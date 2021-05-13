package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	url := "http://127.0.0.1:8080/user/yufei"
	clt := http.Client{}
	resp, err := clt.Get(url)
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Printf("response.body: %v\n", string(body))

	url = "http://127.0.0.1:8080/json/yufei"
	resp, err = clt.Get(url)
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("err: %v", err)
		os.Exit(1)
	}
	resp.Body.Close()
	fmt.Printf("response.body: %v\n", string(body))
}
