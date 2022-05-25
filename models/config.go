package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() {
	db, err = sql.Open("postgres", "postgres:my_photo_blog-emmrys@tcp(database-1.cpyezbxep7pq.us-east-2.rds.amazonaws.com)/myphotoblog?charset=utf8")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you connected to database successfully")
}
