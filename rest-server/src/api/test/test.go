package test

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Message struct {
	Msg string `json:"msg" example:"Pepsi Inc."`
}

func Greeting(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Greeting")
	response := Message{Msg: "Hello from test..."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GreetingWithName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Greeting with name")
	//Fetching name from url
	vars := mux.Vars(r)
	name := vars["name"]
	response := Message{Msg: "Hello " + name + " welcome to test..."}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
