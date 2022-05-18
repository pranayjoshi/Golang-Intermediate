package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup
var signals = []string{"test"}
var mut sync.Mutex

func main() {
	// go gretter("Hello")
	// gretter("World!")
	websitelist := []string{
		"https://google.com",
		"https://go.dev",
		"https://pranayjoshi.github.io",
		"https://github.com",
	}
	for _, web := range websitelist {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	print(signals)
}

func gretter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(2 * time.Millisecond)
		print(s)
	}
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		log.Fatal("oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println("Status code:", result.StatusCode)
	}

}
