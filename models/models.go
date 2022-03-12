package models

import (
	"database/sql"
	"fmt"
	"net/http"
)

var db *sql.DB
var err error

type PicInfo struct {
	Id     int
	Uname  string
	Pname  string
	Photop string
	Desc   string
}

var Pics []PicInfo

var DisplayNum int8
var DbRowsAdded int

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
	_, err = db.Exec(`INSERT INTO myphotoblog.userspb VALUES(?, ?, ?);`, uname, email, psword)
	return err
}

func AddPicture(uname, pname, desc, photop string) (sql.Result, error) {
	res, err := db.Exec("INSERT INTO myphotoblog.photob(uname, ptitle, photo, descp) VALUES(?, ?, ?, ?)", uname, pname, photop, desc)
	return res, err
}

func GetUser(uname string) *sql.Row {
	row := db.QueryRow(`SELECT uname, psword FROM myphotoblog.userspb WHERE uname = ?;`, uname)
	return row
}

func GetPics(dbRows int) error {
	var p PicInfo
	var rows *sql.Rows
	var err error
	if dbRows == 0 {
		rows, err = db.Query(`SELECT * FROM myphotoblog.photob ORDER BY id ASC;`)
	} else {
		stmt := fmt.Sprintf(`SELECT * FROM (
			SELECT * FROM myphotoblog.photob ORDER BY id DESC LIMIT %d
		 )Var1 ORDER BY id ASC;`, dbRows)
		rows, err = db.Query(stmt)
	}
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Uname, &p.Pname, &p.Photop, &p.Desc)
		if err != nil {
			return err
		}
		Pics = append(Pics, p)
	}
	return nil
}
