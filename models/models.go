package models

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB
var err error

func ConnectDB() {
	db, err = sql.Open("mysql", "admin:my_photo_blog-emmrys@tcp(database-2.cpyezbxep7pq.us-east-2.rds.amazonaws.com)/myphotoblog?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you connected to database successfully")
}

func Adduser(w http.ResponseWriter, uname, email, psword string) error {
	stmt := fmt.Sprintf("INSERT INTO myphotoblog.userspb VALUES('%s', '%s', '%s')", uname, email, psword)
	_, err = db.Exec(stmt)
	return err
}

func AddPicture(uname, pname, desc, photop string) error {
}

func GetUser(uname string) *sql.Row {
	stmt := fmt.Sprintf("SELECT uname, psword FROM myphotoblog.userspb WHERE uname = '%s'", uname)
	row := db.QueryRow(stmt)
	return row
}
