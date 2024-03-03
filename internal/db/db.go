package db

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

func MakeConnection() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "db.sqlite3")
    if err != nil {
        log.Fatal(err)
        return nil, err
    }
    return db, nil
}

func ExecQuery(db *sql.DB, query string) error {
    _, err := db.Exec(query)
    return err
}

func SelectQuery(db *sql.DB, query string) (*sql.Rows, error) {
    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    return rows, nil
}

