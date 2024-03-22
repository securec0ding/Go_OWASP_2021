// main.go

package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        // Vulnerable code: processing form data without CSRF protection
        if r.Method == http.MethodPost {
            r.ParseForm()
            username := r.Form.Get("username")
            fmt.Fprintf(w, "Hello, %s!", username)
        } else {
            // Render a simple HTML form
            fmt.Fprintf(w, `
                <form action="/" method="post">
                    <input type="text" name="username" placeholder="Enter your username">
                    <input type="submit" value="Submit">
                </form>
            `)
        }
    })

    http.ListenAndServe(":9080", nil)
}

