package controllers

import (
	"net/http"
	"text/template"
	"uuid"
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

}

func (m *muxVar) Signin(w http.ResponseWriter, r *http.Request) {

}

func (m *muxVar) Signup(w http.ResponseWriter, r *http.Request) {

}

func (m *muxVar) Add(w http.ResponseWriter, r *http.Request) {

}

func setCookie(w http.ResponseWriter) *http.Cookie {
	uuid := uuid.New()
	c := &http.Cookie{
		Name:  "session",
		Value: uuid,
	}

	http.SetCookie(w, c)
	return c
}
