package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Print("Init server.")

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHomeEndpoint).Methods("GET")
	router.HandleFunc("/api/register", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/api/relation", RelationEndpoint).Methods("POST")
	router.Handle("/", http.FileSerover(http.Dir("static/")))
	log.Fatal(http.ListenAndServe(":3004", router))
}
