package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler function from your original code
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello from Go AAAAA!</h1>")
}

func main() {
	// Register the handler function for all paths
	http.HandleFunc("/", handler)

	// Start the server on port 3000
	fmt.Println("Server starting on http://localhost:3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
