package main

import (
	"encoding/json"
	"fmt"
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
		log.Print("Error parsing register ")
		log.Print(err.Error())
		return
	}
	log.Printf("%s", user.Name)
	if _, err := fmt.Fprintf(w, "Welcome %s", user.Name); err != nil {
		log.Fatal("Error register response")
	}
}

type Relation struct {
	Scanned User `json:"trusts"`
	Scans   User `json:"trusted"`
}

func RelationEndpoint(w http.ResponseWriter, req *http.Request) {
	var rel Relation

	d := json.NewDecoder(req.Body)
	if err := d.Decode(&rel); err != nil {
		http.Error(w, err.Error(), 400)
		log.Print("Error parsing relation")
		log.Print(err.Error())
		return
	}
	log.Printf("%s tursts %s", rel.Scanned.Name, rel.Scans.Name)
	if _, err := fmt.Fprintf(w, "%s added to your connections.", rel.Scanned.Name); err != nil {
		log.Fatal("Error relation response")
	}
}
