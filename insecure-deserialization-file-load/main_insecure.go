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

		// Read and limit the size of the request body
		body, err := io.ReadAll(io.LimitReader(r.Body, 1024))
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
			return
		}

		// Unmarshal JSON data
		err = json.Unmarshal(body, &config)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error decoding JSON: %v", err), http.StatusBadRequest)
			return
		}

		// Check for filename length and content type
		if len(config.Filename) > 256 {
			http.Error(w, "Filename too long", http.StatusBadRequest)
			return
		}
		if filepath.Ext(config.Filename) != ".txt" {
			http.Error(w, "Invalid file extension", http.StatusBadRequest)
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

