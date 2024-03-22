// main.go

package main

import (
    "html/template"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        userInput := r.URL.Query().Get("input")
        
        // Secure code: using html/template to safely render user input
        tmpl := template.Must(template.New("index").Parse(`<h1>User Input: {{.}}</h1>`))
        tmpl.Execute(w, userInput)
    })

    http.ListenAndServe(":9080", nil)
}

