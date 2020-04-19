package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Card struct {
	Name              string `json:"name"`
	VictoryPoints     int    `json:"victoryPoints"`
	AdditionalActions int    `json:"additionalActions"`
	Type              string `json:"type"`
	AdditionalBuys    int    `json:"additionalBuys"`
	Treasure          int    `json:"treasure"`
	AdditionalCards   int    `json:"additionalCards"`
	Cost              int    `json:"cost"`
}

type BankCard struct {
	Card     Card `json:"card"`
	Quantity int  `json:"quantity"`
}

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

func bankData(w http.ResponseWriter, r *http.Request) {
	var allCards = []BankCard{
		BankCard{Card{"Gold", 0, 0, "Treasure", 0, 3, 0, 6}, 10},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(allCards)
}

func main() {
	fmt.Println("Starting Dominion web server")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/bank", bankData)
	log.Fatal(http.ListenAndServe(":8080", router))
}
