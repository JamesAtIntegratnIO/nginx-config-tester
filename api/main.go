package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func health2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
	fmt.Println("Endpoint Hit: health2")
}

func handleRequests() {
	http.HandleFunc("/health", homePage)
	http.HandleFunc("/health2", health2)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
