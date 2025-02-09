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
	CourseID    string  `json:"courseiD"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"` //`json:"courseprice"` `json:"-" it will not print
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

//fake DB

var course []Course

// middleware, helper -file

func (c *Course) Isempty() bool {

	//if  c.CourseID != "" && c.CourseName !=""{

	//}

	//return c.CourseID == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {

	fmt.Println("Checking")

	r := mux.NewRouter()

	//seeding

	course = append(course, Course{CourseID: "2", CourseName: "Golang", CoursePrice: 499, Author: &Author{Fullname: "Vijay", Website: "www.Golearn.com"}})

	course = append(course, Course{CourseID: "4", CourseName: "Python", CoursePrice: 499, Author: &Author{Fullname: "Vijay", Website: "www.python.com"}})

	// routing

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/course", getallcourse).Methods("GET")
	r.HandleFunc("/course/{id}", getonecourse).Methods("GET")
	r.HandleFunc("/Course", createonecouesr).Methods("POST")
	r.HandleFunc("/course/{id}", updateonecourse).Methods("PUT")
	r.HandleFunc("/Course/{id}", deleteonecourse).Methods("DELETE")
	// listen to port

	log.Fatal(http.ListenAndServe(":4000", r))

}

// controllere -file
// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("<h1> Welcome Vijay </h1>"))

	// r is get value from whoever sending value
	// w where you write the value
}

func getallcourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get All course")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(course)
}

func getonecourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get one course")
	w.Header().Set("content-Type", "application/json")

	//grap id from request
	params := mux.Vars(r)

	// loop through courses, find matching Id and return responce

	for _, co := range course {

		if co.CourseID == params["id"] {

			json.NewEncoder(w).Encode(co)
			return
		}

	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createonecouesr(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Create one course")
	w.Header().Set("content-Type", "application/json")

	// what if:  body is empty

	if r.Body == nil {

		json.NewEncoder(w).Encode("Please send some data")
	}

	//what about -{}

	var co Course

	_ = json.NewDecoder(r.Body).Decode(&co)

	if co.Isempty() {

		json.NewEncoder(w).Encode("No data inside")
		return
	}

	//Todo: check only if titile is dublicate
	//loop,titile matches with coursenmae and json reponce and saying this is alreday exist.

	// generate one unique id
	//append newcouse into course
	rand.Seed(time.Now().UnixNano())
	co.CourseID = strconv.Itoa(rand.Intn(100)) // this int value converting as astring.
	course = append(course, co)
	json.NewEncoder(w).Encode(co)
	return
}

func updateonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("content-Type", "application/json")
	// first - grap id from req

	params := mux.Vars(r)

	// loop,id , remove add with my id

	for index, co := range course {

		if co.CourseID == params["id"] {

			course = append(course[:index], course[index+1:]...)
			var co Course
			_ = json.NewDecoder(r.Body).Decode(&co)
			co.CourseID = params["id"]
			course = append(course, co)
			json.NewEncoder(w).Encode(co)
			return
		}
	}
	//Todo send a responce when id is not found
}

func deleteonecourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Deete one course")
	w.Header().Set("content-Type", "application/json")
	//loop, match id remove

	params := mux.Vars(r)

	for index, Course := range course {

		if Course.CourseID == params["id"] {
			course = append(course[:index], course[index+1:]...)
			break
		}

	}

}
