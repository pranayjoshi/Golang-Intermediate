package main

import "fmt"

func main() {
	fmt.Println("hi")
	days := []string{"monday", "tuesday", "thursday", "friday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}
}
