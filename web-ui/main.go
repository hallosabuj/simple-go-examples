package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	temp, _ := template.ParseGlob("static/*.html")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "home.html", nil)
	})
	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("static/contact.html", "static/header-footer.html")
		type Contact struct {
			Name  string
			Phone string
			Marks map[string]string
		}
		temp.Execute(w, Contact{Name: "Sabuj Mondal", Phone: "9876543210", Marks: map[string]string{"Math": "20", "Hindi": "20"}})
	})
	http.Handle("/asset/", http.StripPrefix("/asset", http.FileServer(http.Dir("./asset"))))

	fmt.Println("Server started at 8888")
	http.ListenAndServe("localhost:8888", nil)
}
