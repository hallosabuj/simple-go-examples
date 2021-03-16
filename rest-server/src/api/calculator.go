package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Numbers struct {
	First  int16 `json:"num1" example:"Pepsi Inc."`
	Second int16 `json:"num2" example:"Pepsi Inc."`
}
type Result struct {
	First     int16  `json:"num1" example:"Pepsi Inc."`
	Second    int16  `json:"num2" example:"Pepsi Inc."`
	Operation string `json:"operation" example:"Pepsi Inc."`
	Result    int32  `json:"result" example:"Pepsi Inc."`
}

func Addition(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Addition")
	var numbers Numbers
	var result Result
	json.NewDecoder(r.Body).Decode(&numbers)
	result.First = numbers.First
	result.Second = numbers.Second
	result.Operation = "add"
	result.Result = int32(numbers.First) + int32(numbers.Second)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Subtraction(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Subtraction")
	var numbers Numbers
	var result Result
	json.NewDecoder(r.Body).Decode(&numbers)
	result.First = numbers.First
	result.Second = numbers.Second
	result.Operation = "sub"
	result.Result = int32(numbers.First) - int32(numbers.Second)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Multiplication(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Multiplication")
	var numbers Numbers
	var result Result
	json.NewDecoder(r.Body).Decode(&numbers)
	result.First = numbers.First
	result.Second = numbers.Second
	result.Operation = "mult"
	result.Result = int32(numbers.First) * int32(numbers.Second)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func Division(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Division")
	var numbers Numbers
	var result Result
	json.NewDecoder(r.Body).Decode(&numbers)
	result.First = numbers.First
	result.Second = numbers.Second
	result.Operation = "div"
	result.Result = int32(numbers.First) / int32(numbers.Second)
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
