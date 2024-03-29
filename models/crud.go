package models

import (
	"database/sql"
	"net/http"
)

var db *sql.DB
var err error

// PicInfo stores is the DB struct model
type PicInfo struct {
	Id     int
	Uname  string
	Pname  string
	Photop string
	Desc   string
}

// Adduser adds a new user to the database
func Adduser(w http.ResponseWriter, uname, email, psword string) error {
	_, err = db.Exec(`INSERT INTO userspb VALUES($1, $2, $3);`, uname, email, psword)
	return err
}

// AddPicture adds a new picture to the database
func AddPicture(uname, pname, desc, photop string) error {
	_, err := db.Exec(`INSERT INTO photob(uname, ptitle, photo, descp) VALUES($1, $2, $3, $4)`, uname, pname, photop, desc)
	return err
}

// GetUser gets a user from the database during sign in requests
func GetUser(uname string) *sql.Row {
	row := db.QueryRow(`SELECT uname, psword FROM userspb WHERE uname = $1;`, uname)
	return row
}

// GetOnePic gets a single picture from the database during update requests
func GetOnePic(uname, photop string) (PicInfo, error) {
	var p PicInfo
	rows, err := db.Query(`SELECT * FROM photob 
						   WHERE uname = $1 AND photo = $2 LIMIT 1;`, uname, photop)
	if err != nil {
		return p, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Uname, &p.Pname, &p.Photop, &p.Desc)
	}
	if err != nil {
		return p, err
	}
	return p, err
}

// GetPics gets all pictures from the database during read requests
func GetPics() ([]PicInfo, error) {
	var p PicInfo
	var SliceOfP []PicInfo
	var rows *sql.Rows
	var err error

	rows, err = db.Query(`SELECT * FROM photob ORDER BY id ASC;`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Uname, &p.Pname, &p.Photop, &p.Desc)
		if err != nil {
			return nil, err
		}
		SliceOfP = append(SliceOfP, p)
	}
	return SliceOfP, nil
}

// DeletePicture deletes a picture from the database
func DeletePicture(uname, photoPath string) error {
	_, err := db.Exec("DELETE FROM photob WHERE uname = $1 AND photo = $2", uname, photoPath)
	return err
}

// UpdatePic updates a single picture in the database
func UpdatePic(olduname, oldphotoPath, pname, desc, photop string) error {
	if photop == "" {
		_, err := db.Exec(`UPDATE photob 
						  SET ptitle = $1, descp = $2
						  WHERE uname = $3 AND photo = $4`, pname, desc, olduname, oldphotoPath)

		return err
	}
	_, err := db.Exec(`UPDATE photob 
					  SET ptitle = $1, descp = $2, photo = $3
					  WHERE uname = $4 AND photo = $5`, pname, desc, photop, olduname, oldphotoPath)
	return err
}

// SearchPics searches and returns pictures from the database
func SearchPics(query string) ([]PicInfo, error) {
	var p PicInfo
	var SliceOfP []PicInfo
	var rows *sql.Rows
	var err error

	rows, err = db.Query(`SELECT * FROM photob 
				WHERE ptitle ~ $1 OR descp ~ $2`, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Uname, &p.Pname, &p.Photop, &p.Desc)
		if err != nil {
			return nil, err
		}
		SliceOfP = append(SliceOfP, p)
	}
	return SliceOfP, nil
}
