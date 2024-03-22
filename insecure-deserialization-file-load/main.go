// main.go

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Config struct {
	Filename string `json:"filename"`
}

func main() {
	http.HandleFunc("/load", func(w http.ResponseWriter, r *http.Request) {
		var config Config

		// Vulnerable code: directly deserializing untrusted JSON data
		err := json.NewDecoder(r.Body).Decode(&config)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		// Open the specified file
		file, err := os.Open(config.Filename)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error opening file: %v", err), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Serve the file
		_, err = io.Copy(w, file)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error serving file: %v", err), http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":9080", nil)
}

