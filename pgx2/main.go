// https://future-architect.github.io/articles/20210727a/

package main

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/jackc/pgx"
)

type copyFromSource struct {
	r       *csv.Reader
	nextRow []interface{}
	err     error
}

func (s *copyFromSource) Next() bool {
	s.nextRow = nil
	s.err = nil
	record, err := s.r.Read()
	if err == io.EOF {
		return false
	} else if err != nil {
		s.err = err
		return false
	}

	s.nextRow = []interface{}{
		record[0], record[1],
		record[2] != "", record[3] != "", record[4] != "", record[5] != "",
		record[6] != "", record[7] != "", record[8] != "", record[9] != "",
	}
	return true
}

func (s copyFromSource) Values() ([]interface{}, error) {
	if s.err != nil {
		return nil, s.err
	}
	return s.nextRow, nil
}

func (s copyFromSource) Err() error {
	return s.err
}

var _ pgx.CopyFromSource = &copyFromSource{}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	// f, err := os.Open("output.csv")
	// if err != nil {
	// 	log.Fatal("🐙", err)
	// }
	// r := csv.NewReader(f)
	// r.FieldsPerRecord = -1

	conn, err := pgx.Connect(context.Background(), "postgres://dba:@cafe24@192.168.123.213:9999/mall")
	if err != nil {
		log.Fatal("🦑", err)
	}

	txn, err := conn.Begin(ctx)
	if err != nil {
		log.Fatal("🐣", err)
	}
	_, err = txn.CopyFrom(ctx, pgx.Identifier{"allergies"}, []string{
		"category", "menu",
		"shrimp", "crab", "wheat", "soba", "eggs", "milk", "peanuts", "walnuts",
	}, &copyFromSource{r: r})

	if err != nil {
		log.Fatal("🐬", err)
	}

	err = txn.Commit(ctx)
	if err != nil {
		log.Fatal("🐱", err)
	}
}
