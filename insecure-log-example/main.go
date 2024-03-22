// main.go

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Vulnerable code: logging sensitive information
		log.Printf("Login attempt: username=%s, password=%s\n", username, password)

		if username == "admin" && password == "admin123" {
			w.Write([]byte("Login successful"))
		} else {
			w.Write([]byte("Login failed"))
		}
	})

	http.ListenAndServe(":9080", nil)
}

