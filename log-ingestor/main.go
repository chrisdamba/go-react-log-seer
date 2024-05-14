package main

import (
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
	"github.com/chrisdamba/go-react-log-seer/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/logs", handlers.IngestLog).Methods("POST")
	
	http.Handle("/", r)
	log.Println("Starting server on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
