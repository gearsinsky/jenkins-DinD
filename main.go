package main

import (
	"fmt"
	"io"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Path: %s\n", r.URL.Path)
	fmt.Fprintf(w, "Headers:\n")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Fprintf(w, "  %s: %s\n", name, value)
		}
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusInternalServerError)
		return
	}

	if len(body) > 0 {
		fmt.Fprintf(w, "\nBody:\n%s\n", string(body))
	} else {
		fmt.Fprintf(w, "\n(No Body)\n")
	}
}

func main() {
	http.HandleFunc("/", echoHandler)
	fmt.Println("Server running on http://0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

