package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Story struct {
	Title string
	Body  string
}

func save(p *Story) error {
	return errors.New("Not Implemented")
}

func loadStory(id string) (*Story, error) {
	db := getDB()

	var title string
	var body string
	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	err := db.QueryRow("select title, body from story where id = ?", id).Scan(&title, &body)

	if err != nil {
		log.Fatal(err)
	}

	return &Story{Title: title, Body: body}, nil
}

func editStory(id string) (*Story, error) {
	return nil, errors.New("editStory not implemented")
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	// id := r.URL.Path[len("/view/"):]
	story := new(Story)
	err := json.NewDecoder(r.Body).Decode(story)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "title: %s, body: %s", story.Title, story.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/view/"):]
	s, _ := loadStory(id)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", s.Title, s.Body)
}

var db *sql.DB

func getDB() *sql.DB {
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
	// defer db.Close()
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
