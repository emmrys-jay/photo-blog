package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() {

	// Connect ot remote database
	// Edit details when connecting to your PostgreSQL server instance
	db, err = sql.Open("postgres", "user=postgres dbname=myphotoblog password=my_photo_blog-emmrys host=database-1.cpyezbxep7pq.us-east-2.rds.amazonaws.com sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Ping database to see if connection has been successfully established
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("you connected to database successfully")
}
