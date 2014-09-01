package stories

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wasbazi/kanbanner/db/db"
)

type Story struct {
	Id       string    `json:"id" bind:"require"`
	Title    string    `json:"title" bind:"required"`
	Body     string    `json:"body" bind:"required"`
	Created  time.Time `json:"created" bind:"require"`
	Modified time.Time `json:"modified" bind:"require"`
	State    string    `json:"state" bind:"state"`
}

func LoadStories() (map[string][]*Story, error) {
	conn := db.GetDB()

	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	rows, err := conn.Query("select title, body, state, id from stories")

	if err != nil {
		log.Fatal(err)
	}

	stories := make(map[string][]*Story)

	defer rows.Close()
	for rows.Next() {
		var title string
		var body string
		var state string
		var id string

		if err := rows.Scan(&title, &body, &state, &id); err != nil {
			log.Fatal(err)
		}
		story := &Story{Title: title, Body: body, State: state, Id: id}
		stories[story.State] = append(stories[story.State], story)
	}

	return stories, nil
}

func LoadStory(id string) (*Story, error) {
	conn := db.GetDB()

	var title string
	var body string
	var state string
	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	err := conn.QueryRow("select title, body, state from stories where id = ?", id).Scan(&title, &body, &state)

	if err != nil {
		log.Fatal(err)
	}

	return &Story{Title: title, Body: body, Id: id, State: state}, nil
}

func EditStory(id string, story Story) (*Story, error) {
	fmt.Printf("state: %s, story: %#v", story.State, story)

	conn := db.GetDB()
	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	_, err := conn.Exec("update stories set title = ?, body = ?, state = ? where id = ?", story.Title, story.Body, story.State, id)

	if err != nil {
		log.Fatal(err)
	}

	return &story, nil
}
