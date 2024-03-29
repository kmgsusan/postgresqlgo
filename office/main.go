package main

import (
	"database/sql"
	"fmt"

	"github.com/joho/sqltocsv"
	_ "github.com/lib/pq"
)

const (
	host     = "192.168.123.213"
	port     = 9999
	database = "mall"
	username = "dba"
	password = "@cafe24"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Initialize connection string.
	var connectionString string = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	fmt.Println(connectionString)

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

	var nspname string
	for rows.Next() {
		switch err := rows.Scan(&nspname); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			//fmt.Printf("Data row = (%s)\n", nspname)
			fmt.Println("set search_path to " + nspname)
			db.Exec("set search_path to " + nspname)
			rows2, _ := db.Query("SELECT mall_id, current_Schema() as sn from mall ")
			fmt.Println(rows2)
			rows, _ := db.Query("select mall_id from mall")
			// fmt.Println(rows['mall_id'])
			// rows, err := db.Query("select * from mall")
			// CheckError(err)
			csvConverter := sqltocsv.New(rows)
			csvConverter.WriteFile("result_" + nspname + ".csv")
		default:
			checkError(err)
		}
	}
}
