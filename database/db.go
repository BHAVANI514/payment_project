package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// func Connect() {
// 	connStr := `host=10.10.19.138 username=postgres password=root port=5433 dbname=bank sslmode=disable`
// 	var err error
// 	DB, err = sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatalf("Failed to open database: %v", err)
// 	}

// 	// Optional but recommended: test the connection
// 	if err = DB.Ping(); err != nil {
// 		log.Fatalf("Failed to connect to DB: %v", err)
// 	}
// }

func Connect() {
	var err error
	DB, err = sql.Open("postgres", "host=localhost port=5433 user=postgres password=root dbname=bank sslmode=disable")

	if err != nil {
		log.Fatalf("error opening DB: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("connection failed to postgres: %+v", err)
	} else {
		fmt.Println("connection success to postgres")
	}
}
