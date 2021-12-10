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

	insert := "insert into students(name, role) values('mgkim', 1)"
	_, e := db.Exec(insert)
	CheckError(e)

}

// CheckError func
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
