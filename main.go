package main

import (
	"github.com/mmatur/vercel-multi-headers/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.Handler)
	_ = http.ListenAndServe(":8000", mux)
}
