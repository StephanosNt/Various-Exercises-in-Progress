package main

import (
	"net/http"
)

func GetArtists(w http.ResponseWriter, r *http.Request) {
	data, err := FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, "Failed to fetch artists", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func GetArtist(w http.ResponseWriter, r *http.Request) {
	artistID := r.URL.Query().Get("id")
	if artistID == "" {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	data, err := FetchDataFromAPI("https://groupietrackers.herokuapp.com/api/artists/" + artistID)
	if err != nil {
		http.Error(w, "Failed to fetch artist details", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
