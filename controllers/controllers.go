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

// muxVar is used to pass handler methods to main
type muxVar struct{}

var tpl *template.Template

// init parses all gohtml templates
func init() {
	tpl = template.Must(template.ParseGlob("views/templates/*.gohtml"))
}

// GetMuxVar is used to pass muxVar variable to main.go
func GetMuxVar() *muxVar {
	var m muxVar
	return &m
}

// Signin handles sign in requests from user
func (m *muxVar) Signin(w http.ResponseWriter, r *http.Request) {

	//check if user is already logged in
	if _, ok := alreadyLoggedIn(r); ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
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
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		} else {
			fmt.Fprint(w, "Invalid Username or Password")
			return
		}
	}

	tpl.ExecuteTemplate(w, "signin.gohtml", nil)
}

// Signup handles sign up request from user
func (m *muxVar) Signup(w http.ResponseWriter, r *http.Request) {
	if _, ok := alreadyLoggedIn(r); ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
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
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
}

// Signout handles sign out requests from user
func (m *muxVar) Signout(w http.ResponseWriter, r *http.Request) {

	if _, ok := alreadyLoggedIn(r); !ok {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	c := &http.Cookie{
		Name:   "token",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// setCookie stores the generated JWT token in a cookie
func setCookie(w http.ResponseWriter, token string) {
	// u := uuid.New().String()
	// uuid := u + "|" + uname

	c := &http.Cookie{
		Name:  "token",
		Value: token,
	}

	http.SetCookie(w, c)
}

// alreadyLoggedIn checks if a user is ligged in by getting token from browser cookie
func alreadyLoggedIn(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return "", false
	}
	jwtMaker := token.NewJWTMaker()
	token := cookie.Value

	payload, err := jwtMaker.VerifyToken(token)
	if err != nil {
		return "", false
	}

	return payload.Username, err == nil
}

// check is used to return an Internal Server Error to the client
func check(w http.ResponseWriter, err error) {
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
