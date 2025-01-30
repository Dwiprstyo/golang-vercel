package main

import (
	"go-vercel/api/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.Handler)

	port := "8080"
	println("Server starting on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
