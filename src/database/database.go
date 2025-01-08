package database

import (
	"database/sql"
	"embed"
	"os"

	"test/src/tools/log"
	"test/src/tools/path"

	_ "github.com/mattn/go-sqlite3"
)

const (
	databasesPath  = "databases"
	migrationsPath = "migrations"
)

//go:embed migrations/*
var migrations embed.FS

func DatabasePath(database string) string {
	databasesPath := path.StorageDir() + "\\" + databasesPath

	err := path.MkDir(databasesPath)

	if err != nil {
		log.Fatal(err)
	}

	return databasesPath + "\\" + database + ".sqlite"
}

func Open(database string, migration bool) *sql.DB {
	databasePath := DatabasePath(database)

	db, err := sql.Open("sqlite3", databasePath)

	if err != nil {
		log.Fatal(err)
	}

	if migration {
		runMigration(db, database)
	}

	return db
}

func migrateDB(db *sql.DB, file *os.File, database string) {
	stat, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	if stat.Size() == 0 {
		migrate(db, database)
	}
}

func runMigration(db *sql.DB, database string) {
	dbPath := DatabasePath(database)

	if path.FileExits(dbPath) {
		file, err := os.Open(dbPath)

		if err != nil {
			log.Fatal(err)
		}

		migrateDB(db, file, database)

		return
	}

	file, err := os.Create(dbPath)

	if err != nil {
		log.Fatal(err)
	}

	migrateDB(db, file, database)
}

func migrate(db *sql.DB, database string) {
	query, err := migrations.ReadFile(migrationsPath + "/" + database + ".sql")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(query))

	if err != nil {
		log.Fatal(err)
	}
}
