package controllers

import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Emmrys-Jay/my-photo-blog/models"
)

func (m *muxVar) ReadPics(w http.ResponseWriter, r *http.Request) {
	var err error
	var rows []models.PicInfo

	rows, err = models.GetPics()
	if err != nil {
		check(w, err)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", rows)
}

func (m *muxVar) Addpic(w http.ResponseWriter, r *http.Request) {

	if !alreadyLoggedIn(w, r) {
		redirectToView(w, r)
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
		err1 := models.AddPicture(uname, t, d, path)
		if err1 != nil {
			check(w, err)
			return
		}
	}

	tpl.ExecuteTemplate(w, "addpics.gohtml", nil)
}

func (m *muxVar) UpdatePic(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(w, r) {
		redirectToView(w, r)
		return
	}
	var rowStruct models.PicInfo
	var err error

	//get params from the url
	uname := r.FormValue("uname")
	photoPath := r.FormValue("pic")

	//check if the url contains no uname and pic params
	if uname == "" || photoPath == "" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	//get currently stored cookie
	c, _ := r.Cookie("session")
	unameCookie := strings.Split(c.Value, "|")[1] //get username from the current cookie

	if uname == unameCookie { //compare the current signed in user to the user who added the picture.

		//get all values stored of the single picture stored in the db
		rowStruct, err = models.GetOnePic(uname, photoPath)
		if err != nil {
			check(w, err)
			return
		}

		if r.Method == http.MethodPost {

			//get current form user inputs
			t := r.FormValue("picname")
			d := r.FormValue("desc")

			//Get the newly added file (if added)
			mf, fh, err := r.FormFile("pic")
			if err != nil {

				//call function to add the updated details of the image since the image wasn't changed
				err := models.UpdatePic(uname, photoPath, t, d, "")
				if err != nil {
					//check(w, err)
					return
				}
				redirectToView(w, r)
				return
			}
			defer mf.Close()

			//create a new file for the new picture added
			ext := strings.Split(fh.Filename, ".")[1] //gets the extension of the file name
			h := sha1.New()
			io.Copy(h, mf)
			fname := fmt.Sprintf("%x", h.Sum([]byte(uname))) + "." + ext
			// get working directory, create file path and create new file
			wd, err := os.Getwd()
			if err != nil {
				check(w, err)
				return
			}

			//delete the old picture from the server
			toBeDeleted := filepath.Join(wd, photoPath)
			os.Remove(toBeDeleted)

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

			path = filepath.Join("views", "pics", fname) //mpdify the path variable to the form to be stored on the database

			//update the database after a new file has been added
			err = models.UpdatePic(uname, photoPath, t, d, path)
			if err != nil {
				check(w, err)
				return
			}
			redirectToView(w, r)
			return
		}
	} else {
		fmt.Fprintln(w, "Oops! \n You aren't the owner of this picture, therefore you can't update it")
		return
	}

	tpl.ExecuteTemplate(w, "update.gohtml", rowStruct)
}

func (m *muxVar) DeletePic(w http.ResponseWriter, r *http.Request) {

	//check if user is already logged in
	if !alreadyLoggedIn(w, r) {
		redirectToView(w, r)
		return
	}

	//get currently stored cookie
	c, _ := r.Cookie("session")
	unameCookie := strings.Split(c.Value, "|")[1] //get username from the current cookie

	//get params from the url
	uname := r.FormValue("uname")
	photoPath := r.FormValue("pic")

	//check if current signed in user is the owner of the image
	if uname != unameCookie {
		fmt.Fprintln(w, "Oops! \n You aren't the owner of this picture, therefore you can't delete it")
		return
	}

	//delete the file stored in server
	wd, err := os.Getwd()
	if err != nil {
		check(w, err)
	}
	f := filepath.Join(wd, photoPath)
	os.Remove(f)

	//delete pricture from db
	err = models.DeletePicture(uname, photoPath)
	if err != nil {
		check(w, err)
		return
	}
	//models.DisplayNum = 0
	redirectToView(w, r)
}

func redirectToView(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view", http.StatusSeeOther)
}
