// main.go

package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"

    _ "github.com/mattn/go-sqlite3"
)

func main() {
    // Initialize SQLite database
    db, err := sql.Open("sqlite3", "test.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Create table if not exists
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, username TEXT, password TEXT)")
    if err != nil {
        log.Fatal(err)
    }

    // Insert some sample data
    _, err = db.Exec(`INSERT INTO users (username, password) VALUES ('admin', 'admin123')`)
    if err != nil {
        log.Fatal(err)
    }

    http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        username := r.FormValue("username")
        password := r.FormValue("password")

        // Vulnerable code: directly interpolating user input into SQL query
        query := fmt.Sprintf("SELECT * FROM users WHERE username='%s' AND password='%s'", username, password)
        rows, err := db.Query(query)
        if err != nil {
            http.Error(w, "Error querying database", http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var found bool
        for rows.Next() {
            found = true
            break
        }

        if found {
            fmt.Fprintf(w, "Login successful for user: %s", username)
        } else {
            fmt.Fprintf(w, "Login failed")
        }
    })

    log.Fatal(http.ListenAndServe(":4080", nil))
}

