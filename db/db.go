package db

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	Pending   = "pending"
	Progress  = "progress"
	Completed = "completed"
)

type Story struct {
	Title    string `json:"title" bind:"required"`
	Body     string `json:"body" bind:"required"`
	Created  time.Time
	Modified time.Time
	// State    State
}

func LoadStories() ([]*Story, error) {
	db := GetDB()

	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	rows, err := db.Query("select title, body from story")

	if err != nil {
		log.Fatal(err)
	}

	stories := make([]*Story, 0, 0)

	defer rows.Close()
	for rows.Next() {
		var title string
		var body string

		if err := rows.Scan(&title, &body); err != nil {
			log.Fatal(err)
		}
		story := &Story{Title: title, Body: body}
		stories = append(stories, story)
	}

	return stories, nil
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

func EditStory(id string, story Story) (Story, error) {
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
