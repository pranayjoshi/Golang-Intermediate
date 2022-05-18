package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	pranay := User{"Pranay", "pj@gmail.com", true, 16}
	fmt.Println(pranay)
}
