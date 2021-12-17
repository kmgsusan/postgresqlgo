package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jackc/pgx"
)

func main() {
	// var PgConn ConnConfig
	// PgConn = ConnConfig{Host: "192.168.123.213", Port: 9999, Database: "mall", User: "dba", Password: "@cafe24"}
	// conn, err := pgx.Connect(PgConn)
	databaseUrl := "postgres://dba:@cafe24@192.168.123.213:9999/mall"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// rows, err := conn.Query(context.Background(), "copy(select *, null::text from mono_ec.misc_code) to stdout delimiter '\t'")
	// rows, err := conn.Query(context.Background(), "select *, null::text as txt, ''::text as val from mono_ec.misc_code")
	// if err != nil {
	// 	log.Fatal("error while executing query")
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	values, err := rows.Values()
	// 	if err != nil {
	// 		log.Fatal("error while ite")
	// 	}
	// 	fmt.Println(values)
	// }
	// data, _ := conn.Exec(context.Background(), "copy(select *, null::text from mono_ec.misc_code) to stdout csv header")
	// fmt.Println(data)

	var output io.Reader
	conn.PgConn().CopyFrom(context.Background(), output, "copyselect *, null::text from mono_ec.misc_code")
	fmt.Println(output)

}
