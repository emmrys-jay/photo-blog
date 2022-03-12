package controllers

import (
	"Golang/my-photo-blog/models"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type muxVar struct{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/templates/*.gohtml"))
}

func GetMuxVar() *muxVar {
	var m muxVar
	return &m
}

func (m *muxVar) Index(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		if models.DisplayNum == 0 {
			err := models.GetPics(models.DbRowsAdded)
			if err != nil {
				check(w, err)
				return
			}
			models.DisplayNum = 1
			models.DbRowsAdded = 0
		}
	}

	tpl.ExecuteTemplate(w, "index.gohtml", models.Pics)
}

func (m *muxVar) Signin(w http.ResponseWriter, r *http.Request) {

	//check if user is already logged in
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	if r.Method == http.MethodPost {
		u := r.FormValue("uname")
		p := r.FormValue("psword")

		var pd, un string
		//get the user details from db
		row := models.GetUser(u)
		err := row.Scan(&un, &pd)
		if err == sql.ErrNoRows {
			fmt.Fprint(w, "Invalid Username or Password")
			return
		}

		if pd == p {
			//set cookie
			setCookie(w, u)

			//redirect to home page
			http.Redirect(w, r, "/view", http.StatusSeeOther)
			return
		} else {
			fmt.Fprint(w, "Invalid Username or Password")
			return
		}
	}

	tpl.ExecuteTemplate(w, "signin.gohtml", nil)
}

func (m *muxVar) Signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/view", http.StatusSeeOther)
	}

	if r.Method == http.MethodPost {
		//get form inputs
		e := r.FormValue("email")
		u := r.FormValue("uname")
		p := r.FormValue("psword")

		//encrypt password using sha256

		//add new user to the database
		err := models.Adduser(w, u, e, p)
		if err != nil {
			check(w, err)
			return
		}

		//set cookie
		setCookie(w, u)

		//redirect to home page
		http.Redirect(w, r, "/view", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func (m *muxVar) Signout(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/view", http.StatusSeeOther)
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("session")

	return err == nil
}

func check(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
