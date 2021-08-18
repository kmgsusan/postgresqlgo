package main

import (
	"database/sql"
	"fmt"
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

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected")

}

// CheckError func
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
