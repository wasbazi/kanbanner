package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/wasbazi/kanbanner/db"
)

func editHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/edit/"):]

	story := new(db.Story)
	err := json.NewDecoder(r.Body).Decode(story)
	if err != nil {
		log.Println(err)
	}

	db.EditStory(id)
	story, err = db.LoadStory(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", story.Title, story.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/view/"):]
	s, _ := db.LoadStory(id)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", s.Title, s.Body)
}

func AcceptConnections() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	// http.HandleFunc("/save/", saveHandler)
	http.ListenAndServe(":8080", nil)
}
