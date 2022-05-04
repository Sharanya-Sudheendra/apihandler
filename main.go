package main

import (
	"log"
	"net/http"
	"encoding/json" 
	"github.com/gorilla/mux" 
	"gorm.io/gorm" 
)


type Handler struct {
	
	gorm .Model 
	Sender string `json:"sender"`
	Subject string `json:"subject"`
	Body string `json:"body"`
	Recepient string `json:"recepient"` 

} 

var handler []Handler

func GetHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(handler) 
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&handler) 
	json.NewEncoder(w).Encode(handler) 
}



func main() {

	router := mux.NewRouter() 
	router.HandleFunc("/handler", GetHandler).Methods("GET") 
	router.HandleFunc("/handler", CreateHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router)) 

} 