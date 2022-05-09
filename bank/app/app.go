package main

import (
	"fmt"
	"log"
	"net/http"
)

func start() {
	http.HandleFunc("/greet", handleGreet)
	http.HandleFunc("/customers", handleCusomters)

	fmt.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
