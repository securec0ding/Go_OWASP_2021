package main

import (
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        userInput := r.URL.Query().Get("input")
        
        // Non-compliant code: rendering user input directly in the HTML response
        html := "<h1>User Input: " + userInput + "</h1>"
        w.Write([]byte(html))
    })

    http.ListenAndServe(":9080", nil)
}

