package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipCode"`
}

func main() {
	http.HandleFunc("/greet", handleGreet)
	http.HandleFunc("/customers", handleCusomters)

	fmt.Println("Listening on Port 8000")
	http.ListenAndServe("localhost:8000", nil)
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func handleCusomters(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Ashish", "New Delhi", "110075"},
		{"Rob", "New Delhi", "110075"},
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
