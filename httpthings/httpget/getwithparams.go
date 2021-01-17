package httpget

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

func GetWithParams() {
	params := url.Values{}
	//Url, err := url.Parse("http://httpbin.org/get")
	//if err != nil {
	//	return
	//}
	req, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	params.Set("name", "zhaofan")
	params.Set("name", "yufei") // it will overwrite zhaofan
	params.Set("age", "23")
	//如果参数中有中文参数,这个方法会进行URLEncode
	req.URL.RawQuery = params.Encode()
	fmt.Printf("Url.RawQuery: %v\n", req.URL.RawQuery)
	urlPath := req.URL.String()
	fmt.Println(urlPath) // https://httpbin.org/get?age=23&name=zhaofan
	resp, err := http.Get(urlPath)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func GetWithParams2() {
	req, err := http.NewRequest("GET", "http://api.themoviedb.org/3/tv/popular", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}
	q := req.URL.Query()
	q.Add("api_key", "key_from_environment_or_flag")
	q.Add("another_thing", "foo & bar")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
}
