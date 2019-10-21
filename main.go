package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetHomeEndpoint(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "Ez api")
	if err != nil {
		log.Fatal("Error in GetHomeEndpoint")
	}
}

func main() {
	log.Print("Init server.")

	router := mux.NewRouter()
	router.HandleFunc("/", GetHomeEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", router))
}
