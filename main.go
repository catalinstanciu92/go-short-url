// add http server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Missing query parameter", http.StatusBadRequest)
			return
		}

		// get the original URL from the database
		var originalURL string
		originalURL, err := getOriginalURL(query)

		if err != nil {
			http.Error(w, "URL not found", http.StatusNotFound)
			fmt.Println("Error querying database:", err)
			return
		}

		// if the request is a POST, shorten the URL
		if r.Method == http.MethodPost {
			fmt.Fprintf(w, "Shortened URL for %q", originalURL)
			return
		}

		//redirect to the original URL
		http.Redirect(w, r, originalURL, http.StatusFound)
	})

	http.HandleFunc("/shorten", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "Missing query parameter", http.StatusBadRequest)
			return
		}
		shortenedURL := generateShortURL(query)

		// respond with json
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"shortened_url": "%s"}`, shortenedURL)
	})

	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :3009")
	if err := http.ListenAndServe(":3009", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
