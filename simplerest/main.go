package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage endpoint hit")
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	var articles []Article
	articles = []Article{
		Article{Title: "Test Title", Desc: "Test Summary", Content: "Hello World"},
		Article{Title: "Second Article", Desc: "Test Summary", Content: "Hello There"},
	}
	fmt.Fprintf(w, "all articles endpoint hit")
	json.NewEncoder(w).Encode(articles)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8088", myRouter))
}

func main() {
	handleRequests()
}
