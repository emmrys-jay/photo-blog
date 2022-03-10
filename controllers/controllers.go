package controllers

import (
	"Golang/my-photo-blog/models"
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type muxVar struct{}

var tpl *template.Template

var numInd int

var alertFunct = template.FuncMap{
	"a1": alertFunc,
}

func init() {
	tpl = template.Must(template.New("").Funcs(alertFunct).ParseGlob("views/templates/*.gohtml"))
}

func GetMuxVar() *muxVar {
	var m muxVar
	return &m
}

func (m *muxVar) Index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func (m *muxVar) Signin(w http.ResponseWriter, r *http.Request) {

	//check if user is already logged in
	if alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
			http.Redirect(w, r, "/", http.StatusSeeOther)
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

func (m *muxVar) Signout(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	c := &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (m *muxVar) Addpic(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "addpics.gohtml", numInd)
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

func alertFunc(numIndicator int) string {
	if numIndicator == 1 {
		return "Picture Added"
	}
	return ""
}
