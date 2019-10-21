package main

import (
	"encoding/json"
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

type User struct {
	Name string `json:"name"`
}

func RegisterEndpoint(w http.ResponseWriter, req *http.Request) {
	var user User

	d := json.NewDecoder(req.Body)
	if err := d.Decode(&user); err != nil {
		http.Error(w, err.Error(), 400)
		log.Print("Error register parsing")
		log.Print(err.Error())
		return
	}
	log.Printf("%s", user.Name)
	if _, err := fmt.Fprintf(w, "Welcome %s", user.Name); err != nil {
		log.Fatal("Error register response")
	}
}

func main() {
	log.Print("Init server.")

	router := mux.NewRouter()
	router.HandleFunc("/api", GetHomeEndpoint).Methods("GET")
	router.HandleFunc("/api/register", RegisterEndpoint).Methods("POST")
	router.Handle("/", http.FileServer(http.Dir("static/")))
	log.Fatal(http.ListenAndServe(":3004", router))
}
