package main

import (
	"net/http"

	"github.com/Emmrys-Jay/my-photo-blog/controllers"
	"github.com/Emmrys-Jay/my-photo-blog/models"
)

func main() {
	models.ConnectDB()
	m := controllers.GetMuxVar()
	mux := http.NewServeMux()

	mux.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
	mux.HandleFunc("/", m.ReadPics)
	mux.HandleFunc("/signin", m.Signin)
	mux.HandleFunc("/signup", m.Signup)
	mux.HandleFunc("/add", m.Addpic)
	mux.HandleFunc("/logout", m.Signout)
	mux.HandleFunc("/update", m.UpdatePic)
	mux.HandleFunc("/delete", m.DeletePic)
	mux.HandleFunc("/search", m.SearchPics)
	http.ListenAndServe(":8080", mux)
}
