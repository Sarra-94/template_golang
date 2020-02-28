package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//Create the Model of students
type Student struct {
	Name     string
	LastName string
	Email    string
	Note     int
}
/* liste of students*/
var MyStudent = []Student{
	{"Sarra",
		"Guesmi",
		"sarra@gmail.com",
		15},
	{"amen",
		"y",
		"sarra@gmail.com",
		5},
	{"ameny",
		"x",
		"sarra@gmail.com",
		10}}
func showSuccess(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test success")
	tpl, _ := template.ParseFiles("./success.html")
	err := tpl.Execute(w, MyStudent)
	if err != nil {
		fmt.Println(err)
	}
}
func showFailer(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Test success")
	tpl, _ := template.ParseFiles("./failer.html")
	err := tpl.Execute(w, MyStudent)
	if err != nil {
		fmt.Println(err)
	}
}
func showStudents(w http.ResponseWriter, r *http.Request) {

	tpl, _ := template.ParseFiles("./students.html")
	err := tpl.Execute(w, MyStudent)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", showStudents)
	mux.HandleFunc("/succeess", showSuccess)
	mux.HandleFunc("/failer", showFailer)
	fmt.Printf("http://localhost:5000/")
	http.ListenAndServe(":5000", mux)

}