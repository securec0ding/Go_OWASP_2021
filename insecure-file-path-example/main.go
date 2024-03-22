// main.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Query().Get("file")

		// Vulnerable code: directly using user-supplied file name without validation
		filePath := filepath.Join("/path/to/files", fileName)

		// Open the file
		file, err := os.Open(filePath)
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

