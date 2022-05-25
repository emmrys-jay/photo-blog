package controllers

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Emmrys-Jay/my-photo-blog/models"

	"github.com/google/uuid"
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
		ps, err := base64.StdEncoding.DecodeString(pd)
		if err != nil {
			check(w, err)
			return
		}

		if string(ps) == p {
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

		//encode password using base64 encoding
		ps := base64.StdEncoding.EncodeToString([]byte(p))

		//add new user to the database
		err := models.Adduser(w, u, e, ps)
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

func setCookie(w http.ResponseWriter, uname string) {
	u := uuid.New().String()
	uuid := u + "|" + uname

	c := &http.Cookie{
		Name:  "session",
		Value: uuid,
	}

	http.SetCookie(w, c)
}

func alreadyLoggedIn(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("session")

	return err == nil
}

func check(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
