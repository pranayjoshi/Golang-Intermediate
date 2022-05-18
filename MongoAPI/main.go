package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pranayjoshi/mongoapi/router"
)

func main() {
	fmt.Println("Mongo API")
	r := router.Router()
	fmt.Println("Server is Starting...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening")
}
