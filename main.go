package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/theantichris/todo-api/handlers"
	"gopkg.in/mgo.v2"
)

// TODO: add to env
var port = ":8000"
var session, _ = mgo.Dial("127.0.0.1")
var c = session.DB("TodoDB").C("Todo")

func main() {
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	// TODO: gracefully exit on interrupt
	log.Fatal(http.ListenAndServe(port, router))
}
