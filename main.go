package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./focustime.db")

	db, err := sql.Open("sqlite3", "file:focustime.db?_foreign_keys=on")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = execSQL(db, "./sqlqueries/make_tables")
	if err != nil {
		log.Fatal(err)
	}

	err = execSQL(db, "./sqlqueries/insert_data")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done")
}

func readAllSQLFiles(path string) ([]string, error) {
	var files []string

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileInfo, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	for _, file := range fileInfo {
		data, err := os.ReadFile(path + "/" + file.Name())
		if err != nil {
			return nil, err
		}
		files = append(files, string(data))
	}
	return files, nil
}

func execSQL(db *sql.DB, path string) error {
	queries, err := readAllSQLFiles(path)
	if err != nil {
		return err
	}

	if len(queries) == 0 {
		return nil
	}
	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			log.Printf("%q: %s\n", err, query)
			return err
		}
	}
	return nil
}
