package controllers

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Emmrys-Jay/my-photo-blog/models"
	"github.com/Emmrys-Jay/my-photo-blog/token"
)

// ReadPics handles read request in the home page
func (m *muxVar) ReadPics(w http.ResponseWriter, r *http.Request) {

	var err error
	var rows = []models.PicInfo{}

	rows, _ = models.GetPics()
	if err != nil {
		if err != sql.ErrNoRows {
			check(w, err)
			return
		}
	}

	// Get the signed in username from the token.

	if uname, ok := alreadyLoggedIn(r); ok {
		response := struct {
			Username string
			Rows     []models.PicInfo
		}{
			Username: uname,
			Rows:     rows,
		}

		tpl.ExecuteTemplate(w, "index-logged.gohtml", response)
		return
	}
	tpl.ExecuteTemplate(w, "index.gohtml", rows)
}

// Addpic handles authorized client add pictures requests
func (m *muxVar) Addpic(w http.ResponseWriter, r *http.Request) {

	if _, ok := alreadyLoggedIn(r); !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	//get currently stored cookie
	c, _ := r.Cookie("token")
	tokenString := c.Value
	tokenMaker := token.NewJWTMaker()
	payload, err := tokenMaker.VerifyToken(tokenString)
	if err != nil {
		check(w, err)
		return
	}

	if r.Method == http.MethodPost {
		// Modify or Refresh the payload by updating the payload info
		payload = &token.Payload{
			Username:  payload.Username,
			IssuedAt:  time.Now(),
			ExpiresAt: time.Now().Add(time.Minute * 5),
		}

		// Create and set new token in a cookie
		newToken, err := tokenMaker.CreateToken(payload)
		if err != nil {
			check(w, err)
			return
		}
		setCookie(w, newToken)

		t := r.FormValue("picname")
		d := r.FormValue("desc")
		uname := payload.Username

		//Work on the file inputed by getting the file extension and storing the file
		//using sha1 in conjunction with the username and multipart file details.
		//i created a file in a particular directory on the server, then sent the full
		//file path to the postgresql database running on AWS
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

	tpl.ExecuteTemplate(w, "addpics.gohtml", payload)
}

// UpdatePic handles authorized client update requests
func (m *muxVar) UpdatePic(w http.ResponseWriter, r *http.Request) {
	if _, ok := alreadyLoggedIn(r); !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	var rowStruct models.PicInfo
	var err error

	//get params from the url
	uname := r.FormValue("uname")
	photoPath := r.FormValue("pic")

	//check if the url contains no uname and pic params
	if uname == "" || photoPath == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	//get currently stored cookie
	c, _ := r.Cookie("token")
	tokenString := c.Value
	tokenMaker := token.NewJWTMaker()
	payload, err := tokenMaker.VerifyToken(tokenString)
	if err != nil {
		check(w, err)
		return
	}

	if uname == payload.Username { //compare the current signed in user to the user who added the picture.

		//get all values stored of the single picture stored in the db
		rowStruct, err = models.GetOnePic(uname, photoPath)
		if err != nil {
			check(w, err)
			return
		}

		if r.Method == http.MethodPost {

			// Modify payload so it can be used to create a new token
			payload = &token.Payload{
				Username:  payload.Username,
				IssuedAt:  time.Now(),
				ExpiresAt: time.Now().Add(time.Minute * 5),
			}

			// Create and set new token in a cookie
			newToken, err := tokenMaker.CreateToken(payload)
			if err != nil {
				check(w, err)
				return
			}
			setCookie(w, newToken)

			//get current form user inputs
			t := r.FormValue("picname")
			d := r.FormValue("desc")

			//Get the newly added file (if added)
			mf, fh, err := r.FormFile("pic")
			if err != nil {

				//call function to add the updated details of the image since the image wasn't changed
				err := models.UpdatePic(uname, photoPath, t, d, "")
				if err != nil {
					check(w, err)
					return
				}
				redirectToHome(w, r)
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
			redirectToHome(w, r)
			return
		}
	} else {
		fmt.Fprintln(w, "Oops! \n You aren't the owner of this picture, therefore you can't update it")
		return
	}

	tpl.ExecuteTemplate(w, "update.gohtml", rowStruct)
}

// DeletePic handles authorized clients delete requests
func (m *muxVar) DeletePic(w http.ResponseWriter, r *http.Request) {

	//check if user is already logged in
	if _, ok := alreadyLoggedIn(r); !ok {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	//get currently stored cookie
	c, _ := r.Cookie("token")
	tokenString := c.Value
	tokenMaker := token.NewJWTMaker()
	payload, err := tokenMaker.VerifyToken(tokenString)
	if err != nil {
		check(w, err)
		return
	}

	//get params from the url
	uname := r.FormValue("uname")
	photoPath := r.FormValue("pic")

	//check if current signed in user is the owner of the image
	if uname != payload.Username {
		fmt.Fprintln(w, "Oops! \n You aren't the owner of this picture, therefore you can't delete it")
		return
	}

	payload = &token.Payload{
		Username:  payload.Username,
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Minute * 5),
	}

	// Create and set new token in a cookie
	newToken, err := tokenMaker.CreateToken(payload)
	if err != nil {
		check(w, err)
		return
	}
	setCookie(w, newToken)

	//delete the file stored in server
	wd, err := os.Getwd()
	if err != nil {
		check(w, err)
	}
	f := filepath.Join(wd, photoPath)
	os.Remove(f)

	//delete picture from db
	err = models.DeletePicture(uname, photoPath)
	if err != nil {
		check(w, err)
		return
	}
	redirectToHome(w, r)
}

// SearchRequestResponse is the datatype where the data is stored before being sent to the searchgohtml template
type SearchRequestResponse struct {
	Query string
	Rows  []models.PicInfo
}

// SearchPics handles client search requests
func (m *muxVar) SearchPics(w http.ResponseWriter, r *http.Request) {
	//Define variable to store search keyword
	var searchKeyword string

	// Define variable for storing results
	var rows = []models.PicInfo{}
	var err error

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else {
		//get param from the form field
		searchKeyword = r.FormValue("search")

		// Pass search keyword into the db function
		rows, err = models.SearchPics(searchKeyword)
		if err != nil {
			// Handle error when search keyword is not found
			if err == sql.ErrNoRows {
				fmt.Println("No results found! \nPlease go back and try with another keyword.")
				return
			} else {
				check(w, err)
				return
			}
		}
	}

	// Create a new struct so search keyword can be added to the values passed to gohtml template
	response := &SearchRequestResponse{
		Query: searchKeyword,
		Rows:  rows,
	}

	// Decide which gohtml template to render based on if user is logged in or not
	// Get username of the signed in user if logged in
	if uname, ok := alreadyLoggedIn(r); ok {

		ResponseWithUsername := struct {
			Username string
			Response *SearchRequestResponse
		}{
			Username: uname,
			Response: response,
		}

		tpl.ExecuteTemplate(w, "search-result-logged.gohtml", ResponseWithUsername)
	} else {
		tpl.ExecuteTemplate(w, "search-result.gohtml", response)
	}

}

// rediectToHome redirects a page to the homepage
func redirectToHome(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
