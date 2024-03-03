package main

import (
    "fmt"
    "log"
    "example.com/AssetManagement/internal/db"
)

func main() {
    // Establish a connection to the database
    conn, err := db.makeConnection()
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    queries := []string{
        `CREATE TABLE IF NOT EXISTS targets (
            id INTEGER PRIMARY KEY,
            name TEXT UNIQUE
        )`,
        `CREATE TABLE IF NOT EXISTS subdomains (
            id INTEGER PRIMARY KEY,
            domain TEXT,
            target_id INTEGER,
            FOREIGN KEY (target_id) REFERENCES targets(id)
        )`,
        `CREATE TABLE IF NOT EXISTS ip_addresses (
            id INTEGER PRIMARY KEY,
            subdomain_id INTEGER,
            ip_address TEXT,
            FOREIGN KEY (subdomain_id) REFERENCES subdomains(id)
        )`,
        `CREATE TABLE IF NOT EXISTS http_responses (
            id INTEGER PRIMARY KEY,
            subdomain_id INTEGER,
            status_code INTEGER,
            content_length INTEGER,
            page_title TEXT,
            redirect_location TEXT,
            FOREIGN KEY (subdomain_id) REFERENCES subdomains(id)
        )`,
        `CREATE TABLE IF NOT EXISTS screenshots (
            id INTEGER PRIMARY KEY,
            subdomain_id INTEGER,
            screenshot_data BLOB,
            FOREIGN KEY (subdomain_id) REFERENCES subdomains(id)
        )`,
        `CREATE TABLE IF NOT EXISTS response_files (
            id INTEGER PRIMARY KEY,
            subdomain_id INTEGER,
            file_path TEXT,
            FOREIGN KEY (subdomain_id) REFERENCES subdomains(id)
        )`,
    }

    for _, query := range queries {
        err := db.execQuery(conn, query)
        if err != nil {
            log.Fatal(err)
        }
    }

    // Execute SELECT query
    rows, err := db.selectQuery(conn, "SELECT * FROM targets")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // Iterate over the result set
    for rows.Next() {
        var id int
        var name string
        if err := rows.Scan(&id, &name); err != nil {
            log.Fatal(err)
        }
        fmt.Println(id, name)
    }
    if err := rows.Err(); err != nil {
        log.Fatal(err)
    }
}

