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
	pranay.GetStatus()
	pranay.NewEmail()
}

func (U User) GetStatus() {
	fmt.Println("is user there? ", U.Status)
}

func (U User) NewEmail() {
	U.Email = "sa@gma.com"
	fmt.Println("New Email ", U.Email)
}
