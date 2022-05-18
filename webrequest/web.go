package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://beyondwelkin.com/"

func main() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%T\n", response)
	defer response.Body.Close()
	d, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(d)
	fmt.Println(content)
}
