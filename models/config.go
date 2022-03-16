package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

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
