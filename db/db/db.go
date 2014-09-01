package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/wasbazi/kanbanner/db/initializer"
)

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hello")
	initializer.CreateTables(db)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Cannot connect to MySQL server")
		os.Exit(1)
	}

	return db
}
