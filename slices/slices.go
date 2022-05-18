package main

import "fmt"

func main() {
	var fruit = []string{}
	fmt.Printf("%T\n", fruit)
	fruit = append(fruit, "Mango", "Banana", "Peach", "Orange")
	fmt.Println(fruit)
	fruit = append(fruit[1:3])
	fmt.Println(fruit)
	highScores := make([]int, 4)
	highScores[1] = 4
	highScores[3] = 9
	fmt.Println(highScores)

	// remove val based on index
	courses := []string{"python", "javascript", "java", "go"}

	index := 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
