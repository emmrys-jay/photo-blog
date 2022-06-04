package main

import (
	"net/http"

	"github.com/Emmrys-Jay/my-photo-blog/controllers"
	"github.com/Emmrys-Jay/my-photo-blog/models"
)

func main() {
	models.ConnectDB()
	m := controllers.GetMuxVar()
	// http.Handle("/css/", http.StripPrefix("/views/", http.FileServer(http.Dir("views/templates"))))
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("views"))))
	http.HandleFunc("/", m.ReadPics)
	http.HandleFunc("/signin", m.Signin)
	http.HandleFunc("/signup", m.Signup)
	http.HandleFunc("/add", m.Addpic)
	http.HandleFunc("/logout", m.Signout)
	http.HandleFunc("/update", m.UpdatePic)
	http.HandleFunc("/delete", m.DeletePic)
	http.HandleFunc("/search", m.SearchPics)
	http.ListenAndServe(":8080", nil)
}
