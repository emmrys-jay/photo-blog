package controllers

import (
	"Golang/my-photo-blog/models"
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func (m *muxVar) Addpic(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(w, r) {
		http.Redirect(w, r, "/view", http.StatusSeeOther)
		return
	}

	//get currently stored cookie
	c, _ := r.Cookie("session")
	uname := strings.Split(c.Value, "|")[1] //get username from the current cookie

	if r.Method == http.MethodPost {
		t := r.FormValue("picname")
		d := r.FormValue("desc")

		//Work on the file inputed by getting the file extension and storing the file
		//using sha1 in conjunction with the username and multipart file details.
		//i created a file in a particular directory on the server, then sent the full
		//file path to the mysql database running on AWS
		mf, fh, err := r.FormFile("pic")
		if err != nil {
			check(w, err)
			return
		}
		defer mf.Close()
		// create sha1 hash for file name
		ext := strings.Split(fh.Filename, ".")[1] //gets the extension of the file name
		h := sha1.New()
		io.Copy(h, mf)
		fname := fmt.Sprintf("%x", h.Sum([]byte(uname))) + "." + ext
		// create new file
		wd, err := os.Getwd()
		if err != nil {
			check(w, err)
			return
		}
		path := filepath.Join(wd, "views", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			check(w, err)
			return
		}
		defer nf.Close()
		// copy all the picture file details into the new file created
		mf.Seek(0, 0)
		io.Copy(nf, mf)
		//End of working with file with the new file created in the webserver

		path = filepath.Join("views", "pics", fname)

		//Insert picture info into the database
		res, err1 := models.AddPicture(uname, t, d, path)
		if err1 != nil {
			check(w, err)
			return
		}
		models.DisplayNum = 0
		n, _ := res.RowsAffected()
		models.DbRowsAdded += int(n)
	}

	tpl.ExecuteTemplate(w, "addpics.gohtml", nil)
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
