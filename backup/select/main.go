package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/sqltocsv"
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
// https://godoc.org/github.com/joho/sqltocsv
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

	for _, schema := range schemanames {
		db.Exec("set search_path to " + schema)
		// output_data(&db, "select * from mall")
		rows, err := db.Query("select * from mall")
		CheckError(err)
		csvConverter := sqltocsv.New(rows)
		csvConverter.WriteFile("result_" + schema + ".csv")
	}
}

// func output_data(db *DB, sql_query string) {
// 	rows, err := db.Query(sql_query)
// 	CheckError(err)
// 	csvConverter := sqltocsv.New(rows)
// 	csvConverter.WriteFile("result.csv")
// }

// CheckError func
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
