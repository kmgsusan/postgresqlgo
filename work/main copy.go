package main

import (
	"database/sql"
	"fmt"

	"github.com/joho/sqltocsv"
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
		panic(err)
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

	// Drop previous table of same name if one exists.
	_, err = db.Exec("DROP TABLE IF EXISTS inventory;")
	checkError(err)
	fmt.Println("Finished dropping table (if existed)")

	// Create table.
	_, err = db.Exec("CREATE TABLE inventory (id serial PRIMARY KEY, name VARCHAR(50), quantity INTEGER);")
	checkError(err)
	fmt.Println("Finished creating table")

	// Insert some data into table.
	sql_statement := "INSERT INTO inventory (name, quantity) VALUES ($1, $2);"
	_, err = db.Exec(sql_statement, "banana", 150)
	checkError(err)
	_, err = db.Exec(sql_statement, "orange", 154)
	checkError(err)
	_, err = db.Exec(sql_statement, "apple", 100)
	checkError(err)
	fmt.Println("Inserted 3 rows of data")

	// read rows from table
	var id int
	var name string
	var quantity int

	sql_statement = "SELECT * FROM inventory;"
	rows, err := db.Query(sql_statement)
	checkError(err)
	defer rows.Close()

	for rows.Next() {
		switch err := rows.Scan(&id, &name, &quantity); err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned")
		case nil:
			fmt.Printf("Data row = (%d, %s, %d)\n", id, name, quantity)
		default:
			checkError(err)
		}
	}

	sql_statement = "SELECT nspname from pg_namespace where nspname ~ '^ec_';"
	schemarows, err := db.Query(sql_statement)
	checkError(err)
	defer schemarows.Close()
	defer db.Close()

	var nspname string
	for schemarows.Next() {
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
