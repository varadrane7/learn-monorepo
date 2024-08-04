package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
	"url-shortener/shortener"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

var c *cache.Cache

func main() {
	// Initialize cache
	c = cache.New(5*time.Minute, 10*time.Minute)

	// Create a new router
	r := mux.NewRouter()

	// Define the handlers
	r.HandleFunc("/", homeHandler).Methods("GET", "POST")
	r.HandleFunc("/{shortURL}", redirectHandler).Methods("GET")

	// Start the server
	log.Fatal(http.ListenAndServe(":8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()
		originalURL := r.FormValue("url")
		shortURL := shortener.GenerateShortURL()

		for {
			if _, exists := shortener.GetOriginalURL(shortURL); !exists {
				break
			}
			shortURL = shortener.GenerateShortURL()
		}

		shortener.StoreURL(shortURL, originalURL)
		c.Set(shortURL, originalURL, cache.DefaultExpiration)

		tmpl, _ := template.ParseFiles("templates/index.html")
		tmpl.Execute(w, map[string]string{"ShortURL": shortURL})
		return
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortURL := vars["shortURL"]

	originalURL, found := c.Get(shortURL)
	if found {
		// Assert originalURL to string
		if originalURLStr, ok := originalURL.(string); ok {
			http.Redirect(w, r, originalURLStr, http.StatusMovedPermanently)
			return
		}
	}

	originalURL, exists := shortener.GetOriginalURL(shortURL)
	if !exists {
		http.NotFound(w, r)
		return
	}

	c.Set(shortURL, originalURL, cache.DefaultExpiration)
	http.Redirect(w, r, originalURL.(string), http.StatusMovedPermanently)
}
