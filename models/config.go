package models

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() {

	// Connect to remote database
	// Edit details when connecting to your PostgreSQL server instance
	db, err = sql.Open("postgres", "user=postgres dbname=myphotoblog password=password host=my-photo-blog_postgres-server_1 sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Ping database to see if connection has been successfully established
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	err = CreateDatabase(db)
	if err != nil {
		panic(err)
	}

	fmt.Println("you connected to database successfully")
}

func CreateDatabase(db *sql.DB) error {
	userspbQuery := `CREATE TABLE userspb (
						uname VARCHAR PRIMARY KEY NOT NULL,
						email VARCHAR UNIQUE NOT NULL,
						psword VARCHAR
					);`
	_, err := db.Exec(userspbQuery)
	if err != nil {
		return err
	}

	photobQuery := `CREATE TABLE photob (
						id BIGSERIAL,
						uname VARCHAR,
						ptitle VARCHAR,
						photo VARCHAR NOT NULL,
						descp VARCHAR DEFAULT NULL,
						PRIMARY KEY(id, uname),
						FOREIGN KEY(uname) REFERENCES userspb(uname) ON DELETE CASCADE
					);`

	_, err = db.Exec(photobQuery)

	return err
}
