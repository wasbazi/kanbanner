package stories

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
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

func CreateStory(title, body string) (string, error) {
	conn := db.GetDB()

	result, err := conn.Exec("insert into stories (title, body) values (?, ?)", title, body)

	if err != nil {
		return "", err
	}

	id, err := result.LastInsertId()
	strId := strconv.FormatInt(id, 10)

	return strId, err
}

func DeleteStory(id string) error {
	conn := db.GetDB()

	result, err := conn.Exec("delete from stories where id = ?", id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()

	if affected != 1 || err != nil {
		return errors.New("Unexpected rows affected")
	}

	return nil
}

func LoadStories() (map[string][]*Story, error) {
	conn := db.GetDB()

	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	rows, err := conn.Query("select title, body, states.name state, stories.id from stories RIGHT JOIN states on state = states.id")

	if err != nil {
		log.Fatal(err)
	}

	stories := make(map[string][]*Story)

	defer rows.Close()
	for rows.Next() {
		var title sql.NullString
		var body sql.NullString
		var state sql.NullString
		var id sql.NullString

		if err := rows.Scan(&title, &body, &state, &id); err != nil {
			log.Fatal(err)
		}

		if id.Valid {
			story := &Story{Title: title.String, Body: body.String, State: state.String, Id: id.String}
			stories[story.State] = append(stories[story.State], story)
		} else {
			stories[state.String] = make([]*Story, 0, 0)
		}
	}

	return stories, nil
}

func LoadStory(id string) (*Story, error) {
	conn := db.GetDB()

	var title string
	var body string
	var state string
	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	err := conn.QueryRow("select title, body, states.name state from stories LEFT JOIN states on state = states.id where stories.id = ?", id).Scan(&title, &body, &state)

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
