package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Story struct {
	Title    string `json:"title" bind:"required"`
	Body     string `json:"body" bind:"required"`
	Created  time.Time
	Modified time.Time
	State    string
}

func Init(db *sql.DB) {
	row := db.QueryRow("SHOW TABLES LIKE 'story'")
	var result string
	row.Scan(&result)

	if result == "story" {
		return
	}

	tableDefinition := "CREATE TABLE `story` ( " +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`title` varchar(255) DEFAULT NULL," +
		"`body` longtext," +
		"`created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"`modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"`state` enum('pending','progress','completed') DEFAULT NULL," +
		"PRIMARY KEY (`id`)" +
		")"

	_, err := db.Exec(tableDefinition)

	if err != nil {
		fmt.Println("Error creating table story")
		os.Exit(1)
	}

	fmt.Println("Created table story")
}

func LoadStories() (map[string][]*Story, error) {
	db := GetDB()

	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	rows, err := db.Query("select title, body, state from story")

	if err != nil {
		log.Fatal(err)
	}

	stories := make(map[string][]*Story)

	defer rows.Close()
	for rows.Next() {
		var title string
		var body string
		var state string

		if err := rows.Scan(&title, &body, &state); err != nil {
			log.Fatal(err)
		}
		story := &Story{Title: title, Body: body, State: state}
		stories[story.State] = append(stories[story.State], story)
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
	Init(db)

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
