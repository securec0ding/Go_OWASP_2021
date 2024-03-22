// main.go

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/download", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Query().Get("url")

		// Vulnerable code: downloading code without integrity check
		resp, err := http.Get(url)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error downloading file: %v", err), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Create a new file
		file, err := os.Create("downloaded_file")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error creating file: %v", err), http.StatusInternalServerError)
			return
		}
		defer file.Close()

		// Write the downloaded content to the file
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error writing to file: %v", err), http.StatusInternalServerError)
			return
		}

		w.Write([]byte("File downloaded successfully"))
	})

	http.ListenAndServe(":9080", nil)
}

