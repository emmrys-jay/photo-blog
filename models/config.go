package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() {

	// Connect to remote database
	// Edit details when connecting to your PostgreSQL server instance
	databaseString := "user=" + os.Getenv("POSTGRES_USER") + " dbname=" + os.Getenv("POSTGRES_DB") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("DATABASE_HOST") + " sslmode=disable"
	db, err = sql.Open("postgres", databaseString)
	if err != nil {
		fmt.Println("While connecting to database, got error: ", err)
		os.Exit(1)
	}

	// Ping database to see if connection has been successfully established
	err = db.Ping()
	if err != nil {
		fmt.Println("While pinging database, got error: ", err)
		os.Exit(1)
	}

	err = CreateTables(db)
	if err != nil {
		fmt.Println("While creating database schemas, got error: ", err)
		os.Exit(1)
	}

	fmt.Println("you connected to database successfully")
}

func CreateTables(db *sql.DB) error {
	userspbQuery := `CREATE TABLE IF NOT EXISTS userspb (
						uname VARCHAR PRIMARY KEY NOT NULL,
						email VARCHAR UNIQUE NOT NULL,
						psword VARCHAR
					);`
	_, err := db.Exec(userspbQuery)
	if err != nil {
		return err
	}

	photobQuery := `CREATE TABLE IF NOT EXISTS photob (
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
