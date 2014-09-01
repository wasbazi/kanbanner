# Kanbanner

This is my attempt at a simple and usable Kanban application

# Dependencies

* Go 1.3
* MySQL server running

# Installation

To install & run:

```shell
  godep restore
  go run kanbanner.go
```

This will setup the necessary tables in the MySQL database

# API Endpoints

* url: `localhost:8090/stories`, request: `GET`
  * returns an object with all stories from the db, where keys are states
* url: `localhost:8090/story/:id`, request: `GET`
  * returns the story with `:id` from the db
* url: `localhost:8090/story/:id`, request: `POST`
  * update the story with the passed along JSON object, updates all fields

* url: `localhost:8090/states`, request: `GET`
  * returns an array of all states from the db
