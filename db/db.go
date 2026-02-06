package db

// go provides the official sql package but the database specific driver it does not provided
// we need to use third party package for the specific db
import (
	"database/sql"

	// this 3rd party package is imported and will be used by go's sql package, hence we use _ to indicate that
	// it will be imported but no used by us
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	// driver name and datasource..since sqlite used disk storage we have provided filename where data will be stored
	DB, err = sql.Open("sqlite3", "api.db")
	// DB.
	if err != nil {
		panic("Could not connect to DB")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text not null,
	email text not null unique,
	password text not null
)
	`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table")
	}

	createEventsTable :=
		`
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text not null,
	description text not null,
	location text not null,
	dateTime DATETIME not null,
	user_id integer ,
	foreign key(user_id) references users(id)
	)`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}

	createRegistrationsTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id integer not null,
	event_id integer not null,
	foreign key(user_id) references users(id),
	foreign key(event_id) references events(id)
	)`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic("Could not create registrations table")
	}
}
