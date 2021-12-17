package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/jackc/pgx"
)

func main() {
	// pgxConConfig := pgx.ConnConfig{
	// 	Port:     5432,
	// 	Host:     "remote_host",
	// 	Database: "db_name",
	// 	User:     "my_user",
	// 	Password: "my_password",
	// }

	// conn, err := pgx.Connect(pgxConConfig)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	databaseUrl := "postgres://dba:@cafe24@192.168.123.213:9999/mall"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil {
		panic(err)
	}

	res, err := conn.CopyTo(f, fmt.Sprintf("COPY mono_ec.%s FROM STDIN DELIMITER '|' CSV HEADER", table))

	// tables := []string{"misc_code"}

	// import_dir := "/dir_to_import_from"
	// export_dir := "/dir_to_export_to"

	// for _, t := range tables {
	// f, err := os.OpenFile(fmt.Sprintf("%s/table_%s.csv", import_dir, t), os.O_RDONLY, 0777)
	// if err != nil {
	// 	return
	// }
	// f.Close()

	// err = importer(conn, f, t)
	// if err != nil {
	// 	break
	// }

	// fmt.Println("  Done with import and doing export")
	// ef, err := os.OpenFile(fmt.Sprintf("%s/table_%s.csv", export_dir, t), os.O_CREATE|os.O_WRONLY, 0777)
	// if err != nil {
	// 	fmt.Println("error opening file:", err)
	// 	return
	// }
	// ef.Close()

	// err = exporter(conn, ef, t)
	// if err != nil {
	// 	break
	// }
	// }
}

// func importer(conn *pgx.Conn, f *os.File, table string) error {
// 	res, err := conn.CopyFrom(f, fmt.Sprintf("COPY mono_ec.%s FROM STDIN DELIMITER '|' CSV HEADER", table))
// 	if err != nil {
// 		return err
// 	}
// 	fmt.Println("==> import rows affected:", res.RowsAffected())

// 	return nil
// }

func exporter(conn *pgx.Conn, f *os.File, table string) error {
	res, err := conn.CopyTo(f, fmt.Sprintf("COPY mono_ec.%s TO STDOUT DELIMITER '|' CSV HEADER", table))
	if err != nil {
		return fmt.Errorf("error exporting file: %+v", err)
	}
	fmt.Println("==> export rows affected:", res.RowsAffected())
	return nil
}
