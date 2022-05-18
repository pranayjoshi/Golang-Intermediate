package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func getReq() {
	const url = "http://localhost:8081/get"
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.StatusCode, response.ContentLength)
	content, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func main() {
	getReq()
}
