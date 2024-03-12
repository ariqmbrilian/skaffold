package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	appRequest()
}

func homePage(response http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(response, "Welcome to the Go Web API!")
	fmt.Println("Endpoint Hit: homePage")
}

func appRequest() {
	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
