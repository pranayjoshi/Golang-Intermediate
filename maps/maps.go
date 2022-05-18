package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["en"] = "english"
	languages["hi"] = "hindi"
	fmt.Println(languages)
	delete(languages, "en")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Println(key, value)
	}
}
