package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() {

	// Database string when deploying project using docker
	// Comment the code below when deploying project locally
	//databaseString := "user=" + os.Getenv("POSTGRES_USER") + " dbname=" + os.Getenv("POSTGRES_DB") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("POSTGRES_HOST") + " sslmode=disable"

	// Database string when deploying project locally
	// Uncomment the code below when deploying project locally
	databaseString := "user=" + os.Getenv("POSTGRES_USER") + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("POSTGRES_HOST") + " sslmode=disable"

	db, err = sql.Open("postgres", databaseString)
	if err != nil {
		fmt.Println("While connecting to database, got error: ", err)
		os.Exit(1)
	}

	// Uncomment the code below when deploying locally
	db = CreateDatabase(db)
	if err != nil {
		fmt.Println("While creating database, got error: ", err)
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

func CreateDatabase(db *sql.DB) *sql.DB {
	pgDB := os.Getenv("POSTGRES_DB")
	var newDB *sql.DB

	_, err := db.Exec("CREATE DATABASE " + pgDB)
	if err != nil {
		db.Close()
		databaseString := "user=" + os.Getenv("POSTGRES_USER") + " dbname=" + pgDB + " password=" + os.Getenv("POSTGRES_PASSWORD") + " host=" + os.Getenv("POSTGRES_HOST") + " sslmode=disable"
		newDB, err = sql.Open("postgres", databaseString)
		if err != nil {
			fmt.Println("While connecting to new database, got error: ", err)
			os.Exit(1)
		}
		return newDB
	}
	return db
}
