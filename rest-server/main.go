package main

import (
	calculator "api"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

func index() {
	fmt.Println("Inside index function...")
}
func uri_init() {
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
		index()
	}).Methods("POST")
}

func main() {
	uri_init()
	fmt.Println("Hello!!!")
	log.Fatal(http.ListenAndServe(":5000", router))
}
