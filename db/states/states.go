package states

import (
	"log"

	"github.com/wasbazi/kanbanner/db/db"
)

type State struct {
	Id    string `json:"id" bind:"require"`
	Name  string `json:"name" bind:"required"`
	Order string `json:"order" bind:"required"`
}

func LoadStates() ([]*State, error) {
	conn := db.GetDB()

	// this may be ripe for SQL injections, ignore because SQL injections are for noobs
	rows, err := conn.Query("select id, name, `order` from states order by `order`")

	if err != nil {
		log.Fatal(err)
	}

	var states []*State

	defer rows.Close()
	for rows.Next() {
		var id string
		var name string
		var order string

		if err := rows.Scan(&id, &name, &order); err != nil {
			log.Fatal(err)
		}

		state := &State{Id: id, Name: name, Order: order}
		states = append(states, state)
	}

	return states, nil
}
