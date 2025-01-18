package main

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/api/artists", GetArtists)
	mux.HandleFunc("/api/artist", GetArtist)

	return mux
}
