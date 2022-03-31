package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"scrape/scrape"
)

func init() {
	scrape.LOG_LEVEL = 2
}

func main() {
	
	if scrape.LOG_LEVEL > 0 {
		log.Println("starting API server")
	}
	//create a new router
	router := mux.NewRouter()

	if scrape.LOG_LEVEL > 0 {
		log.Println("creating routes")
	}

	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")		
    router.HandleFunc("/addsearch/{regex}", scrape.AddRegex).Methods("GET")
    router.HandleFunc("/clear", scrape.Clear).Methods("GET")
    router.HandleFunc("/reset", scrape.Reset).Methods("GET")



	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}