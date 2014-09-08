package initializer

import (
	"database/sql"
	"fmt"
	"os"
)

func CreateStories(db *sql.DB) {
	row := db.QueryRow("SHOW TABLES LIKE 'stories'")
	var result string
	row.Scan(&result)

	if result == "stories" {
		return
	}

	tableDefinition := "CREATE TABLE `stories` ( " +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`title` varchar(255) DEFAULT NULL," +
		"`body` longtext," +
		"`created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP," +
		"`modified` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP," +
		"`state` INT DEFAULT 1, " +
		"FOREIGN KEY (state) REFERENCES states(id) ON DELETE CASCADE," +
		"PRIMARY KEY (`id`)" +
		")"

	_, err := db.Exec(tableDefinition)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating table stories")
		os.Exit(1)
	}

	fmt.Println("Created table stories")
}

func CreateStates(db *sql.DB) {
	row := db.QueryRow("SHOW TABLES LIKE 'states'")
	var result string
	row.Scan(&result)

	if result == "states" {
		return
	}

	tableDefinition := "CREATE TABLE `states` (" +
		"`id` int(11) NOT NULL AUTO_INCREMENT," +
		"`name` varchar(255) DEFAULT NULL," +
		"`order` int(11) NOT NULL," +
		"PRIMARY KEY (`id`)" +
		");" // +

	_, err := db.Exec(tableDefinition)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error creating table states")
		os.Exit(1)
	}

	fmt.Println("Created table states")

	fillStates(db)
}

func fillStates(db *sql.DB) {
	tableDefinition := "INSERT INTO `states` (name, `order`) VALUES ('pending', 0), ('progress', 1), ('completed', 2);"

	_, err := db.Exec(tableDefinition)

	if err != nil {
		fmt.Println(err)
		fmt.Println("Error inserting into table states")
		os.Exit(1)
	}

	fmt.Println("Initialized states table")
}

func CreateTables(db *sql.DB) {
	CreateStates(db)
	CreateStories(db)
}
