package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetHomeEndpoint(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "API is online")
	if err != nil {
		log.Fatal("Error in GetHomeEndpoint")
	}
}

func main() {
	log.Print("Init server.")

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHomeEndpoint)
	router.Handle("/", http.FileServer(http.Dir("static/")))
	log.Fatal(http.ListenAndServe(":3004", router))
}
