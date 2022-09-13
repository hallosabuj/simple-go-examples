package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var router = mux.NewRouter()

//This is a test function
func testFunc(w http.ResponseWriter, r *http.Request) {
	//This will be printed on server log
	fmt.Println("Hello inside testFunc")
}
func uri_init() {
	router.HandleFunc("/", testFunc).Methods("GET")
}

func main() {
	// Register all URIs
	uri_init()
	fmt.Println("Server started...")
	err := http.ListenAndServeTLS(":5002", "example.crt", "example.key", router)
	if err != nil {
		fmt.Println("Error...", err)
	}
}

//=============================================================
//Create certificate and key
// openssl req -newkey rsa:4096 \
//             -x509 \
//             -sha256 \
//             -days 3650 \
//             -nodes \
//             -out example.crt \
//             -keyout example.key \
//             -subj "/C=SI/ST=Ljubljana/L=Ljubljana/O=Security/OU=IT Department/CN=localhost"

//=============================================================
//Create certificate and key
//openssl req -newkey rsa:2048 -new -nodes -x509 -days 3650 -keyout key.pem -out cert.pem
