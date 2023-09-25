package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Model for courses -file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake db

var courses []Course

//middleware , helper

func IsEmpty(c *Course) bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("API - Vishnu golang")
	r := mux.NewRouter()

	//seeding data

	courses = append(courses, Course{CourseId: "2", CourseName: "Reactjs", CoursePrice: 299, Author: &Author{Fullname: "vishnu", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "1", CourseName: "mern", CoursePrice: 199, Author: &Author{Fullname: " bubu", Website: "lco.dev"}})

	//routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	//listen to port

	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers -file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api by Vishnu in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	//grab id from req

	params := mux.Vars(r)

	//loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return

		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")
	//what if body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about  -{}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if IsEmpty(&course) {
		json.NewEncoder(w).Encode("No data inside the JSON")
		return
	}

	//generate unique id, string
	//append course in to courses

	course.CourseId = strconv.Itoa(rand.Intn(100)) //converts int to str
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req

	params := mux.Vars(r)

	//loop , id, remove, add  with myid

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop id remove

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break

		}

	}
	json.NewEncoder(w).Encode("Course deleted successfully")

}
