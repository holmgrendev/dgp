package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	// Create a new HTTP request Multiplexer
	mux := http.NewServeMux()

	// Handle requests for static files
	mux.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))

	return mux
}
