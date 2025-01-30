package main

import (
	"fmt"
	"golang-vercel/api"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)
	fmt.Println("Server running on http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
