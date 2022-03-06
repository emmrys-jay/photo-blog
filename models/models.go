package models

import (
	"database/sql"
	"mime/multipart"
)

var db *sql.DB

func ConnectDB() {

}

func Adduser(email, name, description string) {

}

func AddPicture(pic *multipart.File, uname, pname, desc string) {

}
