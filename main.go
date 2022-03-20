package main

import (
	"Golang/my-photo-blog/controllers"
	"Golang/my-photo-blog/models"
	"net/http"
)

func main() {
	models.ConnectDB()
	m := controllers.GetMuxVar()
	http.Handle("/", http.StripPrefix("/views/pics", http.FileServer(http.Dir("./views/pics"))))
	http.HandleFunc("/view", m.ReadPics)
	http.HandleFunc("/signin", m.Signin)
	http.HandleFunc("/signup", m.Signup)
	http.HandleFunc("/add", m.Addpic)
	http.HandleFunc("/logout", m.Signout)
	http.HandleFunc("/update", m.UpdatePic)
	http.HandleFunc("/delete", m.DeletePic)
	http.ListenAndServe(":8080", nil)
}
