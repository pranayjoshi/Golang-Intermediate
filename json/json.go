package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"name"`
	Price    int
	Platform string
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	Courses := []course{
		{"C-1", 200, "Coursera", "123", []string{"c-1", "c-2"}},
		{"C-1", 200, "Coursera", "123", []string{"c-1", "c-2"}},
		{"C-1", 200, "Coursera", "123", nil},
	}
	finalJSON, _ := json.MarshalIndent(Courses, "", "\t")
	fmt.Printf("%s/n", finalJSON)
}

func DecodeJson() {
	jsonData := []byte(`
	{
		"name": "C-1",
		"Price": 200,
		"Platform": "Coursera",
		"tags": [
				"c-1",
				"c-2"
		]
	}
	`)

	var Course course

	check := json.Valid(jsonData)
	if check {
		fmt.Println("Json valid")
		json.Unmarshal(jsonData, &Course)
		fmt.Printf("%#v\n", Course)
	} else {
		fmt.Println("Not valid")
	}
}
