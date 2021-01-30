package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dba"
	password = "@cafe24"
	dbname   = "mgkim"
)

// main
// https://golangdocs.com/golang-postgresql-example
func main() {
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	rows, err := db.Query("select nspname from pg_namespace where nspname ~ '^ec_'")
	CheckError(err)

	schemanames := []string{}
	for rows.Next() {
		var schemaname string
		rows.Scan(&schemaname)
		schemanames = append(schemanames, schemaname)
	}

	fmt.Println(schemanames)
}

// CheckError func
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
