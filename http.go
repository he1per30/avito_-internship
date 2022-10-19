package main

import (
	"log"
	"net/http"
)

type clientBalance struct {
	ID      string  `json:"id"`
	Balance float64 `json:"balance"`
}

func firstMethod() {
	http.HandleFunc("/get", GetByID)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func GetByID(resp http.ResponseWriter, req *http.Request) {

}
