package calculator

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Addition(w http.ResponseWriter, r *http.Request) {
	log.Println("Addition")
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
