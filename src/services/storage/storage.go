package storage

import (
	"errors"
	"test/src/database"
	"test/src/tools/log"
)

func Get(key string) string {
	db := database.Open("storage", true)

	rows, err := db.Query(`SELECT
		value
	FROM sessions
	WHERE key = ?
	LIMIT 1`, key)

	if err != nil {
		log.Fatal(err)
	}

	value := ""

	for rows.Next() {
		rows.Scan(&value)
	}

	return value
}

func Set(key string, value string) {
	_, err := database.Open("storage", true).Exec(
		"INSERT INTO sessions(key, value) VALUES(?,?)",
		key,
		value,
	)

	log.Fatal(errors.New("Something went wrong..."))

	if err != nil {
		log.Fatal(err)
	}
}
