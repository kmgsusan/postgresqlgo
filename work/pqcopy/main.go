package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	// Initialize connection constants.
	HOST     = "192.168.123.213"
	PORT     = 9999
	DATABASE = "mall"
	USER     = "dba"
	PASSWORD = "@cafe24"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DATABASE)

	// Initialize connection object.
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	fmt.Println("Successfully created connection to database")

	sql_statement := "SELECT nspname from pg_namespace where nspname ~ '^ec_';"
	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()
	defer db.Close()

	var nspname string
	for rows.Next() {
		switch err := rows.Scan(&nspname); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			db.Exec("set search_path to " + nspname)
			rows2, _ := db.Query("SELECT mall_id, current_Schema() as sn from mall ")
			fmt.Println(rows2)
			rows2.Close()

		default:
			checkError(err)
		}
	}
}
