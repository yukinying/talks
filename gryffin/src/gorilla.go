package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func loadArticleFromDatabase(w http.ResponseWriter, r *http.Request) {
	// load article from database.
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/news/{article-key}", loadArticleFromDatabase)
	http.Handle("/", r)
}
