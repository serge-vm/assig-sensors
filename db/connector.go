package db

import (
	"database/sql"
	"embed"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var migrations embed.FS

var dbConn *sql.DB

func NewConnector() *sql.DB {
	db, err := sql.Open("sqlite3", "sensors.db")
	if err != nil {
		log.Fatal(err)
	}
	applyMigrations(db)
	return db
}

func GetDB() *sql.DB {
	if dbConn == nil {
		dbConn = NewConnector()
	}
	return dbConn
}

func applyMigrations(db *sql.DB) {
	goose.SetDialect("sqlite3")
	goose.SetBaseFS(migrations)

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
	if err := goose.Version(db, "migrations"); err != nil {
		log.Fatal(err)
	}
}
