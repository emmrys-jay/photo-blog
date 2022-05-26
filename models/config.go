package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() {
	db, err = sql.Open("postgres", "user=postgres dbname=myphotoblog password=my_photo_blog-emmrys host=database-1.cpyezbxep7pq.us-east-2.rds.amazonaws.com sslmode=disable")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you connected to database successfully")
}
