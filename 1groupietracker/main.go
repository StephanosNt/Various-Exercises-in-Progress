package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Structures to parse the API data
type Artist struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	YearFormed int      `json:"year_formed"`
	FirstAlbum string   `json:"first_album"`
	Members    []string `json:"members"`
}

type Location struct {
	Locations []string `json:"locations"`
}

type Date struct {
	Dates []string `json:"dates"`
}

type Relation struct {
	Relations map[string][]string `json:"relations"`
}

// Variables to store parsed data
var artists []Artist
var locations map[string][]string
var dates map[string]Date
var relations Relation

// Load data from API
func fetchData(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(target); err != nil {
		return err
	}
	return nil
}

// Handler for the homepage
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, artists)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
	}
}

// Handler for artist details
func artistHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	for _, artist := range artists {
		if artist.Name == name {
			tmpl, err := template.ParseFiles("templates/artist.html")
			if err != nil {
				http.Error(w, "Error loading template", http.StatusInternalServerError)
				return
			}
			artistDetails := struct {
				Artist   Artist
				Location Location
				Dates    Date
			}{
				Artist:   artist,
				Location: Location{Locations: locations[artist.Name]},
				Dates:    dates[artist.Name],
			}
			if err := tmpl.Execute(w, artistDetails); err != nil {
				http.Error(w, "Error rendering template", http.StatusInternalServerError)
			}
			return
		}
	}
	http.NotFound(w, r)
}
func main() {
	// Load data from the API
	if err := fetchData("https://groupietrackers.herokuapp.com/api/artists", &artists); err != nil {
		log.Fatalf("Error fetching artists: %v", err)
	}
	if err := fetchData("https://groupietrackers.herokuapp.com/api/locations", &locations); err != nil {
		log.Fatalf("Error fetching locations: %v", err)
	}
	if err := fetchData("https://groupietrackers.herokuapp.com/api/dates", &dates); err != nil {
		log.Fatalf("Error fetching dates: %v", err)
	}
	if err := fetchData("https://groupietrackers.herokuapp.com/api/relations", &relations); err != nil {
		log.Fatalf("Error fetching relations: %v", err)
	}

	// Serve static files (CSS, JS, Images)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Setup routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/artist", artistHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server started at http://localhost:%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
