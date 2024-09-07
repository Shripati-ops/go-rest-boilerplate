package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)
func main() {
    // Create a new router
    r := mux.NewRouter()
    // Routes consist of a path and a handler function
    r.HandleFunc("/", home).Methods("GET")


    // Start the server
    log.Println("Starting server on :8000")
    if err := http.ListenAndServe(":8000", r); err != nil {
        log.Fatal(err)
    }
}

func home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"Request Received")
}