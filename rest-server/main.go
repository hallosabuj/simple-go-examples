package main

import (
	calculator "api/calculator"
	test "api/test"
	homepage "homepage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func uri_init_calculator() {
	router.HandleFunc("/calculator/addition", func(w http.ResponseWriter, r *http.Request) {
		calculator.Addition(w, r)
	}).Methods("POST")
	router.HandleFunc("/calculator/subtraction", func(w http.ResponseWriter, r *http.Request) {
		calculator.Subtraction(w, r)
	}).Methods("POST")
	router.HandleFunc("/calculator/division", func(w http.ResponseWriter, r *http.Request) {
		calculator.Division(w, r)
	}).Methods("POST")
	router.HandleFunc("/calculator/multiplication", func(w http.ResponseWriter, r *http.Request) {
		calculator.Multiplication(w, r)
	}).Methods("POST")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homepage.Index(w, r)
	}).Methods("GET")
}

func uri_init_test() {
	router.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		test.Greeting(w, r)
	}).Methods("GET")
	router.HandleFunc("/test/{name}", func(w http.ResponseWriter, r *http.Request) {
		test.GreetingWithName(w, r)
	}).Methods("GET")
}

func main() {
	uri_init_calculator()
	uri_init_test()
	log.Println("First log for rest server!!!")
	log.Println("Rest server satrted at port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
