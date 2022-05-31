package controllers

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"

	"github.com/Emmrys-Jay/my-photo-blog/models"
	"github.com/Emmrys-Jay/my-photo-blog/token"
)

type muxVar struct{}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("views/templates/*.html"))
}

func GetMuxVar() *muxVar {
	var m muxVar
	return &m
}

func (m *muxVar) Signin(w http.ResponseWriter, r *http.Request) {

	//check if user is already logged in
	if alreadyLoggedIn(r) {
		redirectToView(w, r)
		return
	}

	if r.Method == http.MethodPost {
		uname := r.FormValue("uname")
		psword := r.FormValue("psword")

		var pd, un string
		//get the user details from db
		row := models.GetUser(uname)
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

		if string(ps) == psword {
			// get payload and use to create token
			payload := token.NewPayload(uname)
			tokenMaker := token.NewJWTMaker()

			tokenString, err := tokenMaker.CreateToken(payload)
			if err != nil {
				check(w, err)
				return
			}
			//set cookie
			setCookie(w, tokenString)

			//redirect to home page
			http.Redirect(w, r, "/view", http.StatusSeeOther)
			return
		} else {
			fmt.Fprint(w, "Invalid Username or Password")
			return
		}
	}

	tpl.ExecuteTemplate(w, "signin.html", nil)
}

func (m *muxVar) Signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		redirectToView(w, r)
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

		// get payload and use to create token
		payload := token.NewPayload(u)
		tokenMaker := token.NewJWTMaker()

		tokenString, err := tokenMaker.CreateToken(payload)
		if err != nil {
			check(w, err)
			return
		}

		//set cookie
		setCookie(w, tokenString)

		//redirect to home page
		http.Redirect(w, r, "/view", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func (m *muxVar) Signout(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(r) {
		redirectToView(w, r)
		return
	}

	c := &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/view", http.StatusSeeOther)
}

func setCookie(w http.ResponseWriter, token string) {
	// u := uuid.New().String()
	// uuid := u + "|" + uname

	c := &http.Cookie{
		Name:  "token",
		Value: token,
	}

	http.SetCookie(w, c)
}

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("token")
	if err != nil {
		return false
	}
	jwtMaker := token.NewJWTMaker()
	token := cookie.Value

	_, err = jwtMaker.VerifyToken(token)

	return err == nil
}

func check(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
