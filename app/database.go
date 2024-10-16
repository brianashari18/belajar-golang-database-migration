package app

import (
	"database/sql"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/golang_database_migrations")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

// migrate create -ext sql -dir db/migrations create_table_category
// migrate -database "mysql://root:Pokemon18*@tcp(localhost:3306)/golang_database_migrations" -path db/migrations up
// migrate -database "mysql://root:Pokemon18*@tcp(localhost:3306)/golang_database_migrations" -path db/migrations down
