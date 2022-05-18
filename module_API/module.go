package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID    string  `json:"id"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

// Middleware

func (c *Course) isEmpty() bool {
	// return c.CourseID == "" && c.CourseName == ""
	return c.CourseID == ""
}
func main() {
	fmt.Println("d")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseID: "2", CourseName: "Golang", CoursePrice: 299, Author: &Author{Fullname: "Pranay Joshi", Website: "pranayjoshi.github.io"}})
	courses = append(courses, Course{CourseID: "3", CourseName: "Mern", CoursePrice: 249, Author: &Author{Fullname: "Pranay Joshi", Website: "pranayjoshi.github.io"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")
	// listen to the port
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
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	//empty body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("JSON: Please send some data")
	}
	for _, c := range courses {
		if course.CourseName == c.CourseName {
			fmt.Println("Same course")
			return
		}
	}
	rand.Seed(time.Now().Unix())
	course.CourseID = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	// grab id from request url
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO:- send a response when id not found
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("successfully deleted!")
			break
		}
	}
}
