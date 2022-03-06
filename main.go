package main

import (
	"Golang/my-photo-blog/controllers"
	"net/http"
)

func main() {
	m := controllers.GetMuxVar()
	http.HandleFunc("/", m.Index)
	http.HandleFunc("/signin", m.Signin)
	http.HandleFunc("/signup", m.Signup)
	http.HandleFunc("/add", m.Add)
	http.ListenAndServe(":8080", nil)
}
