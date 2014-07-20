package main

import (
 "encoding/json"
 "fmt"
 "log"
 "io/ioutil"
 "net/http"
)

type Story struct {
  Title string
  Body []byte
}

func (p *Story) save() error {
    filename := p.Title + ".txt"
    return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadStory(title string) (*Story, error) {
    filename := title + ".txt"
    body, err := ioutil.ReadFile(filename)

    if err != nil {
     return nil, err
    }

    return &Story{Title: title, Body: body}, nil
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	story := new(Story)
	err := json.NewDecoder(r.Body).Decode(story)
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintf(w, "title: %s, body: %s", story.Title, story.Body)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/view/"):]
    s, _ := loadStory(title)
    fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", s.Title, s.Body)
}

func main() {
    http.HandleFunc("/view/", viewHandler)
    http.HandleFunc("/edit/", editHandler)
    // http.HandleFunc("/save/", saveHandler)
    http.ListenAndServe(":8080", nil)
}
