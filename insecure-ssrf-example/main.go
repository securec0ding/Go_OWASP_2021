// main.go

package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        url := r.URL.Query().Get("url")

        // Vulnerable code: making an HTTP request without proper validation
        resp, err := http.Get(url)
        if err != nil {
            http.Error(w, "Error fetching URL", http.StatusInternalServerError)
            return
        }
        defer resp.Body.Close()

        body, err := ioutil.ReadAll(resp.Body)
        if err != nil {
            http.Error(w, "Error reading response body", http.StatusInternalServerError)
            return
        }

        fmt.Fprintf(w, "Response from %s:\n\n%s", url, body)
    })

    http.ListenAndServe(":9080", nil)
}

