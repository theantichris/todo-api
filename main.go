package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/theantichris/todo-api/handlers"
	"gopkg.in/mgo.v2"
)

// TODO: add to env
var port = ":8000"
var db, _ = mgo.Dial("127.0.0.1")
var c = db.DB("TodoDB").C("Todo")

func main() {
	db.SetMode(mgo.Monotonic, true)
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods("GET")

	server := &http.Server{
		Addr:         "0.0.0.0" + port,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)

	<-interruptChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	server.Shutdown(ctx)

	log.Println("shutting down")
	os.Exit(0)
}
