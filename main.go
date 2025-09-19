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

		if r.Method == http.MethodPost {
			fmt.Fprintf(w, "Shortened URL for %q", query)
			return
		}

		//redirect to the original URL
		http.Redirect(w, r, query, http.StatusFound)
	})

	http.HandleFunc("/", helloHandler)
	fmt.Println("Starting server on :3009")
	if err := http.ListenAndServe(":3009", nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
