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
var session, _ = mgo.Dial("127.0.0.1")
var db = session.DB("TodoDB").C("Todo")

func main() {
	session.SetMode(mgo.Monotonic, true)
	defer session.Close()

	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthCheckHandler).Methods(http.MethodGet)
	router.HandleFunc("/todo", handlers.AddTodoItemHandler(db)).Methods(http.MethodPost)

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
