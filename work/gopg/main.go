package main

import (
	"os"

	"github.com/go-pg/pg"
)

func main() {
	db := pg.Connect(&pg.Options{
		User:     "dba",
		Password: "@cafe24",
		Database: "mall",
		Addr:     "192.168.123.213:9999", //database address
	})
	defer db.Close()

	//open output file
	out, err := os.Create("bar.json")
	if err != nil {
		panic(err)
	}
	defer out.Close()

	//execute opy
	copy := `COPY (SELECT row_to_json(foo) FROM (select * from mono_ec.misc_code) foo ) TO STDOUT`
	_, err = db.CopyTo(out, copy)
	if err != nil {
		panic(err)
	}
}
