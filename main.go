package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var manager *Manager

func main() {
	log.Print("Init server.")

	//init env
	if err := os.Mkdir("tmp", os.ModePerm); err != nil {
		if _, err = os.Stat("tmp"); os.IsNotExist(err) {
			log.Print("Error setting up tmp folder. Error is: ", err.Error())
		}
	}

	manager = NewManager()
	graph := NewGraph()
	db := make(map[string]*User)
	go manager.manage(db, graph)

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHomeEndpoint).Methods("GET")
	router.HandleFunc("/api/register", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/api/relation", RelationEndpoint).Methods("POST")
	router.HandleFunc("/api/log", LogEndpoint)
	router.HandleFunc("/api/json", WriteJsonEndpoint)
	router.HandleFunc("/api/load", LoadJsonEndpoint)
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":3004", router))
}
