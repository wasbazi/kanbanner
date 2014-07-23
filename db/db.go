package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Story struct {
	Title string
	Body  string
}

func LoadStory(id string) (*Story, error) {
	db := GetDB()

	var title string
	var body string
	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	err := db.QueryRow("select title, body from story where id = ?", id).Scan(&title, &body)

	if err != nil {
		log.Fatal(err)
	}

	return &Story{Title: title, Body: body}, nil
}

func EditStory(id string) (*Story, error) {
	story := new(Story)
	db := GetDB()
	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	_, err := db.Exec("update story set title = ?, body = ? where id = ?", story.Title, story.Body, id)

	if err != nil {
		log.Fatal(err)
	}

	return story, nil
}

var db *sql.DB

func GetDB() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hello")

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		// do something here
	}

	return db
}
