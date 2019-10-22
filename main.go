package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var manager *Manager

func main() {
	log.Print("Init server.")

	manager = NewManager()
	graph := NewGraph()
	db := make(map[string]*User)
	go manager.manage(db, graph)

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHomeEndpoint).Methods("GET")
	router.HandleFunc("/api/register", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/api/relation", RelationEndpoint).Methods("POST")
	router.HandleFunc("/api/log", LogEndpoint)
	router.Handle("/", http.FileServer(http.Dir("static/")))
	log.Fatal(http.ListenAndServe(":3004", router))
}
