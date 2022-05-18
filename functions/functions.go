package main

import "fmt"

func main() {
	addPrinter("hello", "world")
	total, comp := addInfinite(1, 2, 3, 4, 2)
	fmt.Println(total, comp)
}
func addPrinter(str string, comp_str string) {
	fmt.Println((str + " " + comp_str))
}

func addInfinite(nums ...int) (int, string) {
	total := 0
	for _, i := range nums {
		total += i
	}
	return total, "complement"
}
