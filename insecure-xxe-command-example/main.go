// main.go

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		// Vulnerable code: directly parsing XML without disabling external entities
		data := string(body)
		if strings.Contains(data, "<!DOCTYPE") {
			http.Error(w, "XML parsing error", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Received XML: %s", data)
	})

	http.ListenAndServe(":9080", nil)
}

