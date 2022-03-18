package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice string  `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// Middleware

func (c *Course) isEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}
func main() {
	fmt.Println("d")

	r := mux.NewRouter()
	r.HandleFunc("/", serve).Methods("GET")

	log.Fatal(http.ListenAndServe(":4000", r))
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Welcome to Golang </h1>"))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseID == params["ID"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course")
	return
}
